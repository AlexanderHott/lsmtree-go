package lsmtree

import (
	"errors"
	"fmt"
	"time"
)

type KVStore interface {
	Put(int, int)
	Get(int)
	Delete(int)
}

type Entry struct {
	Key       int
	Value     int
	tombstone bool
	timestamp int64
}

func (e Entry) String() string {
	if !e.tombstone {
		return fmt.Sprintf("%d: %d [%d] | [%t]", e.Key, e.Value, e.timestamp, e.tombstone)
	}
	return "<deleted>"
}

type LSMTree struct {
	buf      []Entry
	size     int
	capacity int
}

func New(capacity int) *LSMTree {
	return &LSMTree{buf: make([]Entry, capacity), size: 0, capacity: capacity}
}

func (l *LSMTree) Put(key int, value int) {
	entry := Entry{Key: key, Value: value, timestamp: time.Now().UnixNano(), tombstone: false}
	l.put(&entry)
}

func (l *LSMTree) put(entry *Entry) {
	l.buf[l.size] = *entry
	l.size++
}

func (l *LSMTree) Get(key int) (int, error) {
	max_timestamp := int64(0)
	var newestEntry Entry

	// scan buffer
	for _, e := range l.buf {
		if e.Key == key && e.timestamp > max_timestamp {
			newestEntry = e
			max_timestamp = newestEntry.timestamp
		}
	}

	// if an entry exists, return it
	if newestEntry.timestamp > 0 && !newestEntry.tombstone {
		return newestEntry.Value, nil
	}
	return 0, errors.New("not found")
}

func (l *LSMTree) GetRange(start_key int, end_key int) map[int]int {
	res := map[int]int{}

	for _, e := range l.buf {
		if e.Key >= start_key && e.Key <= end_key {
			res[e.Key] = e.Value
		}
	}

	return res
}

func (l *LSMTree) Delete(key int) {
	e := Entry{Key: key, timestamp: time.Now().UnixNano(), tombstone: true}
	l.put(&e)
}
