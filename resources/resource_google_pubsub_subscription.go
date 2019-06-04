package resources

import (
	"strings"
)

type google_pubsub_subscription struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			AckDeadlineSeconds       string            `json:"ack_deadline_seconds"`
			ID                       string            `json:"id"`
			Labels                   map[string]string `json:"labels"`
			MessageRetentionDuration string            `json:"message_retention_duration"`
			Name                     string            `json:"name"`
			Path                     string            `json:"path"`
			Project                  string            `json:"project"`
			PushConfig               []Params          `json:"push_config"`
			RetainAckedMessages      string            `json:"retain_acked_messages"`
			Topic                    string            `json:"topic"`
			ExpirationPolicy         []Params          `json:"expiration_policy"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

type PushConfig struct {
	Number       string
	PushEndpoint string `json:"push_endpoint"`
	Attributes   string `json:"attributes"`
}

func (d *PushConfig) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "push_endpoint") {
		d.PushEndpoint = value
	} else if strings.Contains(key, "attributes") {
		d.Attributes = value
	}
}

func (d *PushConfig) SetNumber(num string) {
	d.Number = num
}

func (d *PushConfig) GetNumber() string {
	return d.Number
}

func NewPushConfig(num string) Params {
	var s Params = new(PushConfig)
	s.SetNumber(num)
	return s
}

type ExpirationPolicy struct {
	Number string
	TTL    string `json:"ttl"`
}

func (d *ExpirationPolicy) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "ttl") {
		d.TTL = value
	}
}

func (d *ExpirationPolicy) SetNumber(num string) {
	d.Number = num
}

func (d *ExpirationPolicy) GetNumber() string {
	return d.Number
}

func NewExpirationPolicy(num string) Params {
	var s Params = new(ExpirationPolicy)
	s.SetNumber(num)
	return s
}

func Newgoogle_pubsub_subscription(resourceData map[string]interface{}, resourceName string) *google_pubsub_subscription {

	resource := new(google_pubsub_subscription)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform

	resource.Primary.Attributes.PushConfig = CreateStructSlice(attributes_map, attributesKey, "push_config", 1, NewPushConfig)
	resource.Primary.Attributes.ExpirationPolicy = CreateStructSlice(attributes_map, attributesKey, "expiration_policy", 1, NewExpirationPolicy)
	resource.Primary.Attributes.Labels = CreateMap(attributes_map, "labels")

	return resource
}
