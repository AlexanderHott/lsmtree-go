package lsmtree

import (
	// "bytes"
	// "encoding/gob"
	"errors"
	"time"

	"github.com/AlexanderHOtt/lsmtree/pkg/config"
	"github.com/AlexanderHOtt/lsmtree/pkg/entry"
	"github.com/AlexanderHOtt/lsmtree/pkg/level"
	_logger "github.com/AlexanderHOtt/lsmtree/pkg/logger"
	// "github.com/charmbracelet/log"
)

var logger = _logger.Logger

type KVStore interface {
	Put(int, int)
	Get(int)
	Delete(int)
}

type LSMTree struct {
	buf []entry.Entry
	// len          int
	// capacity     int
	// levels      []level.Level
	levels      level.Levels
	scaleFactor int
}

func (lsm *LSMTree) Len() int {
	return len(lsm.buf)
}

func (lsm *LSMTree) Cap() int {
	return cap(lsm.buf)
}

// New LSMTree
func New(capacity int) *LSMTree {
	return &LSMTree{
		buf: make([]entry.Entry, 0, capacity),
		// levels: []level.Level{
		// 	level.New(capacity * config.Cfg.ScaleFactor),
		// },
		levels:      level.NewLevels(capacity),
		scaleFactor: config.Cfg.ScaleFactor}
}

func (lsm *LSMTree) Put(key int, value int) {
	ent := entry.New(key, value)
	logger.Info("Putting", "kv", ent)
	lsm.put(&ent)
}

func (lsm *LSMTree) put(ent *entry.Entry) {
	lsm.buf = append(lsm.buf, *ent)
	logger.Debugf("mem buf len/cap %d/%d", lsm.Len(), lsm.Cap())

	// compaction
	if lsm.Len() >= lsm.Cap() {
		logger.Info("mem buf full, compacting level", "buf_len", len(lsm.buf))
		lsm.levels.Append(&lsm.buf)

		logger.Debug("replacing mem buf", "cap", lsm.Cap())
		//lsm.buf = make([]entry.Entry, 0, lsm.Cap())
		lsm.buf = lsm.buf[:0]
	}
}

func (lsm *LSMTree) Get(key int) (int, error) {
	maxTimestamp := int64(0)
	var newestEntry entry.Entry

	// scan buffer
	for _, e := range lsm.buf {
		if e.Key == key && e.Timestamp > maxTimestamp {
			newestEntry = e
			maxTimestamp = newestEntry.Timestamp
		}
	}

	// if an entry exists, return it
	if newestEntry.Timestamp > 0 && !newestEntry.Tombstone {
		return newestEntry.Value, nil
	}
	return 0, errors.New("not found")
}

func (lsm *LSMTree) GetRange(start_key int, end_key int) map[int]int {
	res := map[int]int{}

	for _, e := range lsm.buf {
		if e.Key >= start_key && e.Key <= end_key {
			res[e.Key] = e.Value
		}
	}

	return res
}

func (lsm *LSMTree) Delete(key int) {
	e := entry.Entry{Key: key, Timestamp: time.Now().UnixNano(), Tombstone: true}
	lsm.put(&e)
}
