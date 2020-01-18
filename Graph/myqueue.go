package Graph

import "sync"

type NodeQueue struct{
	nodes []Node
	lock sync.RWMutex
}

//new a Node Queue
func NewNodeQueue() *NodeQueue{
	q:=NodeQueue{}
	q.lock.Lock()
	defer q.lock.Unlock()
	q.nodes=[]Node{}
	return &q
}

//enqueue
func(self *NodeQueue) Enqueue(node Node){
	self.lock.Lock()
	defer self.lock.Unlock()
	self.nodes=append(self.nodes,node)
}

//dequeue
func(self *NodeQueue) Dequeue() *Node{
	self.lock.Lock()
	defer self.lock.Unlock()
	node:=self.nodes[0]
	self.nodes=self.nodes[1:]
	return &node
}
//is empty
func(self *NodeQueue) IsEmpty()bool{
	self.lock.Lock()
	defer self.lock.Unlock()
	return len(self.nodes)==0
}