package Graph

import (
	"sync"
	"fmt"
)

type Node struct {
 	value int
 }

 type Graph struct {
	 nodes []*Node          // 节点集
	 edges map[Node][]*Node // 邻接表表示的无向图
	 lock  sync.RWMutex     // 保证线程安全
 }

// 增加节点
func (self *Graph) AddNode(n *Node) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.nodes = append(self.nodes, n)
}
 func(self *Graph) AddEdge(u,v *Node){
 	self.lock.Lock()
 	defer self.lock.Unlock()
 	if self.edges==nil{
 		self.edges=make(map[Node][]*Node)
	}
	self.edges[*u]=append(self.edges[*u],v) //无向图
	self.edges[*v]=append(self.edges[*v],u)
 }

func(self *Graph) String(){
	self.lock.Lock()
	defer self.lock.Unlock()
	for _,node:=range self.nodes{
		print("the value is:",node.value)
		nexts:=self.edges[*node]
		for _,nodenext:=range nexts {
			print("the next value:",nodenext.value)
		}
	}
}

func(self *Graph) BFS(f func(node *Node)){
	self.lock.Lock()
	defer self.lock.Unlock()
	if self.nodes==nil{
		panic("valid Graph")
	}
	q:=NewNodeQueue()
	visited:=make(map[Node]bool)
	head:=self.nodes[0]
	q.Enqueue(*head)
	for{
		if q.IsEmpty(){
			break
		}
		anode:=q.Dequeue()
		visited[*anode]=true
		if nexts,ok:=self.edges[*anode];!ok{
			for _,inode:=range nexts{
				if !visited[*inode]{
					q.Enqueue(*inode)
					visited[*inode]=true
				}else {
					continue
				}

			}
		}
		if f!=nil{
			f(anode)
		}
	}

}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}