package main

import (
	"fmt"
	"mygosnippets/custprotobuf"
	"mygosnippets/rest"
	"mygosnippets/update"
	"sync"
)

var (
	wg *sync.WaitGroup
)

func main() {

	wg = new(sync.WaitGroup)

	fmt.Printf("Enter the choice of operation: \n")
	fmt.Printf("1. Run Update Func\n2. Make Rest Call Async\n3. Make Rest Call Sync\n4. Test Protobuf\n")
	var input string
	fmt.Scanln(&input)
	if input == "1" {
		update.Update()
	} else if input == "2" {
		for i := 0; i < 10; i++ {
			go rest.MakeRestCallAsync()
		}
	} else if input == "3" {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go rest.MakeRestCallSync(wg)
		}
		wg.Wait()
	} else if input == "4" {
		custprotobuf.ViewProtoBuf()
	}
}
