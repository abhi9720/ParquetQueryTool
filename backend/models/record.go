package models

type Record struct {
	Message        string `json:"message"`
	MessageRaw     string `json:"message_raw"`
	StructuredData string `json:"structured_data"`
	Tag            string `json:"tag"`
	Sender         string `json:"sender"`
	Groupings      string `json:"groupings"`
	Event          string `json:"event"`
	EventId        string `json:"event_id"`
	NanoTimestamp  int64  `json:"nano_timestamp"`
	Namespace      string `json:"namespace"`
}
