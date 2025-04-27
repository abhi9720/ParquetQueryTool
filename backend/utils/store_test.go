package utils

import (
	"apica-search/models"
	"testing"
)

func BenchmarkRecordStore_Set(b *testing.B) {

	InitStore()

	record := &models.Record{
		Message:        "Sample message",
		MessageRaw:     "Sample raw message",
		StructuredData: "{\"key\":\"value\"}",
		Tag:            "sample_tag",
		Sender:         "sender",
		Groupings:      "group1,group2",
		Event:          "event1",
		EventId:        "12345",
		NanoTimeStamp:  "1740755524740577500",
		Namespace:      "default_namespace",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Store.Set(record)
	}
}

func BenchmarkRecordStore_Search(b *testing.B) {

	InitStore()

	records := []models.Record{
		{Message: "Test message 1", Tag: "tag1", Sender: "sender1", Groupings: "group1", EventId: "event1", Namespace: "default_namespace"},
		{Message: "Test message 2", Tag: "tag2", Sender: "sender2", Groupings: "group2", EventId: "event2", Namespace: "default_namespace"},
		{Message: "Test message 3", Tag: "tag1", Sender: "sender1", Groupings: "group1", EventId: "event3", Namespace: "default_namespace"},
	}

	SetRecords(records)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Store.Search("tag1")
		if err != nil {
			b.Fatalf("Error searching for word: %v", err)
		}
	}
}
