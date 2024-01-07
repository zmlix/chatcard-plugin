package main

type PluginConfigure struct {
	Version string    `json:"version,omitempty"`
	Auther  string    `json:"auther,omitempty"`
	Display string    `json:"display,omitempty"`
	Img     string    `json:"img,omitempty"`
	Name    string    `json:"name"`
	Cmd     string    `json:"cmd,omitempty"`
	Plugins []Plugins `json:"plugins"`
	Options []Options `json:"options,omitempty"`
}

type Function struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Parameters  any    `json:"parameters"`
}
type Plugins struct {
	Type     string   `json:"type"`
	Function Function `json:"function"`
}
type Options struct {
	Type  int32  `json:"type,omitempty"`
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
