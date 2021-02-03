package custprotobuf

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

func ViewProtoBuf() {
	user := &User{
		Name: "Test User",
		Age:  22,
	}

	data, err := proto.Marshal(user)
	if err != nil {
		log.Fatalf("Error during marshalling")
	}
	fmt.Println(data)

	newUser := &User{}
	err = proto.Unmarshal(data, newUser)
	if err != nil {
		log.Fatalf("Unmarshal Error")
	}
	fmt.Println(newUser.GetAge())
	fmt.Println(newUser.GetName())
}
