package record

import "time"

type Record struct {
	ID        string    `json:"id"`
	Client    string    `json:"client"`    // ex: 192.168.128.1:63678
	Content   string    `json:"content"`   // ex: {"level":"info","cmd":"send","time":1665345472,"caller":"/Users/bmika/go/src/github.com/bartmika/logfarm-backend/cmd/send.go:40","message":"This is a test message"}
	Facility  int       `json:"facility"`  // ex: 0
	Hostname  string    `json:"hostname"`  // ex: MacMini2022.lan
	Priority  int       `json:"priority"`  // ex: 7
	Severity  int       `json:"severity"`  // ex: 7
	Tag       string    `json:"tag"`       // ex: logfarm-backend
	Timestamp time.Time `json:"timestamp"` // ex: 2022-10-09 15:57:52 -0400 -0400
	TLSPeer   string    `json:"tls_peer"`  // ex:
}

type RecordFilter struct {
	Tag                         string    `json:"tag"`
	SortOrder                   string    `json:"sort_order"`
	SortField                   string    `json:"sort_field"`
	IDs                         []uint64  `json:"ids"`
	TimestampGreaterThen        time.Time `json:"timestamp_gt,omitempty"`
	TimestampGreaterThenOrEqual time.Time `json:"timestamp_gte,omitempty"`
	TimestampLessThen           time.Time `json:"timestamp_lt,omitempty"`
	TimestampLessThenOrEqual    time.Time `json:"timestamp_lte,omitempty"`
}
