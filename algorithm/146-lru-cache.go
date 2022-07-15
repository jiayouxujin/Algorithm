package main

//采用双向链表的方式，这样可以保证时间复杂度为O(1)
type Node struct {
	Key, Val   int
	Prev, Next *Node
}

type LRUCache struct {
	head, tail *Node
	capacity   int
	keys       map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		keys:     make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.keys[key]; ok {
		this.Delete(node)
		this.Add(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.keys[key]; ok {
		node.Val = value
		this.Delete(node)
		this.Add(node)
		return
	} else {
		node := &Node{Key: key, Val: value}
		this.Add(node)
		this.keys[key] = node
	}

	if len(this.keys) > this.capacity {
		delete(this.keys, this.tail.Key)
		this.Delete(this.tail)
	}
}

func (this *LRUCache) Delete(node *Node) {
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
