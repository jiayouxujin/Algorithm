package main

type LRUCache struct {
	capacity int
	dict     map[int]int
	queue    []int
}

func Constructor2(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		dict:     make(map[int]int),
		queue:    make([]int, 0),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.dict[key]; ok {
		for i, k := range this.queue {
			if k == key {
				this.queue = append(this.queue[:i], this.queue[i+1:]...)
				break
			}
		}
		this.queue = append(this.queue, key)
		return v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.dict[key]; ok {
		for i, k := range this.queue {
			if k == key {
				this.queue = append(this.queue[:i], this.queue[i+1:]...)
				break
			}
		}
	}
	if len(this.queue) == this.capacity {
		deletedKey := this.queue[0]
		this.queue = this.queue[1:]
		delete(this.dict, deletedKey)
	}
	this.queue = append(this.queue, key)
	this.dict[key] = value
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
