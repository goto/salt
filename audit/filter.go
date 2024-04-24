package audit

import "time"

type Filter struct {
	Actor     string
	Action    string
	Data      map[string]string
	Metadata  map[string]string
	StartTime time.Time
	EndTime   time.Time
	Limit     int32
	Page      int32
}
