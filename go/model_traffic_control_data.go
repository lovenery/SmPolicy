/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type TrafficControlData struct {

	// Univocally identifies the traffic control policy data within a PDU session.
	TcId string `json:"tcId"`

	FlowStatus *FlowStatus `json:"flowStatus,omitempty"`

	RedirectInfo *RedirectInformation `json:"redirectInfo,omitempty"`

	// Indicates whether applicat'on's start or stop notification is to be muted.
	MuteNotif bool `json:"muteNotif,omitempty"`

	// Reference to a pre-configured traffic steering policy for downlink traffic at the SMF.
	TrafficSteeringPolIdDl string `json:"trafficSteeringPolIdDl,omitempty"`

	// Reference to a pre-configured traffic steering policy for uplink traffic at the SMF.
	TrafficSteeringPolIdUl string `json:"trafficSteeringPolIdUl,omitempty"`

	// A list of location which the traffic shall be routed to for the AF request
	RouteToLocs []RouteToLocation `json:"routeToLocs,omitempty"`

	UpPathChgEvent *UpPathChgEvent `json:"upPathChgEvent,omitempty"`
}
