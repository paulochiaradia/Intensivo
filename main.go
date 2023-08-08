package main

import "github.com/paulochiaradia/Intensivo/internal/entity"

type Car struct {
	Model string
	Color string
}

// StartRun metodo
func (c *Car) StartRun() {
	println(c.Model + " has been started")
}

func (c *Car) ChangeColor(color string) {
	c.Color = color
	println("New color: " + c.Color)
}

func main() {
	//car := Car{
	//	Model: "Ferrari",
	//	Color: "Red",
	//}
	//
	//car2 := Car{
	//	Model: "Audi",
	//	Color: "Black",
	//}
	//fmt.Println(car.Model)
	//car.StartRun()
	//car2.StartRun()
	//fmt.Println(car.Color)
	//car.ChangeColor("Blue")
	//fmt.Println(car.Color)
	order, err := entity.NewOrder("1", 0, 10)
	if err != nil {
		println(err.Error())
	} else {
		println(order.ID)
	}
}
