package animalfactory

// defaultAnimal implements the default matcher.
type defaultAnimal struct{}

// init registers the default matcher with the program.
func init() {
	var animal defaultAnimal
	Register("default", animal)
}

// SayHello implements the behavior for the default animal.
func (f defaultAnimal) SayHello() (result string) {
	return ""
}
