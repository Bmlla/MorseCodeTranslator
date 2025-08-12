package entities

type Dictionary struct {
	From map[string]string `json:"from"`
	To   map[string]string `json:"to"`
}
