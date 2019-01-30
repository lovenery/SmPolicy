/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type RanNasRelCause struct {

	NgApCause *NgApCause `json:"ngApCause,omitempty"`

	Var5gMmCause int32 `json:"5gMmCause,omitempty"`

	Var5gSmCause int32 `json:"5gSmCause,omitempty"`
}
