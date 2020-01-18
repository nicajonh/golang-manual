package main

import (
	"./Lru"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lru(t *testing.T){
	c := Lru.NewLRUCache(2)
	c.Set("K1", 1)
	c.Set("K2", 2)
	c.Set("K1", 100)
	assert.Equal(t,c.Get("K1"),100)
	c.Set("K3", 3)
	assert.Equal(t,c.Get("K3"),3)
	t.Log(c.Get("K2"))
	c.Set("K4", 4)
}


