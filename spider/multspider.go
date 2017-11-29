/*
Copyright 2017 by GoSpider author. Email: gdccmcm14@live.com
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package spider

import (
	"sync"
)

var (
	// Pool for many spider, every spider can only serial execution
	Pool = &_Spider{sps: make(map[string]*Spider)}
)

type _Spider struct {
	mux sync.RWMutex // simple lock
	sps map[string]*Spider
}

func (sb *_Spider) Get(name string) (b *Spider, ok bool) {
	sb.mux.RLock()
	b, ok = sb.sps[name]
	sb.mux.RUnlock()
	return
}

func (sb *_Spider) Set(name string, b *Spider) {
	sb.mux.Lock()
	sb.sps[name] = b
	sb.mux.Unlock()
	return
}

func (sb *_Spider) Delete(name string) {
	sb.mux.Lock()
	delete(sb.sps, name)
	sb.mux.Unlock()
	return
}
