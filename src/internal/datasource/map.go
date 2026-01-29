package datasource

import (
	"sync"
)

type Map[K comparable, V any] struct {
	m sync.Map
}
