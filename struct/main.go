package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

type animal struct {
	raptiles bool
	flying   bool
}

type lion struct {
	animal
	category string
	legs     int
}

func main() {
	// this is initilizing the instence
	animels := lion{
		category: "Carnivore",
		legs:     4,
		animal: animal{
			raptiles: false,
			flying:   false,
		},
	}
	myMessage := messageToSend{}
	fmt.Println(myMessage)
	fmt.Println(animels.raptiles)
	fmt.Println(animels.category)
	fmt.Println(animels.flying)
	fmt.Println(animels.legs)
}
