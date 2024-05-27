package main
import "fmt"
type person struct{
	name string
	age int
}
func newPerson(name string) *person{
	p := &person{name: name}
	p.age = 30
	return p
}
func main() {
	fmt.Println(person{"bob", 20})


	fmt.Println(person{name:"Sko"})

	fmt.Println(newPerson("pedro"))

	s := person{name: "sean", age: 10}
	fmt.Println(s.name)

	sp := &s
	sp.age = 20
	fmt.Println(sp.age)
}