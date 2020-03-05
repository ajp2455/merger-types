package merger

type Payload struct {
	Info       *Params                `json:"info,omitempty"`
	Clients    []*Client              `json:"clients,omitempty"`
	ScriptData map[string]interface{} `json:"script,omitempty"`
}
