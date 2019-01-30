/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PccRule struct {

	// An array of IP flow packet filter information.
	FlowInfos []FlowInformation `json:"flowInfos,omitempty"`

	// A reference to the application detection filter configured at the UPF.
	AppId string `json:"appId,omitempty"`

	// Represents the content version of some content.
	ContVer int32 `json:"contVer,omitempty"`

	// Univocally identifies the PCC rule within a PDU session.
	PccRuleId string `json:"pccRuleId"`

	// Determines the order in which this PCC rule is applied relative to other PCC rules within the same PDU session.
	Precedence int32 `json:"precedence,omitempty"`

	AfSigProtocol *AfSigProtocol `json:"afSigProtocol,omitempty"`

	// Indication of application relocation possibility.
	AppReloc bool `json:"appReloc,omitempty"`

	// A reference to the QoSData policy type decision type. It is the qosId described in subclause 5.6.2.8. (NOTE)
	RefQosData []string `json:"refQosData,omitempty"`

	// A reference to the TrafficControlData policy decision type. It is the tcId described in subclause 5.6.2.10. (NOTE)
	RefTcData []string `json:"refTcData,omitempty"`

	// A reference to the ChargingData policy decision type. It is the chgId described in subclause 5.6.2.11. (NOTE)
	RefChgData []string `json:"refChgData,omitempty"`

	// A reference to UsageMonitoringData policy decision type. It is the umId described in subclause 5.6.2.12. (NOTE)
	RefUmData []string `json:"refUmData,omitempty"`

	// A reference to the condition data. It is the condId described in subclause 5.6.2.9.
	RefCondData string `json:"refCondData,omitempty"`
}
