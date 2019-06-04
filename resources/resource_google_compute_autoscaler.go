package resources

import (
	"strings"
)

type google_compute_autoscaler struct {
	Type      string   `json:"type"`
	DependsOn []string `json:"depends_on"`
	Primary   struct {
		ID         string `json:"id"`
		Attributes struct {
			AutoscalingPolicy []Params `json:"autoscaling_policy_cooldown_period"`
			CreationTimestamp string   `json:"creation_timestamp"`
			Description       string   `json:"description"`
			ID                string   `json:"id"`
			Name              string   `json:"name"`
			Project           string   `json:"project"`
			SelfLink          string   `json:"self_link"`
			Target            string   `json:"target"`
			Zone              string   `json:"zone"`
		} `json:"attributes"`
		Meta struct {
			SchemaVersion string `json:"schema_version"`
		} `json:"meta"`
		Tainted bool `json:"tainted"`
	} `json:"primary"`
	Deposed  []interface{} `json:"deposed"`
	Provider string        `json:"provider"`
}

type AutoscalingPolicy struct {
	Number              string
	CooldownPeriod      string   `json:"cooldown_period"`
	CPUTarget           string   `json:"cpu_target"`
	LoadBalancingTarget string   `json:"load_balancing_target"`
	MaxReplica          string   `json:"max_replica"`
	MinReplica          string   `json:"min_replica"`
	Metric              []Params `json:"metric"`
}

func (d *AutoscalingPolicy) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "cooldown_period") {
		d.CooldownPeriod = value
	} else if strings.Contains(key, "cpu_target") {
		d.CPUTarget = value
	} else if strings.Contains(key, "load_balancing_target") {
		d.LoadBalancingTarget = value
	} else if strings.Contains(key, "max_replica") {
		d.MaxReplica = value
	} else if strings.Contains(key, "min_replica") {
		d.MinReplica = value
	} else if strings.Contains(key, "metric") && strings.Contains(key, "name") {
		//autoscaling_policy.#.metricまでを取得
		Number := CreateNumberFromKey(key, 2)
		d.Metric = CreateStructSlice(attributes_map, attributesKey, Number, 3, NewMetric)
	}
}

func (d *AutoscalingPolicy) SetNumber(num string) {
	d.Number = num
}

func (d *AutoscalingPolicy) GetNumber() string {
	return d.Number
}

func NewAutoscalingPolicy(num string) Params {
	var s Params = new(AutoscalingPolicy)
	s.SetNumber(num)
	return s
}

type Metric struct {
	Number string
	Target string
	Name   string
	Type   string
}

func (d *Metric) SetParameter(key, value string, attributes_map map[string]interface{}, attributesKey []string) {
	if strings.Contains(key, "target") {
		d.Target = value
	} else if strings.Contains(key, "name") {
		d.Name = value
	} else if strings.Contains(key, "type") {
		d.Type = value
	}
}

func (d *Metric) SetNumber(num string) {
	d.Number = num
}

func (d *Metric) GetNumber() string {
	return d.Number
}

func NewMetric(num string) Params {
	var s Params = new(Metric)
	s.SetNumber(num)
	return s
}

func Newgoogle_compute_autoscaler(resourceData map[string]interface{}, resourceName string) *google_compute_autoscaler {

	resource := new(google_compute_autoscaler)
	attributes_map := resourceData["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	attributesKey := CreateAttributesKey(attributes_map)
	MapToStruct(resourceData, resource)
	resource.Primary.ID = resourceName
	//resource.DependsOn = InterfaceToStrings(resourceData["depends_on"].([]interface{}))

	//attributes need to transform
	resource.Primary.Attributes.AutoscalingPolicy = CreateStructSlice(attributes_map, attributesKey, "autoscaling_policy", 1, NewAutoscalingPolicy)

	//beta
	if resourceData["provider"] == "provider.google-beta" {
	}
	return resource
}
