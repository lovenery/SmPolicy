/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type QosNotificationControlInfo struct {

	// An array of PCC rule id references to the PCC rules associated with the QoS notification control info.
	RefPccRuleIds []string `json:"refPccRuleIds"`

	NotifType *QosNotifType `json:"notifType"`

	// Represents the content version of some content.
	ContVer int32 `json:"contVer,omitempty"`
}