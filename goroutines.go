package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall_3(i)
	}
	wg.Wait()
	fmt.Printf("the time for execution: %v\n", time.Since(t0))
	fmt.Printf("the results array is%v", results)
}

// bad way to do as it is memory unsafe
func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	fmt.Println(delay, time.Duration(delay), time.Millisecond, time.Duration(delay)*time.Millisecond)
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("the result from the database is:", dbData[i])
	results = append(results, dbData[i]) //bad way to write code
	wg.Done()
}

func dbCall_1(i int) {
	var delay float32 = rand.Float32() * 2000
	fmt.Println(delay, time.Duration(delay), time.Millisecond, time.Duration(delay)*time.Millisecond)
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("the result from the database is:", dbData[i])
	m.Lock()
	results = append(results, dbData[i]) //bad way to write code
	m.Unlock()
	wg.Done()
}

func dbCall_3(i int) {
	// experimenting with read and write blocks
	var delay float32 = rand.Float32() * 2000
	fmt.Println(delay, time.Duration(delay), time.Millisecond, time.Duration(delay)*time.Millisecond)
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(i)
	log()
	wg.Done()
}

func save(i int) {
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\n the current result is %v", results)
	m.RUnlock()
}
