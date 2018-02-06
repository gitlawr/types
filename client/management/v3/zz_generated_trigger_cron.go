package client

const (
	TriggerCronType          = "triggerCron"
	TriggerCronFieldSpec     = "spec"
	TriggerCronFieldTimezone = "timezone"
)

type TriggerCron struct {
	Spec     string `json:"spec,omitempty"`
	Timezone string `json:"timezone,omitempty"`
}
