/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Arp struct {

	PriorityLevel int32 `json:"priorityLevel"`

	PreemptCap *PreemptionCapability `json:"preemptCap"`

	PreemptVuln *PreemptionVulnerability `json:"preemptVuln"`
}