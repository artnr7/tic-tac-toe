package datasource

import (
	"sync"
)

type Storage[K comparable, V any] struct {
	m sync.Map
}
