package hashtable

import (
	"errors"
	"hash/fnv"
)

type kv struct {
	Key, Value string
}

type Table struct {
	m     int
	table [][]kv
}

var errKey = errors.New("error, wrong key")

func (t *Table) hash(key string) int {
	h := fnv.New32()
	h.Write([]byte(key))

	return int(h.Sum32()) % t.m
}

func (t *Table) Get(key string) (string, bool) {
	hash := t.hash(key)

	for i, kv := range t.table[hash] {
		if key == kv.Key {
			return t.table[hash][i].Value, true
		}
	}

	return "", false
}

func (t *Table) Insert(key, value string) {
	i := t.hash(key)

	for j, kv := range t.table[i] {
		if key == kv.Key {
			t.table[i][j].Value = value

			return
		}
	}

	t.table[i] = append(t.table[i], kv{
		Key:   key,
		Value: value,
	})
}

func (t *Table) Delete(key string) error {
	hash := t.hash(key)

	for i, v := range t.table[hash] {
		if v.Key == key {
			current := t.table[hash]
			current[i] = current[len(current)-1]
			current = current[:len(current)-1]
			t.table[hash] = current

			return nil
		}
	}

	return errKey
}

func New(m int) *Table {
	return &Table{
		m:     m,
		table: make([][]kv, m),
	}
}
