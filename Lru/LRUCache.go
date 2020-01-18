package Lru

import (
	"../List"
)

type LRUCache struct {
	capacity int // 缓存空间大小
	items    map[string]*List.Node
	list     *List.List
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
	capacity: capacity,
	items:    make(map[string]*List.Node),
	list:     new(List.List),
	}
}

func (c *LRUCache) Set(k string, v interface{}) {
	// 命中
	if node, ok := c.items[k]; ok {
	node.Val = v                         // 命中后更新值
	c.items[k] = c.list.MoveToHead(node) //
	return
	}

	// 未命中
	node := &List.Node{Key: k, Val: v} // 完整的 node
	if c.capacity == c.list.Size() {
	tail := c.list.Tail()
	delete(c.items, tail.Key) // k-v 数据存储与 node 中
	c.list.Remove(tail)
	}
	c.items[k] = c.list.Prepend(node) // 更新地址
}

func (c *LRUCache) Get(k string) interface{} {
	node, ok := c.items[k]
	if ok {
	c.items[k] = c.list.MoveToHead(node)
	return node.Val
	}
	return -1
}