package config

import (
	"golang-question/errorx"
	"sync"
)

type Manager[T any] interface {
	Get() T
	Update(T) errorx.Error
	OnChange(func(T)) (cancel func())
	Watch() Manager[T]     //執行Watch後，會開始監聽配置的變化，並在變化時自動更新 否則每次Get都會從數據源取得最新資料
	InitData(T) Manager[T] //如果數據源沒有資料，則使用InitData put資料
}

func Local[T any]() Manager[T] {
	return &localManager[T]{
		mu: new(sync.RWMutex),
	}
}

func Etcd[T any]() Manager[T] {
	//TODO: implement
	return nil
}

type localManager[T any] struct {
	mu       *sync.RWMutex
	data     T
	onChange func(T)
	watching bool
}

func (lm *localManager[T]) Get() T {
	lm.mu.RLock()
	defer lm.mu.RUnlock()
	return lm.data
}

func (lm *localManager[T]) Update(newData T) errorx.Error {
	lm.mu.Lock()
	defer lm.mu.Unlock()
	lm.data = newData
	//如果监听状态就调用OnChange回调
	if lm.onChange != nil && lm.watching {
		lm.onChange(newData)
	}
	return nil
}

func (lm *localManager[T]) OnChange(callback func(T)) (cancel func()) {
	lm.mu.Lock()
	defer lm.mu.Unlock()
	lm.onChange = callback
	return func() {
		lm.mu.Lock()
		defer lm.mu.Unlock()
		lm.onChange = nil
	}
}

func (lm *localManager[T]) Watch() Manager[T] {
	lm.mu.Lock()
	defer lm.mu.Unlock()
	lm.watching = true
	return lm
}

func (lm *localManager[T]) InitData(initialData T) Manager[T] {
	lm.mu.Lock()
	defer lm.mu.Unlock()
	lm.data = initialData
	return lm
}
