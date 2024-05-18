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

type PostgresNumberStorer struct {
	// some postgres stuff
}

func (p PostgresNumberStorer) GetAll() ([]int, error) {
	return []int{10, 20, 30}, nil
}

func (p PostgresNumberStorer) Put(number int) error {
	fmt.Println("store the numbers into postgres storage")
	return nil
}

func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (m MongoDBNumberStore) Put(number int) error {
	fmt.Println("store the number into the mongoDB storage")
	return nil
}

func main() {
	apiServerMongo := ApiServer{
		numberStore: MongoDBNumberStore{},
	}

	numbers, err := apiServerMongo.numberStore.GetAll()
	if err != nil {
		panic(err)
	}
	if err = apiServerMongo.numberStore.Put(1); err != nil {
		panic(err)
	}
	fmt.Println(numbers)

	apiServerPostgres := ApiServer{
		numberStore: PostgresNumberStorer{},
	}

	numbers, err = apiServerPostgres.numberStore.GetAll()
	if err != nil {
		panic(err)
	}

	if err = apiServerPostgres.numberStore.Put(1); err != nil {
		panic(err)
	}
	fmt.Println(numbers)
}
