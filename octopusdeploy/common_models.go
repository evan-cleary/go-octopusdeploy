package octopusdeploy

import (
	"fmt"
)

type PagedResults struct {
	ItemType       string `json:"ItemType"`
	TotalResults   int    `json:"TotalResults"`
	NumberOfPages  int    `json:"NumberOfPages"`
	LastPageNumber int    `json:"LastPageNumber"`
	ItemsPerPage   int    `json:"ItemsPerPage"`
	IsStale        bool   `json:"IsStale"`
	Links          Links  `json:"Links"`
}

type Links struct {
	Self        string `json:"Self"`
	Template    string `json:"Template"`
	PageAll     string `json:"Page.All"`
	PageCurrent string `json:"Page.Current"`
	PageLast    string `json:"Page.Last"`
	PageNext    string `json:"Page.Next"`
}

type SensitivePropertyValue struct {
	HasValue bool   `json:"HasValue"`
	NewValue string `json:"NewValue"`
}

type PropertyValue string

// TODO: refactor to use the PropertyValueResource for handling sensitive values - https://blog.gopheracademy.com/advent-2016/advanced-encoding-decoding/
// type PropertyValueResource struct {
// 	IsSensitive    bool           `json:"IsSensitive,omitempty"`
// 	Value          string         `json:"Value,omitempty"`
// 	SensitiveValue SensitiveValue `json:"SensitiveValue,omitempty"`
// }

// type PropertyValueResource map[string]PropertyValueResourceData

// // PropertyValues can either be Secret, or not secret, which means they have different structs. Need custom Marshal/Unmarshal to check this.
// type PropertyValueResource struct {
// 	*SensitivePropertyValue
// 	*PropertyValue
// }

// func (d PropertyValueResource) MarshalJSON() ([]byte, error) {
// 	// check if the HasValue field actually exists on the object, if not, its a PropertyValue
// 	if d.SensitivePropertyValue.HasValue == true || d.SensitivePropertyValue.HasValue == false {
// 		return json.Marshal(d.SensitivePropertyValue)
// 	}

// 	return json.Marshal(d.PropertyValue)
// }

// func (d *PropertyValueResource) UnmarshalJSON(data []byte) error {
// 	// try unmarshal into a sensitive property, if that fails, it's just a normal property

// 	var spv SensitivePropertyValue
// 	errUnmarshalSensitivePropertyValue := json.Unmarshal(data, &spv)

// 	if errUnmarshalSensitivePropertyValue != nil {
// 		var pv PropertyValue
// 		errUnmarshalPropertyValue := json.Unmarshal(data, &pv)

// 		if errUnmarshalPropertyValue != nil {
// 			return errUnmarshalPropertyValue
// 		}

// 		d.PropertyValue = &pv
// 		d.SensitivePropertyValue = nil
// 		return nil
// 	}

// 	d.PropertyValue = nil
// 	d.SensitivePropertyValue = &spv
// 	return nil
// }

type DeploymentStep struct {
	ID                 string                   `json:"Id"`
	Name               string                   `json:"Name"`
	PackageRequirement string                   `json:"PackageRequirement,omitempty"` // may need its own model / enum
	Properties         map[string]PropertyValue `json:"Properties"`                   // TODO: refactor to use the PropertyValueResource for handling sensitive values - https://blog.gopheracademy.com/advent-2016/advanced-encoding-decoding/
	Condition          string                   `json:"Condition,omitempty"`          // needs enum
	StartTrigger       string                   `json:"StartTrigger,omitempty"`       // needs enum
	Actions            []DeploymentAction       `json:"Actions"`
}

type DeploymentAction struct {
	ID                            string                   `json:"Id"`
	Name                          string                   `json:"Name"`
	ActionType                    string                   `json:"ActionType"`
	IsDisabled                    bool                     `json:"IsDisabled"`
	CanBeUsedForProjectVersioning bool                     `json:"CanBeUsedForProjectVersioning"`
	Environments                  []string                 `json:"Environments"`
	ExcludedEnvironments          []string                 `json:"ExcludedEnvironments"`
	Channels                      []string                 `json:"Channels"`
	TenantTags                    []string                 `json:"TenantTags"`
	Properties                    map[string]PropertyValue `json:"Properties"`
	LastModifiedOn                string                   `json:"LastModifiedOn"` // datetime
	LastModifiedBy                string                   `json:"LastModifiedBy"`
	Links                         Links                    `json:"Links"` // may be wrong
}

type APIError struct {
	ErrorMessage  string   `json:"ErrorMessage"`
	Errors        []string `json:"Errors"`
	FullException string   `json:"FullException"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("Octopus Deploy Error Response: %v %+v %v", e.ErrorMessage, e.Errors, e.FullException)
}
