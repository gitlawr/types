package mapper

import (
	"fmt"
	"sort"
	"strings"

	"regexp"

	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/convert"
	"github.com/rancher/norman/types/values"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	exprRegexp = regexp.MustCompile("^(.*?)\\s*(=|!=|<|>| in | notin )\\s*(.*)$")
)

type SchedulingMapper struct {
}

func (s SchedulingMapper) FromInternal(data map[string]interface{}) {
	defer func() {
		delete(data, "nodeSelector")
		delete(data, "affinity")
	}()

	var requireAll []string
	for key, value := range convert.ToMapInterface(data["nodeSelector"]) {
		if value == "" {
			requireAll = append(requireAll, key)
		} else {
			requireAll = append(requireAll, fmt.Sprintf("%s = %s", key, value))
		}
	}

	if len(requireAll) > 0 {
		values.PutValue(data, requireAll, "scheduling", "node", "requireAll")
	}

	v, ok := data["affinity"]
	if !ok || v == nil {
		return
	}

	affinity := &v1.Affinity{}
	if err := convert.ToObj(v, affinity); err != nil {
		return
	}

	if affinity.NodeAffinity != nil {
		s.nodeAffinity(data, affinity.NodeAffinity)
	}

	if affinity.PodAffinity != nil {
		s.podAffinityOrAntiAffinity(data, affinity.PodAffinity)
	}

	if affinity.PodAntiAffinity != nil {
		s.podAffinityOrAntiAffinity(data, affinity.PodAntiAffinity)
	}
}

func (s SchedulingMapper) nodeAffinity(data map[string]interface{}, nodeAffinity *v1.NodeAffinity) {
	var requireAll []string
	var requireAny []string
	var preferred []string

	if nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution != nil {
		for _, term := range nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
			exprs := NodeSelectorTermToStrings(term)
			if len(exprs) == 0 {
				continue
			}
			if len(requireAny) > 0 {
				// Once any is set all new terms go to any
				requireAny = append(requireAny, strings.Join(exprs, " && "))
			} else if len(requireAll) > 0 {
				// If all is already set, we actually need to move everything to any
				requireAny = append(requireAny, strings.Join(requireAll, " && "))
				requireAny = append(requireAny, strings.Join(exprs, " && "))
				requireAll = []string{}
			} else {
				// The first term is considered all
				requireAll = exprs
			}
		}
	}

	if nodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution != nil {
		sortPreferred(nodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution)
		for _, term := range nodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
			exprs := NodeSelectorTermToStrings(term.Preference)
			preferred = append(preferred, strings.Join(exprs, " && "))
		}
	}

	if len(requireAll) > 0 {
		values.PutValue(data, requireAll, "scheduling", "node", "requireAll")
	}
	if len(requireAny) > 0 {
		values.PutValue(data, requireAny, "scheduling", "node", "requireAny")
	}
	if len(preferred) > 0 {
		values.PutValue(data, preferred, "scheduling", "node", "preferred")
	}
}

// podAffinityOrAntiAffinity sets podAffinity/podAntiAffinity to data map
func (s SchedulingMapper) podAffinityOrAntiAffinity(data map[string]interface{}, podAffinity interface{}) {
	var affinity = make(map[string]interface{})
	var requiredField []v1.PodAffinityTerm
	var preferredField []v1.WeightedPodAffinityTerm
	var required []interface{}
	var preferred []interface{}
	var isAffinity bool

	switch t := podAffinity.(type) {
	case *v1.PodAffinity:
		requiredField = t.RequiredDuringSchedulingIgnoredDuringExecution
		preferredField = t.PreferredDuringSchedulingIgnoredDuringExecution
		isAffinity = true
	case *v1.PodAntiAffinity:
		requiredField = t.RequiredDuringSchedulingIgnoredDuringExecution
		preferredField = t.PreferredDuringSchedulingIgnoredDuringExecution
	default:
		return
	}
	if requiredField == nil && preferredField == nil {
		return
	}

	if requiredField != nil {
		for _, term := range requiredField {
			rules := PodAffinityTermToStrings(term)
			namespaces := term.Namespaces
			topologyKey := term.TopologyKey
			required = append(required, map[string]interface{}{
				"namespaces":  namespaces,
				"rules":       rules,
				"topologyKey": topologyKey,
			})
		}
		affinity["required"] = required
	}

	if preferredField != nil {
		for _, preferredTerm := range preferredField {
			term := preferredTerm.PodAffinityTerm
			rules := PodAffinityTermToStrings(term)
			namespaces := term.Namespaces
			topologyKey := term.TopologyKey
			preferred = append(preferred, map[string]interface{}{
				"namespaces":  namespaces,
				"rules":       rules,
				"topologyKey": topologyKey,
			})
		}
		affinity["preferred"] = preferred
	}

	if isAffinity {
		values.PutValue(data, affinity, "scheduling", "pod", "affinity")
	} else {
		values.PutValue(data, affinity, "scheduling", "pod", "antiAffinity")
	}
}

func sortPreferred(terms []v1.PreferredSchedulingTerm) {
	sort.Slice(terms, func(i, j int) bool {
		return terms[i].Weight > terms[j].Weight
	})
}

func NodeSelectorTermToStrings(term v1.NodeSelectorTerm) []string {
	exprs := []string{}

	for _, expr := range term.MatchExpressions {
		nextExpr := ""
		switch expr.Operator {
		case v1.NodeSelectorOpIn:
			if len(expr.Values) > 1 {
				nextExpr = fmt.Sprintf("%s in (%s)", expr.Key, strings.Join(expr.Values, ", "))
			} else if len(expr.Values) == 1 {
				nextExpr = fmt.Sprintf("%s = %s", expr.Key, expr.Values[0])
			}
		case v1.NodeSelectorOpNotIn:
			if len(expr.Values) > 1 {
				nextExpr = fmt.Sprintf("%s notin (%s)", expr.Key, strings.Join(expr.Values, ", "))
			} else if len(expr.Values) == 1 {
				nextExpr = fmt.Sprintf("%s != %s", expr.Key, expr.Values[0])
			}
		case v1.NodeSelectorOpExists:
			nextExpr = expr.Key
		case v1.NodeSelectorOpDoesNotExist:
			nextExpr = "!" + expr.Key
		case v1.NodeSelectorOpGt:
			if len(expr.Values) == 1 {
				nextExpr = fmt.Sprintf("%s > %s", expr.Key, expr.Values[0])
			}
		case v1.NodeSelectorOpLt:
			if len(expr.Values) == 1 {
				nextExpr = fmt.Sprintf("%s < %s", expr.Key, expr.Values[0])
			}
		}

		if nextExpr != "" {
			exprs = append(exprs, nextExpr)
		}
	}

	return exprs
}

func PodAffinityTermToStrings(term v1.PodAffinityTerm) []string {
	exprs := []string{}

	for _, expr := range term.LabelSelector.MatchExpressions {
		nextExpr := ""
		switch expr.Operator {
		case metav1.LabelSelectorOpIn:
			if len(expr.Values) > 1 {
				nextExpr = fmt.Sprintf("%s in (%s)", expr.Key, strings.Join(expr.Values, ", "))
			} else if len(expr.Values) == 1 {
				nextExpr = fmt.Sprintf("%s = %s", expr.Key, expr.Values[0])
			}
		case metav1.LabelSelectorOpNotIn:
			if len(expr.Values) > 1 {
				nextExpr = fmt.Sprintf("%s notin (%s)", expr.Key, strings.Join(expr.Values, ", "))
			} else if len(expr.Values) == 1 {
				nextExpr = fmt.Sprintf("%s != %s", expr.Key, expr.Values[0])
			}
		case metav1.LabelSelectorOpExists:
			nextExpr = expr.Key
		case metav1.LabelSelectorOpDoesNotExist:
			nextExpr = "!" + expr.Key
		}
		if nextExpr != "" {
			exprs = append(exprs, nextExpr)
		}
	}

	for k, v := range term.LabelSelector.MatchLabels {
		exprs = append(exprs, fmt.Sprintf("%s = %s", k, v))
	}
	return exprs
}

func StringsToNodeSelectorTerm(exprs []string) []v1.NodeSelectorTerm {
	result := []v1.NodeSelectorTerm{}

	for _, inter := range exprs {
		term := v1.NodeSelectorTerm{}

		for _, expr := range strings.Split(inter, "&&") {
			groups := exprRegexp.FindStringSubmatch(expr)
			selectorRequirement := v1.NodeSelectorRequirement{}

			if groups == nil {
				if strings.HasPrefix(expr, "!") {
					selectorRequirement.Key = strings.TrimSpace(expr[1:])
					selectorRequirement.Operator = v1.NodeSelectorOpDoesNotExist
				} else {
					selectorRequirement.Key = strings.TrimSpace(expr)
					selectorRequirement.Operator = v1.NodeSelectorOpExists
				}
			} else {
				selectorRequirement.Key = strings.TrimSpace(groups[1])
				selectorRequirement.Values = convert.ToValuesSlice(groups[3])
				op := strings.TrimSpace(groups[2])
				switch op {
				case "=":
					selectorRequirement.Operator = v1.NodeSelectorOpIn
				case "!=":
					selectorRequirement.Operator = v1.NodeSelectorOpNotIn
				case "notin":
					selectorRequirement.Operator = v1.NodeSelectorOpNotIn
				case "in":
					selectorRequirement.Operator = v1.NodeSelectorOpIn
				case "<":
					selectorRequirement.Operator = v1.NodeSelectorOpLt
				case ">":
					selectorRequirement.Operator = v1.NodeSelectorOpGt
				}
			}

			term.MatchExpressions = append(term.MatchExpressions, selectorRequirement)
		}

		result = append(result, term)
	}

	return result
}

func StringsToLabelSelectorTerm(exprs []string) []metav1.LabelSelectorRequirement {
	result := []metav1.LabelSelectorRequirement{}

	for _, inter := range exprs {
		for _, expr := range strings.Split(inter, "&&") {
			term := metav1.LabelSelectorRequirement{}
			groups := exprRegexp.FindStringSubmatch(expr)
			if groups == nil {
				if strings.HasPrefix(expr, "!") {
					term.Key = strings.TrimSpace(expr[1:])
					term.Operator = metav1.LabelSelectorOpDoesNotExist
				} else {
					term.Key = strings.TrimSpace(expr)
					term.Operator = metav1.LabelSelectorOpExists
				}
			} else {
				term.Key = strings.TrimSpace(groups[1])
				term.Values = convert.ToValuesSlice(groups[3])
				op := strings.TrimSpace(groups[2])
				switch op {
				case "=":
					term.Operator = metav1.LabelSelectorOpIn
				case "!=":
					term.Operator = metav1.LabelSelectorOpNotIn
				case "notin":
					term.Operator = metav1.LabelSelectorOpNotIn
				case "in":
					term.Operator = metav1.LabelSelectorOpIn
				}
			}
			result = append(result, term)
		}
	}
	return result
}

func (s SchedulingMapper) ToInternal(data map[string]interface{}) error {
	defer func() {
		delete(data, "scheduling")
	}()

	setNodeAffinity(data)
	setPodAffinity(data)
	setPodAntiAffinity(data)

	return nil
}

func setNodeAffinity(data map[string]interface{}) {
	nodeName := convert.ToString(values.GetValueN(data, "scheduling", "node", "nodeId"))
	if nodeName != "" {
		data["nodeName"] = nodeName
	}

	requireAllV := values.GetValueN(data, "scheduling", "node", "requireAll")
	requireAnyV := values.GetValueN(data, "scheduling", "node", "requireAny")
	preferredV := values.GetValueN(data, "scheduling", "node", "preferred")

	if requireAllV == nil && requireAnyV == nil && preferredV == nil {
		return
	}

	requireAll := convert.ToStringSlice(requireAllV)
	requireAny := convert.ToStringSlice(requireAnyV)
	preferred := convert.ToStringSlice(preferredV)

	if len(requireAll) == 0 && len(requireAny) == 0 && len(preferred) == 0 {
		values.PutValue(data, nil, "affinity", "nodeAffinity")
		return
	}

	nodeAffinity := v1.NodeAffinity{}

	if len(requireAll) > 0 {
		nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution = &v1.NodeSelector{
			NodeSelectorTerms: []v1.NodeSelectorTerm{
				AggregateTerms(StringsToNodeSelectorTerm(requireAll)),
			},
		}
	}

	if len(requireAny) > 0 {
		if nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution == nil {
			nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution = &v1.NodeSelector{}
		}
		nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms = StringsToNodeSelectorTerm(requireAny)
	}

	if len(preferred) > 0 {
		count := int32(100)
		for _, term := range StringsToNodeSelectorTerm(preferred) {
			nodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution = append(
				nodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution, v1.PreferredSchedulingTerm{
					Weight:     count,
					Preference: term,
				})
			count--
		}
	}

	affinity, _ := convert.EncodeToMap(&v1.Affinity{
		NodeAffinity: &nodeAffinity,
	})

	if nodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution == nil {
		values.PutValue(affinity, nil, "nodeAffinity", "preferredDuringSchedulingIgnoredDuringExecution")
	}

	if nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution == nil {
		values.PutValue(affinity, nil, "nodeAffinity", "requiredDuringSchedulingIgnoredDuringExecution")
	}

	data["affinity"] = affinity
}

func setPodAffinity(data map[string]interface{}) {
	requiredV := values.GetValueN(data, "scheduling", "pod", "affinity", "required")
	preferredV := values.GetValueN(data, "scheduling", "pod", "affinity", "preferred")
	required := convert.ToMapSlice(requiredV)
	preferred := convert.ToMapSlice(preferredV)
	podAffinity := &v1.PodAffinity{}
	if len(required) == 0 && len(preferred) == 0 {
		values.RemoveValue(data, "affinity", "podAffinity")
		return
	}

	if len(required) > 0 {
		var podAffinityTerms []v1.PodAffinityTerm
		for _, rule := range required {
			podAffinityTerm := v1.PodAffinityTerm{
				Namespaces:  convert.ToStringSlice(rule["namespaces"]),
				TopologyKey: convert.ToString(rule["topologyKey"]),
				LabelSelector: &metav1.LabelSelector{
					MatchExpressions: StringsToLabelSelectorTerm(convert.ToStringSlice(rule["rules"])),
				},
			}
			podAffinityTerms = append(podAffinityTerms, podAffinityTerm)
		}
		podAffinity.RequiredDuringSchedulingIgnoredDuringExecution = podAffinityTerms
	}

	if len(preferred) > 0 {
		var weightedPodAffinityTerms []v1.WeightedPodAffinityTerm
		count := int32(100)
		for _, rule := range preferred {
			weightedPodAffinityTerm := v1.WeightedPodAffinityTerm{
				Weight: count,
				PodAffinityTerm: v1.PodAffinityTerm{
					Namespaces:  convert.ToStringSlice(rule["namespaces"]),
					TopologyKey: convert.ToString(rule["topologyKey"]),
					LabelSelector: &metav1.LabelSelector{
						MatchExpressions: StringsToLabelSelectorTerm(convert.ToStringSlice(rule["rules"])),
					},
				},
			}
			count--
			weightedPodAffinityTerms = append(weightedPodAffinityTerms, weightedPodAffinityTerm)
		}
		podAffinity.PreferredDuringSchedulingIgnoredDuringExecution = weightedPodAffinityTerms
	}
	//FIXME
	podAffinityM, _ := convert.EncodeToMap(podAffinity)
	values.PutValue(data, podAffinityM, "affinity", "podAffinity")
	return
}

func setPodAntiAffinity(data map[string]interface{}) {
	requiredV := values.GetValueN(data, "scheduling", "pod", "antiAffinity", "required")
	preferredV := values.GetValueN(data, "scheduling", "pod", "antiAffinity", "preferred")
	required := convert.ToMapSlice(requiredV)
	preferred := convert.ToMapSlice(preferredV)
	podAntiAffinity := &v1.PodAntiAffinity{}
	if len(required) == 0 && len(preferred) == 0 {
		values.RemoveValue(data, "affinity", "podAntiAffinity")
		return
	}

	if len(required) > 0 {
		var podAffinityTerms []v1.PodAffinityTerm
		for _, rule := range required {
			podAffinityTerm := v1.PodAffinityTerm{
				Namespaces:  convert.ToStringSlice(rule["namespaces"]),
				TopologyKey: convert.ToString(rule["topologyKey"]),
				LabelSelector: &metav1.LabelSelector{
					MatchExpressions: StringsToLabelSelectorTerm(convert.ToStringSlice(rule["rules"])),
				},
			}
			podAffinityTerms = append(podAffinityTerms, podAffinityTerm)
		}
		podAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution = podAffinityTerms
	}

	if len(preferred) > 0 {
		var weightedPodAffinityTerms []v1.WeightedPodAffinityTerm
		count := int32(100)
		for _, rule := range preferred {
			weightedPodAffinityTerm := v1.WeightedPodAffinityTerm{
				Weight: count,
				PodAffinityTerm: v1.PodAffinityTerm{
					Namespaces:  convert.ToStringSlice(rule["namespaces"]),
					TopologyKey: convert.ToString(rule["topologyKey"]),
					LabelSelector: &metav1.LabelSelector{
						MatchExpressions: StringsToLabelSelectorTerm(convert.ToStringSlice(rule["rules"])),
					},
				},
			}
			count--
			weightedPodAffinityTerms = append(weightedPodAffinityTerms, weightedPodAffinityTerm)
		}
		podAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution = weightedPodAffinityTerms
	}
	//FIXME
	podAntiAffinityM, _ := convert.EncodeToMap(podAntiAffinity)
	values.PutValue(data, podAntiAffinityM, "affinity", "podAntiAffinity")
	return
}

func AggregateTerms(terms []v1.NodeSelectorTerm) v1.NodeSelectorTerm {
	result := v1.NodeSelectorTerm{}
	for _, term := range terms {
		result.MatchExpressions = append(result.MatchExpressions, term.MatchExpressions...)
	}
	return result
}

func (s SchedulingMapper) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	delete(schema.ResourceFields, "nodeSelector")
	delete(schema.ResourceFields, "affinity")
	return nil
}
