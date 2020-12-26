package main

import "fmt"

type animal interface {
	growl()
}
type dog struct {

}

type cat struct {

}
func (dog dog) growl()  {
	fmt.Println("wang")
}
func (cat cat) growl()  {
	fmt.Println("miao")
}
func main() {
	var animal animal

	animal=new(dog)
	animal.growl()

	animal=new(cat)
	animal.growl()
}