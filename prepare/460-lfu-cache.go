package main

import "runtime/debug"

func init() { debug.SetGCPercent(-1) }

type Node struct {
	Key       interface{}
	Value     interface{}
	pre, next *Node
	list      *Dll // 该节点属于哪个双向链表
}

type Dll struct {
	head, tail *Node
	pre, next  *Dll // Dll作为二维双向链表结点，用于LFU
	cnt        int  // 该双向链表中的结点被访问的次数，用于LFU
	len        int
}

type Tdll struct {
	head, tail *Dll
}

func NewDll() *Dll {
	l := new(Dll)
	head, tail := new(Node), new(Node)
	head.list, tail.list = l, l
	head.next = tail
	tail.pre = head
	l.head = head
	l.tail = tail
	return l
}

func (l *Dll) moveToEnd(node *Node) {
	pre, next := node.pre, node.next
	pre.next = next
	next.pre = pre
	last := l.tail.pre
	last.next = node
	node.next = l.tail
	node.pre = last
	l.tail.pre = node
}

func (l *Dll) addNewNodeToEnd(node *Node) {
	last := l.tail.pre
	last.next = node
	l.tail.pre = node
	node.pre = last
	node.next = l.tail
	node.list = l
	l.len++
}

func (l *Dll) delFirst() *Node {
	first := l.head.next
	l.head.next = first.next
	first.next.pre = l.head
	first.pre, first.next = nil, nil
	l.len--
	return first
}

func NewTdll() *Tdll {
	tdll := &Tdll{}
	head, tail := new(Dll), new(Dll)
	head.next = tail
	tail.pre = head
	tdll.head, tdll.tail = head, tail
	return tdll
}

type LFUCache struct {
	cap, len int
	m        map[int]*Node
	list     *Tdll
}

func Constructor(capacity int) LFUCache {
	c := LFUCache{
		cap:  capacity,
		len:  0,
		m:    make(map[int]*Node),
		list: NewTdll(),
	}
	return c
}

func (cache *LFUCache) Get(key int) int {
	if v, ok := cache.m[key]; ok {
		// TODO: update cnt
		cache.updateCnt(v)
		return v.Value.(int)
	} else {
		return -1
	}
}

func (cache *LFUCache) Put(key int, value int) {
	if v, ok := cache.m[key]; ok { // key 已经存在于缓存中
		v.Value = value
		// TODO: update cnt
		cache.updateCnt(v)
		return
	}
	if cache.cap == 0 {
		return
	}
	// key 不存在于缓存中
	newNode := &Node{
		Key:   key,
		Value: value,
		pre:   nil,
		next:  nil,
		list:  nil,
	}
	if cache.len == cache.cap {
		// TODO: delete the first node of min cnt Dll
		firstDll := cache.list.head.next
		deletedNode := firstDll.delFirst()
		delete(cache.m, deletedNode.Key.(int))
		cache.len--
		// check and delete empty dll
		if firstDll.head.next == firstDll.tail {
			cache.list.head.next = firstDll.next
			firstDll.next.pre = cache.list.head
			firstDll.pre, firstDll.next = nil, nil
		}
	}
	// TODO: add new cache to end of Dll which cnt == 1
	if cache.list.head.next == cache.list.tail || cache.list.head.next.cnt != 1 { // need new Dll
		newDll := NewDll()
		newDll.cnt = 1
		head, secondDll := cache.list.head, cache.list.head.next
		head.next = newDll
		newDll.pre = head
		newDll.next = secondDll
		secondDll.pre = newDll
		newDll.addNewNodeToEnd(newNode)
	} else if cache.list.head.next.cnt == 1 { // do not need new dll
		cache.list.head.next.addNewNodeToEnd(newNode)
	}
	cache.m[key] = newNode
	cache.len++
}

func (cache *LFUCache) updateCnt(node *Node) {
	dll := node.list
	cnt := dll.cnt
	nextDll := dll.next
	if nextDll == cache.list.tail || nextDll.cnt != cnt+1 {
		newDll := NewDll()
		// insert new Dll between dll and nextDll
		dll.next = newDll
		newDll.pre = dll
		newDll.next = nextDll
		nextDll.pre = newDll
		// update newDll's attr
		newDll.cnt = cnt + 1
		// move node from dll to newDll
		preNode, nextNode := node.pre, node.next
		preNode.next = nextNode
		nextNode.pre = preNode
		newDll.addNewNodeToEnd(node)
		// if len of dll equals 0, delete dll from tdll
		if preNode == dll.head && nextNode == dll.tail {
			dll.pre.next = dll.next
			dll.next.pre = dll.pre
			dll.pre, dll.next = nil, nil
		}
	} else if nextDll.cnt == cnt+1 {
		// move node from dll to nextDll
		preNode, nextNode := node.pre, node.next
		preNode.next = nextNode
		nextNode.pre = preNode
		nextDll.addNewNodeToEnd(node)
		// if len of dll equals 0, delete dll from tdll
		if preNode == dll.head && nextNode == dll.tail {
			dll.pre.next = dll.next
			dll.next.pre = dll.pre
			dll.pre, dll.next = nil, nil
		}
	}
}
