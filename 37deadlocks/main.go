package main

import (
	"fmt"
	"sync"
)

// how to make thread safe collection
type Collection struct {
	//to make it thread safe pass mutex as well
	Mt   sync.Mutex
	Data map[string]string
}

func NewCollection() Collection {
	return Collection{
		Mt:   sync.Mutex{},
		Data: map[string]string{},
	}
}

func (c *Collection) Has(key string) bool {
	c.Mt.Lock()
	defer c.Mt.Unlock()
	if _, ok := c.Data[key]; ok {
		return true
	}
	return false
}

func (c *Collection) Add(key, value string) {
	c.Mt.Lock()
	defer c.Mt.Unlock()
	//lock() and lock() will occur again  by calling another Has() function which will have lock()
	//if c.Has(key) {
	//it will duplicate the code but will make sure your code doesn't go deadlock
	if _, ok := c.Data[key]; ok {
		return
	}
	c.Data[key] = value
}

func main() {
	fmt.Println("testing deadlocks...")

	/*mt := sync.Mutex{}

	mt.Lock()
	mt.Lock()
	//mutex is expecting unlock after lock
	//if it finds lock again it will be in deadlock situation

	mt.Unlock()
	mt.Unlock()

	fmt.Println("finished...")
	*/

	//deadlock for collection

	coll := NewCollection()
	coll.Add("ajay", "kumar")
	coll.Add("raghav", "sharma")

	check := coll.Has("ajay")
	fmt.Printf("name is availble: %v", check)

	//fmt.Println(coll)
	fmt.Printf("collection: %v", coll)

}
