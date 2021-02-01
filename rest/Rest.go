package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// MakeRestCallSync : Makes Sync Rest Call
func MakeRestCallSync(wg *sync.WaitGroup) {
	defer wg.Done()
	makeCall()
}

// MakeRestCallAsync : Makes Async Rest Call
func MakeRestCallAsync() {
	makeCall()
}

func makeCall() {
	fmt.Println("Making Rest Call")
	data := map[string]string{"firstName": "Test", "lastName": "User"}
	value, _ := json.Marshal(data)
	response, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(value))
	if err != nil {
		fmt.Println("Error while making request")
	} else {
		resp, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(resp))
	}
	fmt.Println("Completing Application")
}
