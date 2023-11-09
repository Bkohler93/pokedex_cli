package cache

import (
	"testing"
	"time"
)

func TestSuccessfulAddGet(t *testing.T) {
	duration := time.Duration(time.Second * 10)
	cache := NewCache(duration)

	entries := []struct {
		key string
		val []byte
	}{
		{
			key: "www.example.com",
			val: []byte("some data"),
		},
		{
			key: "www.example2.com",
			val: []byte("some more data"),
		},
	}

	for _, e := range entries {
		cache.Add(e.key, e.val)

		v, ok := cache.Get(&e.key)
		if !ok {
			t.Error("Expected to find a key")
		}

		if string(v) != string(e.val) {
			t.Errorf("Expected %s to be %s", string(v), string(e.val))
		}
	}
}

func TestUnsuccessfulAddGet(t *testing.T) {
	duration := time.Duration(time.Second * 10)
	cache := NewCache(duration)

	key := "test"

	_, ok := cache.Get(&key)
	if ok {
		t.Error("Expected false value after retrieving invalid value from cache")
	}
}

func TestReapLoop(t *testing.T) {
	duration := time.Duration(time.Second)
	cache := NewCache(duration)

	entries := []struct {
		key string
		val []byte
	}{
		{
			key: "www.example.com",
			val: []byte("some data"),
		},
		{
			key: "www.example2.com",
			val: []byte("some more data"),
		},
	}

	for _, e := range entries {
		cache.Add(e.key, e.val)
	}

	time.Sleep(time.Second * 3)

	numEntries := len(cache.entries)
	if numEntries != 0 {
		t.Error("Expected entries to be reaped after 3 seconds")
	}
}

func TestReapLoopDoesntReap(t *testing.T) {
	duration := time.Duration(time.Second * 5)
	cache := NewCache(duration)

	entries := []struct {
		key string
		val []byte
	}{
		{
			key: "www.example.com",
			val: []byte("some data"),
		},
		{
			key: "www.example2.com",
			val: []byte("some more data"),
		},
	}

	for _, e := range entries {
		cache.Add(e.key, e.val)
	}

	time.Sleep(time.Second * 3)

	numEntries := len(cache.entries)
	if numEntries != len(entries) {
		t.Error("Expected entries to be present before reap time")
	}
}

func TestReapLoopUpdateTimeAfterRetrieving(t *testing.T) {
	duration := time.Duration(time.Second * 5)
	cache := NewCache(duration)

	entries := []struct {
		key string
		val []byte
	}{
		{
			key: "www.example.com",
			val: []byte("some data"),
		},
		{
			key: "www.example2.com",
			val: []byte("some more data"),
		},
	}

	for _, e := range entries {
		cache.Add(e.key, e.val)
	}

	entryOne := entries[0].key
	entryTwo := entries[1].key

	time.Sleep(time.Second * 3)

	_, ok := cache.Get(&entryOne)
	if !ok {
		t.Error("Expected entry to remain")
	}

	time.Sleep(time.Second * 3)

	_, ok = cache.Get(&entryOne)
	if !ok {
		t.Error("Expected entry to remain after being updated")
	}

	_, ok = cache.Get(&entryTwo)
	if ok {
		t.Error("Expected entry to be deleted since not being updated")
	}
}
