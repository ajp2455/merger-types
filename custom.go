package merger

type Payload struct {
	Info       *Params                `json:"info,omitempty"`
	Clients    []*Client              `json:"clients,omitempty"`
	ScriptData map[string]interface{} `json:"script,omitempty"`
}

func (w *Window) IsInsidePeriod(low, high int64) bool {
	return w.Start <= high && w.End >= low
}
