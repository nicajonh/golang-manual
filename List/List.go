package List

type Node struct {
	Key        string // 淘汰 tail 时需在维护的哈希表中删除，不是冗余存储
	Val        interface{}
	Freq int// 根据使用频率，将节点从旧梯队移除时使用，非冗余存储
	Prev, Next *Node // 双向指针
}
type List struct{
	head,tail *Node
	size int// size of cache
}

func  NewList() *List{
	return &List{
		head:nil,
		tail:nil,
		size:0,
	}
}
func (l *List) Prepend(node *Node) *Node{
	if l.head==nil{
		l.head=node
		l.tail=node
	}else {
		node.Prev = nil
		node.Next = l.head
		l.head.Prev = node
		l.head = node
	}
	l.size++
	return node
}

func(l *List) Remove(node *Node) *Node{
	if node == nil {
		return nil
	}
	prev, next := node.Prev, node.Next
	if prev == nil {
		l.head = next // 删除头结点
	} else {
		prev.Next = next
	}

	if next == nil {
		l.tail = prev // 删除尾结点
	} else {
		next.Prev = prev
	}

	l.size--
	node.Prev, node.Next = nil, nil
	return node
}

// 封装数据已存在缓存的后续操作
func (l *List) MoveToHead(node *Node) *Node {
	if node == nil {
		return nil
	}
	n := l.Remove(node)
	return l.Prepend(n)
}

func (l *List) Tail() *Node{
	return l.tail
}

func (l *List) Size() int{
	return l.size
}



