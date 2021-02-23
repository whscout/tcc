package model

// Nodes request services
type Nodes struct {
	Url     string `json:"url"`
	Method  string `json:"method"`
	Timeout int    `json:"timeout"`
}


