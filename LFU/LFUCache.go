package LFU

import (
	aList "../List"
)

type LFUCache struct {
	capacity int // 缓存空间大小
	minfreq int// 低频率
	items    map[string]*aList.Node
	freqs map[int]*aList.List // 不同频率梯队
}

func NewLFUCache(capacity int) *LFUCache{
	return &LFUCache{
		capacity:capacity,
		minfreq:0,
		items:make(map[string]*aList.Node),
		freqs:make(map[int]*aList.List),
	}
}

func (lfu *LFUCache) Set(k string,v interface{}){
	//命中
	if lfu.capacity <= 0 {
		return
	}

	// 命中，需要更新频率
	if val := lfu.Get(k); val != -1 {
		lfu.items[k].val = v // 直接更新值即可
		return
	}

	node := &aList.Node{key: k, val: v, freq: 1}

	// 未命中
	// 缓存已满
	if lfu.capacity == len(lfu.items) {
		old := lfu.freqs[lfu.minfreq].Tail() // 最低最旧
		lfu.freqs[lfu.minfreq].Remove(old)
		delete(lfu.items, old.key)
	}

	// 缓存未满，放入第 1 梯队
	c.items[k] = node
	if _, ok := lfu.freqs[1]; !ok {
		lfu.freqs[1] = aList.NewList()
	}
	lfu.freqs[1].Prepend(node)
	lfu.minfreq = 1
}
func (lfu *LFUCache) Get(k string) interface{}{
	node, ok := lfu.items[k]
	if !ok {
		return -1
	}

	// 移到 +1 梯队中
	lfu.freqs[node.freq].Remove(node)
	node.freq++
	if _, ok := lfu.freqs[node.freq]; !ok {
		lfu.freqs[node.freq] = NewList()
	}
	newNode := lfu.freqs[node.freq].Prepend(node)
	lfu.items[k] = newNode // 新地址更新到 map
	if lfu.freqs[lfu.minfreq].Size() == 0 {
		lfu.minfreq++ // Get 的正好是当前值
	}
	return newNode.val
}


