package main

import "fmt"

type vehicle struct {
	doors uint8
	color string
}

type truck struct {
	specs          vehicle
	fourWheelDrive bool
}

type sedan struct {
	specs  vehicle
	luxury bool
}

func main() {

	t := truck{
		specs: vehicle{
			doors: 4,
			color: "black",
		},
		fourWheelDrive: true,
	}

	s := sedan{
		specs: vehicle{
			doors: 2,
			color: "red",
		},
		luxury: true,
	}

	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.specs.doors)
	fmt.Println(s.specs.color)
	fmt.Println(t.fourWheelDrive)
	fmt.Println(s.luxury)

}
