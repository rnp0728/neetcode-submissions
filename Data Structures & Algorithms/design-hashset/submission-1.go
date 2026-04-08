type MyHashSet struct {
	storage map[int]struct{}
}

func Constructor() MyHashSet {
	s := make(map[int]struct{})
    return MyHashSet{
		storage: s,
	}
}

func (this *MyHashSet) Add(key int) {
    this.storage[key] = struct{}{}
}

func (this *MyHashSet) Remove(key int) {
	delete(this.storage, key)
}

func (this *MyHashSet) Contains(key int) bool {
    _, ok := this.storage[key]
	if ok {
		return true
	} else {
		return false
	}
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 