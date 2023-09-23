package models

type Response struct {
	Action string `json:"action"`
	Value  string `json:"value"`
}

type SendData struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type View struct {
	IP      string `json:"ip"`
	Country string `json:"cc"`
}
