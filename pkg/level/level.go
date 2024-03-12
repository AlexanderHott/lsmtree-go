package level

import (
	"slices"

	"github.com/AlexanderHOtt/lsmtree/pkg/entry"

	_logger "github.com/AlexanderHOtt/lsmtree/pkg/logger"
)

var logger = _logger.Logger

type Level struct {
	// filepath   string
	sorted_run []entry.Entry
	// sorted_run []Page
}

// type Page struct {
//   entries []entry.Entry
// }

func New(capacity int) Level {
	logger.Debug("Allocating new level")
	return Level{
		// filepath:   "./data/level1",
		sorted_run: make([]entry.Entry, 0, capacity),
	}
}

func (lvl *Level) Append(buf []entry.Entry) []entry.Entry {
	logger.Debugf("Appending buf sorted_run=%d/%d buf=%d %v %v lvl_addr=%p", len(lvl.sorted_run), cap(lvl.sorted_run), len(buf), buf, lvl.sorted_run, lvl.Buf())

	if lvl.Len()+len(buf) > lvl.Cap() {
		logger.Info("Buf too big for level")
		newBuf := append(lvl.sorted_run, buf...)
		slices.SortFunc(newBuf, func(a, b entry.Entry) int {
			return a.Key - b.Key
		})
		lvl.sorted_run = nil
		return newBuf
	}

	logger.Debug("pre append", "sorted_run", lvl.sorted_run)
	lvl.sorted_run = append(lvl.sorted_run, buf...)
	buf = nil

	logger.Debug("post append", "sorted_run", lvl.sorted_run)
	logger.Debug("sorted_run len pre sort", "len", len(lvl.sorted_run), "cap", cap(lvl.sorted_run), "sorted_run", lvl.sorted_run)
	// logger.Debug("", "sorted_run_len", len(lvl.sorted_run))
	slices.SortFunc(lvl.sorted_run, func(a, b entry.Entry) int {
		return a.Key - b.Key
	})
	logger.Debug("sorted_run len post sort", "len", len(lvl.sorted_run), "cap", cap(lvl.sorted_run), "sorted_run", lvl.sorted_run)

	// return lvl.sorted_run
	return nil
}

func (lvl *Level) Len() int {
	return len(lvl.sorted_run)
}

func (lvl *Level) Cap() int {
	return cap(lvl.sorted_run)
}

func (lvl *Level) Buf() []entry.Entry {
	return lvl.sorted_run
}

// func (l *Level) WriteBuf(buf []Entry) {
// 	var buffer bytes.Buffer
// 	encoder := gob.NewEncoder(&buffer)
// 	for _, e := range buf {
// 		err := encoder.Encode(e)
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			return
// 		}
// 	}
// 	gobData := buffer.Bytes()
// 	fmt.Println("encoded stuff", gobData)
// }

// // Deserialization (decoding)
// var newPerson Person
// decoder := gob.NewDecoder(bytes.NewBuffer(gobData))
// err = decoder.Decode(&newPerson)
// if err != nil {
//     fmt.Println("Error:", err)
//     return
// }
// fmt.Println(newPerson)
