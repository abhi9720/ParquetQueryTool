package models

type Record struct {
	Message        string `parquet:"name=Message, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" json:"message"`
	MessageRaw     string `parquet:"name=MessageRaw, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" json:"message_raw"`
	StructuredData string `parquet:"name=StructuredData, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" json:"structured_data"`
	Tag            string `parquet:"name=Tag, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" json:"tag"`
	Sender         string `parquet:"name=Sender, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" json:"sender"`
	Groupings      string `parquet:"name=Groupings, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" json:"groupings"`
	Event          string `parquet:"name=Event, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" json:"event"`
	EventId        string `parquet:"name=EventId, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" json:"event_id"`
	NanoTimeStamp  string `parquet:"name=NanoTimeStamp, type=INT64" json:"nano_timestamp"`
	Namespace      string `parquet:"name=namespace, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" json:"namespace"`
}
