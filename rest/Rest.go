package main

import (
	"fmt"
	"sync"

	rest "github.com/neo7337/goRep/rest/restutils"
)

var (
	wg *sync.WaitGroup
)

func main() {

	wg = new(sync.WaitGroup)

	fmt.Printf("Enter the choice of operation: \n")
	fmt.Printf("1. Make Rest Call Async\n2. Make Rest Call Sync\n")
	var input string
	fmt.Scanln(&input)

	if input == "1" {
		for i := 0; i < 10; i++ {
			go rest.MakeRestCallAsync()
		}
	} else if input == "2" {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go rest.MakeRestCallSync(wg)
		}
		wg.Wait()
	}

}
