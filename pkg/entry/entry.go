package entry

import (
	"fmt"
	"time"

  // _logger "github.com/AlexanderHOtt/lsmtree/pkg/logger"
)


type Entry struct {
	Key       int
	Value     int
	Tombstone bool
	Timestamp int64
}

func (e Entry) String() string {
	if !e.Tombstone {
		return fmt.Sprintf("%d:%d", e.Key, e.Value)
	}
	return "<deleted>"
}

func (e Entry) Debug() string {
	if !e.Tombstone {
		return fmt.Sprintf("%d: %d [%d] | [%t]", e.Key, e.Value, e.Timestamp, e.Tombstone)
	}
	return "<deleted>"
}

func New(key, value int) Entry {
	return Entry{
		Key:       key,
		Value:     value,
		Tombstone: false,
		Timestamp: time.Now().UnixNano(),
	}
}
