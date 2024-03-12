package level

import (
	"testing"

	"github.com/AlexanderHOtt/lsmtree/pkg/entry"
	"github.com/stretchr/testify/assert"
)

func TestLevelsAppend(t *testing.T) {
	lvls := NewLevels(5)

	ents := make([]entry.Entry, 0, 5)
	for i := 0; i < 5; i++ {
		ents = append(ents, entry.New(i, i))
	}

	lvls.Append(&ents)

	assert.Equal(t, 1, lvls.Len())
	assert.Equal(t, 5, lvls.LvlLen())
	assert.Equal(t, 50, lvls.LvlCap())
}

func TestLevelsAppendALot(t *testing.T) {
	lvls := NewLevels(5)

	ents := make([]entry.Entry, 0, 5)
	for i := 0; i < 3000; i++ {
		ents = append(ents, entry.New(i, i))
		if len(ents) == 5 {
			lvls.Append(&ents)
			ents = make([]entry.Entry, 0, 5)
		}
	}
}
