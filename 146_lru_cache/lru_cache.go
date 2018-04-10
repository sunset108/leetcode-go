type Node struct {
	Key   int
	Value int
	Next  *Node
	Pre   *Node
}

type LRUCache struct {
	Cap   int
	Cache map[int]*Node
	Head  *Node
	Tail  *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:   capacity,
		Cache: make(map[int]*Node, 0),
		Head:  &Node{},
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.Cache[key]; ok {
		this.moveToTail(v)
		return v.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if len(this.Cache) == 0 {
		tmp := &Node{Key: key, Value: value}
		this.Cache[key] = tmp
		this.Head = tmp
		this.Tail = tmp
	} else if this.Get(key) == -1 {
		tmp := &Node{Key: key, Value: value}
		this.addToTail(tmp)
		this.Cache[key] = tmp
		if len(this.Cache) > this.Cap {
			delete(this.Cache, this.Head.Key)
			this.Head = this.Head.Next
			this.Head.Pre = nil
		}
	} else {
		tmp := this.Cache[key]
		tmp.Value = value
		this.moveToTail(tmp)
	}
}

func (this *LRUCache) moveToTail(n *Node) {
	if n != this.Tail {
		if n != this.Head {
			n.Pre.Next = n.Next
			n.Next.Pre = n.Pre
		} else {
			this.Head = n.Next
			this.Head.Pre = nil
		}
		this.addToTail(n)
	}
}

func (this *LRUCache) addToTail(n *Node) {
	this.Tail.Next = n
	n.Next = nil
	n.Pre = this.Tail
	this.Tail = n
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */