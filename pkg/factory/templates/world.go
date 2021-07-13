package animalfactory

type World struct{}

func init() {
	var world World
	Register("world", world)

}

func (d World) SayHello() (result string) {
	return "world"
}
