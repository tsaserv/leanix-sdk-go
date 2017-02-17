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

type ConfigTypeOnOffWithTime struct {

	Enabled bool `json:"enabled,omitempty"`

	Value int64 `json:"value,omitempty"`

	Unit string `json:"unit,omitempty"`
}
