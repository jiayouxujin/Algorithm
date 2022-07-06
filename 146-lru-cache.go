package main

type Node struct {
	Key, Value int
	Prev, Next *Node
}
type LRUCache struct {
	head, tail *Node
	Keys       map[int]*Node
	Cap        int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{Cap: capacity, Keys: make(map[int]*Node)}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Keys[key]; ok {
		node.Value = value
		this.Remove(node)
		this.Add(node)
		return
	} else {
		node := &Node{
			Key:   key,
			Value: value,
		}
		this.Keys[key] = node
		this.Add(node)
	}
	if len(this.Keys) > this.Cap {
		delete(this.Keys, this.tail.Key)
		this.Remove(this.tail)
	}
}

func (this *LRUCache) Add(node *Node) {
	node.Prev = nil
	node.Next = this.head
	if this.head != nil {
		this.head.Prev = node
	}
	this.head = node

	if this.tail == nil {
		this.tail = node
		this.tail.Next = nil
	}
}

func (this *LRUCache) Remove(node *Node) {
	if node == this.head {
		this.head = node.Next
		node.Next = nil
		return
	} else if node == this.tail {
		this.tail = node.Prev
		this.tail.Next = nil
		node.Prev = nil
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
