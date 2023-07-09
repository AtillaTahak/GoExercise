package main

type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func (m *CustomMap[K,V]) Set(key K, value V) {
	m.data[key] = value
}

func NewCustomMap[K comparable, V any]() *CustomMap[K,V] {
	return &CustomMap[K,V]{data: make(map[K]V)}
}

func main(){
	m := NewCustomMap[string, int]()
	m.Set("one", 1)
	m.Set("two", 2)
	m.Set("three", 3)
	print(m.data)
}

