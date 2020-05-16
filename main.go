package lru

import (
	"container/list"
	"errors"
)

type Cache struct {
  // capacity of cache
	capacity int
  // doubly linked list
	items    *list.List
  // hashmap to store references of node of DLL
	hashmap  map[string]*list.Element
}

// item that is stored with node of DLL
type item struct {
	k string
	v string
}

// create cache
func New(cp int) (*Cache, error) {
	if cp <= 0 {
		return nil, errors.New("invalid cache capacity. it must be >= 1")
	}
	c := &Cache{
		capacity: cp,
		items:    list.New(),
		hashmap:  make(map[string]*list.Element),
	}
	return c, nil
}

// add to cache
func (c *Cache) Add(val string) bool {
  // TODO: we can make key and value separate
	if itm, ok := c.hashmap[val]; ok {
    // update recently used
		c.items.MoveToFront(itm)
		itm.Value.(*item).v = val
		return false
	}

	itm := &item{val, val}
  // make it recently used
	pushed := c.items.PushFront(itm)
	c.hashmap[val] = pushed

	shouldEvict := c.items.Len() > c.capacity
	if shouldEvict {
    // capacity exceeded, balance it out
		c.evict()
	}
	return shouldEvict
}

// evict from cache
func (c *Cache) evict() (val string, ok bool) {
	itm := c.items.Back()
	if itm != nil {
    // remove from DLL
    c.items.Remove(itm)
    found := itm.Value.(*item)
    // remove from hashmap
    delete(c.hashmap, found.k)
		return found.k, true
	}
	return
}

// query to cache
func (c *Cache) Query(qr string) (val string, ok bool) {
	if itm, ok := c.hashmap[qr]; ok {
    // update it as recently used
		c.items.MoveToFront(itm)
		if itm.Value.(*item) == nil {
			return "", false
		}
		return itm.Value.(*item).v, true
	}
	return
}
