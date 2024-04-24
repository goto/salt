package audit

import (
	"time"
)

type Log struct {
	Timestamp time.Time `json:"timestamp"`
	Action    string    `json:"action"`
	Actor     string    `json:"actor"`
	Data      any       `json:"data"`
	Metadata  any       `json:"metadata"`
}

type PagedLog struct {
	Count int32
	Logs  []Log
}
