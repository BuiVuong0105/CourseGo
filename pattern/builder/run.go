package builder

import "fmt"

func TestBuilder() {

	var houseNormal *House = Builder("normal").WindowType("Wooden Window").DoorType("Wooden Door").Floor(1).Build()
	fmt.Println("House Nornal: ", *houseNormal)

	var houseigLoo *House = Builder("igloo").WindowType("Snow Window").DoorType("Snow Door").Floor(2).Build()
	fmt.Println("House Igloo: ", *houseigLoo)
}
