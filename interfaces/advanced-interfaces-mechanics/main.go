package main

type Storage interface {
	Get(id int) (any, error)
	Put(id int, val any) error
}

type Server struct {
	store Storage
}

func main() {

	s := &Server{}
	s.store.Get(1)
	s.store.Put(1, "foo")
}
