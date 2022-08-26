package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	oman := &Person{
		Name: "oman",
		Age:  24,
		SocialFollowers: &SocialFollowers{
			Instagram: 100,
			Youtube:   200,
		},
	}
	data, err := proto.Marshal(oman)
	if err != nil {
		log.Fatalf("marshalling error %v\n", err)
	}

	fmt.Println(data)

	newOman := &Person{}
	err = proto.Unmarshal(data, newOman)
	if err != nil {
		log.Fatalf("unmarshalling error %v\n", err)
	}
	fmt.Println(newOman.GetAge())
	fmt.Println(newOman.GetName())
	fmt.Println(newOman.SocialFollowers.GetInstagram())
	fmt.Println(newOman.SocialFollowers.GetYoutube())
}
