package singleton

import (
	"fmt"
	"sync"
)

var (
	once           sync.Once
	globalResource *resource
)

type resource struct {
}

func GetResource() *resource {
	once.Do(func() {
		globalResource = new(resource)
	})
	return globalResource
	fmt.Println()
}
