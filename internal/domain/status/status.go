package status

import "time"

type Snapshot struct {
	Status       string            `json:"status"`
	Service      string            `json:"service"`
	Environment  string            `json:"environment"`
	Timestamp    time.Time         `json:"timestamp"`
	Dependencies map[string]string `json:"dependencies"`
}
