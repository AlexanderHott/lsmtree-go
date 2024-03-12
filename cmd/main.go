package main

import (
	"github.com/AlexanderHOtt/lsmtree/pkg/lsmtree"
)

func main() {
	lsm := lsmtree.New(5)
	for i := 0; i < 5000; i++ {
		lsm.Put(i, i)
	}

	lsm.Put(100, 100)

}
