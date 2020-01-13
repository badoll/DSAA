package main

import "fmt"

//type Type interface {}
//
//type Hashable interface {
//	Hash(tablelen int) int
//}
//
//type String string
//
//func (s String) Hash(tbl int) int {
//	hashcode := 0
//	for _,v := range s {
//		hashcode += int(v)
//	}
//	return hashcode % tbl
//}
//
//type Int int
//
//func (i Int) Hash(tbl int) int {
//	return int(i) % tbl
//}

const (
	loadfactor = 0.75
)

type Entry struct {
	key  string
	val  int
	next *Entry
}

type Hashtable struct {
	buckets []Entry
	len     int
	cap     int
}

func Hash(key string, cap int) int {
	hashcode := 0
	for _, v := range key {
		hashcode += int(v)
	}
	hashcode = hashcode % cap
	return hashcode
}

func Createhtable(cap int) *Hashtable {
	c := 2
	for c < cap {
		c <<= 1
	}
	//容量为大于所需容量的2的整数次幂
	ht := &Hashtable{make([]Entry, c), 0, c}
	return ht
}

func (ht *Hashtable) Resize(size int) {
	newht := &Hashtable{make([]Entry, size), 0, size}
	ht.cap = size
	for i := range ht.buckets {
		p := ht.buckets[i].next
		for p != nil {
			newht.Insert(p.key, p.val)
			p = p.next
		}
	}
	ht.buckets = newht.buckets
}

func (ht *Hashtable) Insert(key string, val int) {
	if float32(ht.len) >= loadfactor*float32(ht.cap) {
		ht.Resize(ht.cap * 2)
	}
	index := Hash(key, ht.cap)
	for p := ht.buckets[index].next; p != nil; {
		if p.key == key {
			p.val = val
			return
		}
		p = p.next
	}
	//新插入的数据放在链表前端，因为新插入的数据使用的概率更大
	p := &Entry{key, val, nil}
	p.next = ht.buckets[index].next
	ht.buckets[index].next = p
	ht.len++
}

func (ht Hashtable) Search(key string) (int, bool) {
	index := Hash(key, ht.cap)
	p := ht.buckets[index].next
	for p != nil {
		if p.key == key {
			return p.val, true
		}
		p = p.next
	}
	return 0,false
}

func (ht *Hashtable) Remove(key string) bool {
	index := Hash(key, ht.cap)
	p := &ht.buckets[index]
	for q := p.next; q != nil; {
		if key == q.key {
			p.next = q.next
			ht.len--
			return true
		}
		p = q
		q = q.next
	}
	return false
}

func main() {
	ht := Createhtable(1)
	ht.Insert("GO", 1)
	ht.Insert("c", 2)
	ht.Insert("c++", 3)
	ht.Insert("lua",4)
	if v, ok := ht.Search("GO"); ok {
		fmt.Println(v)
	} else {
		fmt.Println("search error")
	}
	ht.Remove("GO")
	if v, ok := ht.Search("GO"); ok {
		fmt.Println(v)
	} else {
		fmt.Println("search error")
	}
}
