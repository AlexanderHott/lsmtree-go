package lsmtree

import (
	"testing"
)

func TestNewLSMTree(t *testing.T) {
	capacity := 20
	tree := New(capacity)
	if cap(tree.buf) != capacity || tree.capacity != capacity {
		t.Errorf("Incorrect capacity, expected %d, got %d", capacity, cap(tree.buf))
	}
	if tree.size != 0 {
		t.Errorf("Incorrect size, expected 0, not %d", tree.size)
	}
}

func TestPut(t *testing.T) {
	capacity := 20
	tree := New(capacity)
	tree.Put(1, 1)
	if tree.buf[0].Key != 1 {
		t.Errorf("Invalid put")
	}
}

func TestUpdate(t *testing.T) {
	capacity := 20
	tree := New(capacity)
	tree.Put(1, 1)
	tree.Put(1, 2)

	v, e := tree.Get(1)
	if e != nil {
		t.Errorf("error updating %s", e)
	}
	if v != 2 {
		t.Errorf("Updated failed. Expected %d, got %d", 2, v)
	}
}

func TestDelete(t *testing.T) {
	capacity := 20
	tree := New(capacity)
	tree.Put(1, 1)
	tree.Delete(1)

	// expect error
	v, err := tree.Get(1)
	if err == nil {
		t.Errorf("Error deleting struct, got %d", v)
	}
}
