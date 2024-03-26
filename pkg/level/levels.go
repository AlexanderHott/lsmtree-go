package level

import (
	"github.com/AlexanderHOtt/lsmtree/pkg/config"
	"github.com/AlexanderHOtt/lsmtree/pkg/entry"
)

type Levels struct {
	levels      []Level
	lsmCapacity int
}

func NewLevels(lsmCapacity int) Levels {
	return Levels{
		levels:      make([]Level, 0), // biggest level at 0, smallest at len-1
		lsmCapacity: lsmCapacity,
	}
}

func (lvls *Levels) Append(buf *[]entry.Entry) {
	// naive iterative compaction

	// check if there is space left in the whole tree
	logger.Debug("append buf", "lvllen", lvls.LvlLen(), "lvlcap", lvls.LvlCap())
	if lvls.LvlLen()+len(*buf) > lvls.LvlCap() {
		logger.Info("tree full, allocating new level", "lvllen", lvls.LvlLen(), "lvlcap", lvls.LvlCap())
		for i := range lvls.levels {
			// https://go.dev/wiki/SliceTricks
			// Extend Capacity
			lvl := &lvls.levels[i]
			lvl.sorted_run = append(
				make(
					[]entry.Entry,
					0,
					lvl.Cap()*config.Cfg.ScaleFactor,
				),
				lvl.sorted_run...,
			)
		}

		// insert new level 1 at the end
		lvls.levels = append(lvls.levels, New(lvls.lsmCapacity*config.Cfg.ScaleFactor))
	}

	// append buf to correct page
	toAppend := *buf
	for i := len(lvls.levels) - 1; i >= 0; i-- {
		lvl := &lvls.levels[i]
		toAppend = lvl.Append(toAppend) // TODO: fix append logic for full levels. they dont return the correct value
		// case 1: buf can fit into level -> just append
		// case 2: buf cant fit into level -> returns levels sorted run + buf as a sorted array
		//if len(toAppend)+lvl.Len() < lvl.Cap() {
		//	// can fit, just append
		//	lvl.sorted_run = append(lvl.sorted_run, toAppend...)
		//} else {
		//	// can't fit. compact level and append to next down
		//	toAppend, _ = lvl.Append(toAppend)
		//	toAppend = append(lvl.sorted_run, toAppend...)
		//	lvl.sorted_run = lvl.sorted_run[0:]
		//}

		if toAppend == nil {
			break
		}
	}
	if toAppend != nil {
		panic("Buf wasn't appended")
	}
}

// Total capacity of all the levels
func (lvls *Levels) LvlCap() int {
	total := 0
	for _, lvl := range lvls.levels {
		total += lvl.Cap()
	}
	return total
}

// Total length of all the levels
func (lvls *Levels) LvlLen() int {
	total := 0
	logger.Warn("lvl len", "len", len(lvls.levels))
	for _, lvl := range lvls.levels {
		total += lvl.Len()
	}
	return total
}

// Number of levels
func (lvls *Levels) Len() int {
	return len(lvls.levels)
}

func (lvls *Levels) Get(key int) (int, error) {
	return 0, nil
}
