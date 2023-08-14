package main

import "fmt"

type car interface {
	stop()
	start()
}

type suv struct {
	model string
}

type sedan struct {
	model string
}

func (s sedan) stop() {
	fmt.Println("%s is stoping", s.model)
}

func (s sedan) start() {
	fmt.Println("%s is starting", s.model)
}

func (s suv) stop() {
	fmt.Println("%s is stoping", s.model)
}

func (s suv) start() {
	fmt.Println("%s is starting", s.model)
}
func main() {

	sedan := sedan{model: "toyota camry"}
	suv := suv{model: "ford"}

	fmt.Println("Using Sedan:")
	useCar(sedan)

	fmt.Println("Using SUV:")
	useCar(suv)
}

func useCar(c Car) {
	c.Start()
	c.Stop()
}
