package animalfactory

type Hello struct{}

func init() {
	var hello Hello
	Register("hello", hello)

}

func (f Hello) SayHello() (result string) {
	return "Hello"
}
