package utils

import (
	"apica-search/models"
	"regexp"
	"strings"
	"sync"
)

type RecordStore struct {
	records []*models.Record
	index   map[string][]*models.Record
	mu      sync.RWMutex
}

var Store *RecordStore

var once sync.Once

func InitStore() {
	once.Do(func() {
		Store = NewRecordStore()
	})
}

func NewRecordStore() *RecordStore {
	return &RecordStore{
		index: make(map[string][]*models.Record),
	}
}

func SetRecords(records []models.Record) {
	ptrs := make([]*models.Record, 0, len(records))
	for i := range records {
		ptrs = append(ptrs, &records[i])
	}
	Store.SetAll(ptrs)
}

func (store *RecordStore) SetAll(records []*models.Record) {
	for _, record := range records {
		store.Set(record)
	}
}

func (store *RecordStore) Set(record *models.Record) {
	store.mu.Lock()
	defer store.mu.Unlock()

	store.records = append(store.records, record)

	for _, field := range store.getFields(record) {
		if field == "" {
			continue
		}
		words := splitWords(field)
		uniqueWords := make(map[string]struct{})
		for _, word := range words {
			if _, exists := uniqueWords[word]; exists {
				continue
			}
			uniqueWords[word] = struct{}{}
			store.index[word] = append(store.index[word], record)
		}
	}
}

func (store *RecordStore) Search(word string) ([]*models.Record, error) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	if store.index == nil {
		return nil, ErrSearchFailed
	}

	records, ok := store.index[word]
	if !ok {
		return nil, ErrRecordNotFound
	}

	return records, nil
}

func (store *RecordStore) getFields(record *models.Record) []string {
	return []string{
		record.Message,
		record.MessageRaw,
		record.StructuredData,
		record.Tag,
		record.Sender,
		record.Groupings,
		record.Event,
		record.EventId,
		record.NanoTimeStamp,
		record.Namespace,
	}
}

var wordSplitter = regexp.MustCompile(`[a-zA-Z0-9]+`)

func splitWords(s string) []string {
	s = strings.ToLower(s)
	return wordSplitter.FindAllString(s, -1)
}

func (store *RecordStore) Clear() {
	store.mu.Lock()
	defer store.mu.Unlock()
	store.records = nil
	store.index = make(map[string][]*models.Record)
}
