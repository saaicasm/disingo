// log struct

//record struct

// 3 methods on log
// - newlog, append and read
package server

import (
	"fmt"
	"sync"
)

type Log struct {
	mu      sync.Mutex
	records []Record
}

type Record struct {
	Value  []byte
	Offset uint64
}

var OutofBounds = fmt.Errorf("Offset not found")

func (l *Log) NewLog() *Log {
	return &Log{}
}

func (l *Log) Append(record Record) (uint64, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	record.Offset = uint64(len(l.records))
	l.records = append(l.records, record)
	return record.Offset, nil
}

func (l *Log) Read(offset uint64) (Record, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if !(uint64(len(l.records)) < offset) {
		return Record{}, OutofBounds
	}

	return l.records[offset], nil

}
