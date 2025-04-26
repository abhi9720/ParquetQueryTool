package utils

import (
	"apica-search/models"
	"sync"
)

var (
	recordStore []models.Record
	storeLock   sync.RWMutex
)

func SetRecords(records []models.Record) {
	storeLock.Lock()
	defer storeLock.Unlock()
	recordStore = records
}

func GetRecords() []models.Record {
	storeLock.RLock()
	defer storeLock.RUnlock()
	return recordStore
}
