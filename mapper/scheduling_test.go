package mapper

import (
	"encoding/json"
	"github.com/rancher/norman/types/convert"
	"reflect"
	"testing"

	"github.com/rancher/norman/types/values"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestNodeAffinity(t *testing.T) {
	s := SchedulingMapper{}

	testInputs := []map[string]interface{}{
		{
			"scheduling": map[string]interface{}{
				"node": map[string]interface{}{
					"requireAll": []string{"foo = bar && foo1 != bar1"},
				},
			},
		},
	}

	expected := []v1.Affinity{
		{
			NodeAffinity: &v1.NodeAffinity{
				RequiredDuringSchedulingIgnoredDuringExecution: &v1.NodeSelector{
					NodeSelectorTerms: []v1.NodeSelectorTerm{
						{
							MatchExpressions: []v1.NodeSelectorRequirement{
								{
									Key:      "foo",
									Operator: v1.NodeSelectorOpIn,
									Values:   []string{"bar"},
								},
								{
									Key:      "foo1",
									Operator: v1.NodeSelectorOpNotIn,
									Values:   []string{"bar1"},
								},
							},
						},
					},
				},
			},
		},
	}
	for i, data := range testInputs {
		s.ToInternal(data)
		buf, err := json.Marshal(expected[i])
		if err != nil {
			t.Fatal(err)
		}
		var m map[string]interface{}
		if err := json.Unmarshal(buf, &m); err != nil {
			t.Fatal(err)
		}
		gotNodeAffinity, ok := values.GetValue(data, "affinity", "nodeAffinity", "requiredDuringSchedulingIgnoredDuringExecution")
		if !ok {
			t.Fatal("key not found")
		}
		expectedNodeAffinity, ok := values.GetValue(m, "nodeAffinity", "requiredDuringSchedulingIgnoredDuringExecution")
		if !ok {
			t.Fatal("key not found")
		}
		if !reflect.DeepEqual(gotNodeAffinity, expectedNodeAffinity) {
			t.Fatalf("test result not match! expected: %+v, got: %+v", expectedNodeAffinity, gotNodeAffinity)
		}
	}
}

func TestPodAffinity(t *testing.T) {
	s := SchedulingMapper{}

	testCases := []struct {
		name     string
		input    map[string]interface{}
		expected v1.Affinity
	}{
		{
			name: "pod affinity and antiAffinity test",
			input: map[string]interface{}{
				"scheduling": map[string]interface{}{
					"pod": map[string]interface{}{
						"affinity": map[string]interface{}{
							"required": []map[string]interface{}{
								{
									"topologyKey": "kubernetes.io/hostname",
									"rules":       []string{"foo = bar && foo1 != bar1"},
								},
							},
							"preferred": []map[string]interface{}{
								{
									"topologyKey": "kubernetes.io/hostname",
									"rules":       []string{"foo = bar", "foo1 != bar1"},
								},
							},
						},
						"antiAffinity": map[string]interface{}{
							"required": []map[string]interface{}{
								{
									"topologyKey": "kubernetes.io/hostname",
									"rules":       []string{"foo2 = bar2"},
								},
							},
							"preferred": []map[string]interface{}{
								{
									"topologyKey": "kubernetes.io/hostname",
									"rules":       []string{"foo3 = bar3"},
								},
							},
						},
					},
				},
			},
			expected: v1.Affinity{
				PodAffinity: &v1.PodAffinity{
					RequiredDuringSchedulingIgnoredDuringExecution: []v1.PodAffinityTerm{
						{
							LabelSelector: &metav1.LabelSelector{
								MatchExpressions: []metav1.LabelSelectorRequirement{
									{
										Key:      "foo",
										Operator: metav1.LabelSelectorOpIn,
										Values:   []string{"bar"},
									},
									{
										Key:      "foo1",
										Operator: metav1.LabelSelectorOpNotIn,
										Values:   []string{"bar1"},
									},
								},
							},
							TopologyKey: "kubernetes.io/hostname",
						},
					},
					PreferredDuringSchedulingIgnoredDuringExecution: []v1.WeightedPodAffinityTerm{
						{
							Weight: 100,
							PodAffinityTerm: v1.PodAffinityTerm{
								LabelSelector: &metav1.LabelSelector{
									MatchExpressions: []metav1.LabelSelectorRequirement{
										{
											Key:      "foo",
											Operator: metav1.LabelSelectorOpIn,
											Values:   []string{"bar"},
										},
										{
											Key:      "foo1",
											Operator: metav1.LabelSelectorOpNotIn,
											Values:   []string{"bar1"},
										},
									},
								},
								TopologyKey: "kubernetes.io/hostname",
							},
						},
					},
				},
				PodAntiAffinity: &v1.PodAntiAffinity{
					RequiredDuringSchedulingIgnoredDuringExecution: []v1.PodAffinityTerm{
						{
							LabelSelector: &metav1.LabelSelector{
								MatchExpressions: []metav1.LabelSelectorRequirement{
									{
										Key:      "foo2",
										Operator: metav1.LabelSelectorOpIn,
										Values:   []string{"bar2"},
									},
								},
							},
							TopologyKey: "kubernetes.io/hostname",
						},
					},
					PreferredDuringSchedulingIgnoredDuringExecution: []v1.WeightedPodAffinityTerm{
						{
							Weight: 100,
							PodAffinityTerm: v1.PodAffinityTerm{
								LabelSelector: &metav1.LabelSelector{
									MatchExpressions: []metav1.LabelSelectorRequirement{
										{
											Key:      "foo3",
											Operator: metav1.LabelSelectorOpIn,
											Values:   []string{"bar3"},
										},
									},
								},
								TopologyKey: "kubernetes.io/hostname",
							},
						},
					},
				},
			},
		},
		{
			name: "pod affinity rules test",
			input: map[string]interface{}{
				"scheduling": map[string]interface{}{
					"pod": map[string]interface{}{
						"affinity": map[string]interface{}{
							"required": []map[string]interface{}{
								{
									"topologyKey": "kubernetes.io/hostname",
									"rules":       []string{"foo = bar", "foo1 != bar1", "foo2", "!foo3"},
								},
							},
						},
					},
				},
			},
			expected: v1.Affinity{
				PodAffinity: &v1.PodAffinity{
					RequiredDuringSchedulingIgnoredDuringExecution: []v1.PodAffinityTerm{
						{
							LabelSelector: &metav1.LabelSelector{
								MatchExpressions: []metav1.LabelSelectorRequirement{
									{
										Key:      "foo",
										Operator: metav1.LabelSelectorOpIn,
										Values:   []string{"bar"},
									},
									{
										Key:      "foo1",
										Operator: metav1.LabelSelectorOpNotIn,
										Values:   []string{"bar1"},
									},
									{
										Key:      "foo2",
										Operator: metav1.LabelSelectorOpExists,
									},
									{
										Key:      "foo3",
										Operator: metav1.LabelSelectorOpDoesNotExist,
									},
								},
							},
							TopologyKey: "kubernetes.io/hostname",
						},
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		data := testCase.input
		err := s.ToInternal(data)
		if err != nil {
			t.Fatal(err)
		}
		var expectedAffinity map[string]interface{}
		expectedAffinity, err = convert.EncodeToMap(testCase.expected)
		if err != nil {
			t.Fatal(err)
		}
		gotAffinity, ok := values.GetValue(data, "affinity")
		if !ok {
			t.Fatal("key not found")
		}
		if !reflect.DeepEqual(gotAffinity, expectedAffinity) {
			got, _ := json.Marshal(gotAffinity)
			expected, _ := json.Marshal(expectedAffinity)
			t.Fatalf("test result not match! expected:\n%s, got:\n%s", string(expected), string(got))
			//t.Fatalf("test result not match! expected: %+v, got: %+v", gotAffinity, expectedAffinity)
		}
	}
}

//func TestPodAffinityFromInternal(t *testing.T) {
//	s := SchedulingMapper{}
//
//	testCases := []struct {
//		name     string
//		input    map[string]interface{}
//		expected v1.Affinity
//	}{
//		{
//			name: "podAffinity",
//			input: map[string]interface{}{
//				"scheduling": map[string]interface{}{
//					"pod": map[string]interface{}{
//						"affinity": map[string]interface{}{
//							"required": []map[string]interface{}{
//								{
//									"namespaces":  nil,
//									"topologyKey": "kubernetes.io/hostname",
//									"rules":       []string{"foo = bar", "foo1 != bar1"},
//								},
//							},
//							"preferred": []map[string]interface{}{
//								{
//									"namespaces":  nil,
//									"topologyKey": "kubernetes.io/hostname",
//									"rules":       []string{"foo = bar", "foo1 != bar1"},
//								},
//							},
//						},
//					},
//				},
//			},
//			expected: v1.Affinity{
//				PodAffinity: &v1.PodAffinity{
//					RequiredDuringSchedulingIgnoredDuringExecution: []v1.PodAffinityTerm{
//						{
//							LabelSelector: &metav1.LabelSelector{
//								MatchExpressions: []metav1.LabelSelectorRequirement{
//									{
//										Key:      "foo",
//										Operator: metav1.LabelSelectorOpIn,
//										Values:   []string{"bar"},
//									},
//									{
//										Key:      "foo1",
//										Operator: metav1.LabelSelectorOpNotIn,
//										Values:   []string{"bar1"},
//									},
//								},
//							},
//							TopologyKey: "kubernetes.io/hostname",
//						},
//					},
//					PreferredDuringSchedulingIgnoredDuringExecution: []v1.WeightedPodAffinityTerm{
//						{
//							Weight: 100,
//							PodAffinityTerm: v1.PodAffinityTerm{
//								LabelSelector: &metav1.LabelSelector{
//									MatchExpressions: []metav1.LabelSelectorRequirement{
//										{
//											Key:      "foo",
//											Operator: metav1.LabelSelectorOpIn,
//											Values:   []string{"bar"},
//										},
//										{
//											Key:      "foo1",
//											Operator: metav1.LabelSelectorOpNotIn,
//											Values:   []string{"bar1"},
//										},
//									},
//								},
//								TopologyKey: "kubernetes.io/hostname",
//							},
//						},
//					},
//				},
//			},
//		},
//	}
//
//	for _, testCase := range testCases {
//		expectedAffinity, err := convert.EncodeToMap(testCase.expected)
//		if err != nil {
//			t.Fatal(err)
//		}
//		data := map[string]interface{}{"affinity": expectedAffinity}
//		s.FromInternal(data)
//		if !reflect.DeepEqual(data, testCase.input) {
//			expected, _ := json.Marshal(testCase.input)
//			got, _ := json.Marshal(data)
//			t.Fatalf("test result not match! expected:\n%s, got:\n%s", string(expected), string(got))
//			//t.Fatalf("test result not match! expected: %+v, got: %+v", gotAffinity, expectedAffinity)
//		}
//	}
//}
