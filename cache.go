package main

import (
	"container/list"
	"fmt"
)

// DO NOT CHANGE THIS CACHE SIZE VALUE
const CACHE_SIZE int = 3

type cacheEntry struct {
	key   int
	value int
}

type Cache struct {
	limit     int
	list      *list.List
	hashTable map[int]*list.Element
	Size      int
}

var myCache *Cache = New()

func New() *Cache {
	return &Cache{
		limit:     CACHE_SIZE,
		list:      list.New(),
		hashTable: make(map[int]*list.Element),
	}
}

func Get(key int) int {

	element, success := myCache.hashTable[key]
	if !success {
		return -1
	}
	myCache.list.MoveToFront(element)
	return element.Value.(*cacheEntry).value
}

func Delete(key int) {
	element, success := myCache.hashTable[key]
	if !success {
		return
	}
	delete(myCache.hashTable, key)
	v := myCache.list.Remove(element).(*cacheEntry)
	fmt.Println(v.value)
	myCache.Size -= 1
}

func resize() {
	for myCache.Size > myCache.limit {
		element := myCache.list.Back()
		if element == nil {
			return
		}
		v := myCache.list.Remove(element).(*cacheEntry)
		delete(myCache.hashTable, v.key)
		myCache.Size -= 1
	}
}

func Set(key int, value int) {
	_, success := myCache.hashTable[key]
	if success {
		element, success := myCache.hashTable[key]

		if success {
			myCache.list.MoveToFront(element)
			return
		}

	}
	v := &cacheEntry{key, value}
	element := myCache.list.PushFront(v)
	myCache.hashTable[key] = element
	myCache.Size += 1
	resize()
}
