package level

import (
	"testing"

	"github.com/AlexanderHOtt/lsmtree/pkg/entry"
	"github.com/stretchr/testify/assert"
)

func TestLevelAppendFull(t *testing.T) {
	lvl := New(5)
	ents := make([]entry.Entry, 0, 5)
	for i := 0; i < 5; i++ {
		ents = append(ents, entry.New(5-i, 5-i))
	}

	ents = lvl.Append(ents)

	// no entries left
	assert.Nil(t, ents)

	// sorted run is sorted
	prevKey := 0
	for _, e := range lvl.sorted_run {
		assert.Greater(t, e.Key, prevKey)
		prevKey = e.Key
	}
}

func TestLevelAppendOverflow(t *testing.T) {
	lvl := New(5)
	ents := make([]entry.Entry, 0, 10)
	for i := 0; i < 10; i++ {
		ents = append(ents, entry.New(10-i, 10-i))
	}

	ents = lvl.Append(ents)

	assert.Equal(t, 10, len(ents))
	assert.Equal(t, 0, len(lvl.sorted_run))

	// sorted run is sorted
	prevKey := 0
	for _, e := range lvl.sorted_run {
		assert.Greater(t, e.Key, prevKey)
		prevKey = e.Key
	}
}

func TestLevelAppendOverflowWithPartiallyFilledSortedRun(t *testing.T) {
	lvl := New(5)

	ents := make([]entry.Entry, 0, 3)
	for i := 0; i < 3; i++ {
		ents = append(ents, entry.New(i, i))
	}
	ents = lvl.Append(ents)

	assert.Nil(t, ents)

	ents = make([]entry.Entry, 0, 10)
	for i := 0; i < 10; i++ {
		ents = append(ents, entry.New(10-i, 10-i))
	}

	ents = lvl.Append(ents)

	assert.Equal(t, 13, len(ents))
	assert.Equal(t, 0, len(lvl.sorted_run))

	// sorted run is sorted
	prevKey := 0
	for _, e := range lvl.sorted_run {
		assert.GreaterOrEqual(t, e.Key, prevKey)
		prevKey = e.Key
	}
}
