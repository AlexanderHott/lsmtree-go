package lsmtree

import (
	"testing"
)

func TestNewLSMTree(t *testing.T) {
	capacity := 20
	lsm := New(capacity)
	if cap(lsm.buf) != capacity || lsm.Cap() != capacity {
		t.Errorf("Incorrect capacity, expected %d, got %d", capacity, cap(lsm.buf))
	}
	if lsm.Len() != 0 {
		t.Errorf("Incorrect size, expected 0, not %d", lsm.Len())
	}
}

func TestPut(t *testing.T) {
	capacity := 20
	lsm := New(capacity)
	lsm.Put(1, 1)
	if lsm.buf[0].Key != 1 {
		t.Errorf("Invalid put")
	}
  val, err := lsm.Get(1)
  if val != 1 || err !=  nil {
    t.Errorf("Invalid put key=%d err=%s", val, err)
  }
}

func TestUpdate(t *testing.T) {
	capacity := 20
	lsm := New(capacity)
	lsm.Put(1, 1)
	lsm.Put(1, 2)

	val, err := lsm.Get(1)
	if err != nil {
		t.Errorf("error updating %s", err)
	}
	if val != 2 {
		t.Errorf("Updated failed. Expected %d, got %d", 2, val)
	}
}

func TestDelete(t *testing.T) {
	capacity := 20
	lsm := New(capacity)
	lsm.Put(1, 1)
	lsm.Delete(1)

	// expect error
	val, err := lsm.Get(1)
	if err == nil {
		t.Errorf("Error deleting struct, got %d", val)
	}
}
