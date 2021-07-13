package animalfactory

import "log"

type Animal interface {
	SayHello() (result string)
}

var animals = make(map[string]Animal)

// Register is called to register a animal.
func Register(factType string, animal Animal) {
	if _, exists := animals[factType]; exists {
		log.Println(factType, "Animal already registered")
	}

	log.Println("Register", factType, "animal")
	animals[factType] = animal
}

func Call(factType string) (animal Animal, exists bool) {
	if animal, exists = animals[factType]; !exists {
		log.Println(factType, "Animal not registered")
	}
	return
}
