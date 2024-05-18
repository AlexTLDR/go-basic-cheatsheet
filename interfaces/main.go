package main

import "fmt"

type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

type ApiServer struct {
	numberStore NumberStorer
}

type MongoDBNumberStore struct {
	// some values
}

func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (m MongoDBNumberStore) Put(number int) error {
	fmt.Println("store the number into the mongoDB storage")
	return nil
}

func main() {
	apiServer := ApiServer{
		numberStore: MongoDBNumberStore{},
	}

	numbers, err := apiServer.numberStore.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(numbers)
}
