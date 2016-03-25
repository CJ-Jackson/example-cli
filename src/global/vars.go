package global

import (
	"sync"
)

var (
	global     = &Global{}
	globalSync sync.RWMutex
)
