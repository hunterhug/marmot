/*
	版权所有，侵权必究
	署名-非商业性使用-禁止演绎 4.0 国际
	警告： 以下的代码版权归属hunterhug，请不要传播或修改代码
	你可以在教育用途下使用该代码，但是禁止公司或个人用于商业用途(在未授权情况下不得用于盈利)
	商业授权请联系邮箱：gdccmcm14@live.com QQ:459527502

	All right reserved
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
	For more information on commercial licensing please contact hunterhug.
	Ask for commercial licensing please contact Mail:gdccmcm14@live.com Or QQ:459527502
*
*/

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
