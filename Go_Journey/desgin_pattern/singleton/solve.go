package main

import (
	"fmt"
	"sync"
)

type Database interface {
	GetPopulation(name string) int
}

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
		result += GetSingletonDatabase().GetPopulation(city) // => Dependency inversion principle all functions depend on GetSingletonDatabase => hard to write unit test
	}
	return result
}

func GetTotalPopulationEx(d Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += d.GetPopulation(city)
	}
	return result
}

type DummyDatabase struct {
	dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
		}
	}
	return d.dummyData[name]
}

func main() {
	cities := []string{"alpha", "beta"}
	tp := GetTotalPopulationEx(&DummyDatabase{}, cities)
	fmt.Println(tp)
}
