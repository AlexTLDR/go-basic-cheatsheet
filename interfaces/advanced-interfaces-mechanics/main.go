package main

import "fmt"

type Putter interface {
	Put(id int, val any) error
	fmt.Stringer
}

type Storage interface {
	Putter
	Get(id int) (any, error)
}

type simplePutter struct{}

func (s *simplePutter) Put(id int, val any) error {
	return nil
}

func (*simplePutter) String() string { return "" }

type FooStorage struct{}

func (s *FooStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s *FooStorage) Put(id int, val any) error {
	return nil
}

type BarStorage struct{}

func (s *BarStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s *BarStorage) Put(id int, val any) error {
	return nil
}

type Server struct {
	store Storage
}

func updateValue(id int, val any, p Putter) error {
	return p.Put(id, val)
}

func main() {

	// s := &Server{
	// 	store: &FooStorage{},
	// }
	sputter := &simplePutter{}
	updateValue(1, "foo", sputter)
}
