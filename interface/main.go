package main

type Storage interface {
	Get(id int) (any,error)
	Put(id int, val any) error
}

type FooStorage struct {}

func (s *FooStorage) Get(id int) (any,error) {
	return nil, nil
}
func (s *FooStorage) Put(id int, val any) error {
	return nil
}

type BarStorage struct {}

func (s *BarStorage) Get(id int) (any,error) {
	return nil, nil
}
func (s *BarStorage) Put(id int, val any) error {
	return nil
}
type Server struct {
	storage Storage
}
func update(id int, val any,s Storage) error{
	return s.Put(id,val)
}

func main() {
	s := Server{storage: &FooStorage{}}
	println(s.storage.Get(1))
	update(1, "foo", s.storage)
	// do something with s
}
