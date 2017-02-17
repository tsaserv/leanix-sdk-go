/* 
 * LeanIX Pathfinder REST API
 *
 * Core application for storage and analysis of IT landscape data
 *
 * OpenAPI spec version: 0.1.10-SNAPSHOT
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

import (
	"time"
)

type FactSheet struct {

	Id string `json:"id,omitempty"`

	Name string `json:"name"`

	Description string `json:"description"`

	DisplayName string `json:"displayName,omitempty"`

	Type_ string `json:"type"`

	Tags []Tag `json:"tags,omitempty"`

	Fields []FieldnameAndData `json:"fields,omitempty"`

	Relations []Relation `json:"relations,omitempty"`

	RelationIds []string `json:"relationIds,omitempty"`

	IntentionallyNotSet []string `json:"intentionallyNotSet,omitempty"`

	Completion Completion `json:"completion,omitempty"`

	// when the fact sheet was last changed
	LastChanged time.Time `json:"lastChanged,omitempty"`

	Comments []GqlComment `json:"comments,omitempty"`

	// when the fact sheet was last approved
	ApprovedDate time.Time `json:"approvedDate,omitempty"`

	Status string `json:"status"`

	Rev int64 `json:"rev,omitempty"`
}