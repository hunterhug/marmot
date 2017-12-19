package miner

import (
	"sync"
)

// Pool for many Worker, every Worker can only serial execution
var Pool = &_Workers{ws: make(map[string]*Worker)}

type _Workers struct {
	mux sync.RWMutex // simple lock
	ws  map[string]*Worker
}

func (pool *_Workers) Get(name string) (b *Worker, ok bool) {
	pool.mux.RLock()
	b, ok = pool.ws[name]
	pool.mux.RUnlock()
	return
}

func (pool *_Workers) Set(name string, b *Worker) {
	pool.mux.Lock()
	pool.ws[name] = b
	pool.mux.Unlock()
	return
}

func (pool *_Workers) Delete(name string) {
	pool.mux.Lock()
	delete(pool.ws, name)
	pool.mux.Unlock()
	return
}
