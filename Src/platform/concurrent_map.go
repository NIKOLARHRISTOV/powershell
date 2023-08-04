package platform

import "sync"

func NewConcurrentMap() *ConcurrentMap {
	var cm ConcurrentMap
	return &cm
}

type ConcurrentMap sync.Map

func (cm *ConcurrentMap) Set(key string, value any) {
	(*sync.Map)(cm).Store(key, value)
}

<<<<<<< HEAD:Src/platform/concurrent_map.go
func (c *ConcurrentMap) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	if val, ok := c.values[key]; ok {
		return val, true
	}

	return "", false
=======
func (cm *ConcurrentMap) Get(key string) (any, bool) {
	return (*sync.Map)(cm).Load(key)
>>>>>>> upstream/main:src/platform/concurrent_map.go
}

func (cm *ConcurrentMap) Delete(key string) {
	(*sync.Map)(cm).Delete(key)
}

func (cm *ConcurrentMap) Contains(key string) bool {
	_, ok := (*sync.Map)(cm).Load(key)
	return ok
}

func (cm *ConcurrentMap) List() map[string]any {
	list := make(map[string]any)
	(*sync.Map)(cm).Range(func(key, value any) bool {
		list[key.(string)] = value
		return true
	})
	return list
}
