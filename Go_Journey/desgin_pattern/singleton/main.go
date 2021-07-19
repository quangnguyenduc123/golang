package main

import (
	"fmt"
	"sync"
)

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// in addition to sync.Once, there're init(), laziness
var once sync.Once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		caps := make(map[string]int) // get from db
		db := singletonDatabase{caps}
		instance = &db
	})
	return instance
}

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city) // => Dependency inversion principle all functions depend on GetSingletonDatabase
	}
	return result
}

func main() {
	db := GetSingletonDatabase()
	pop := db.GetPopulation("London")
	fmt.Println(pop)
}
