package main

import "fmt"

func main() {

	// Map declarations
	var dev map[string]string
	fmt.Println(dev) // []

    // Check map values
    if dev == nil {
        fmt.Println("dev map is nil")
    } else {
        fmt.Println("dev map is not nil")
    }


	// Get the length of the map
	fmt.Println("Length", len(dev))

	// Instantiate map object
	var dev2  = map[string]string{}
    fmt.Println("Instantiate map object", dev2)


	// Map with predefined values and keys
	var dev3 = map[string]string{
		"name": "Hemant",
		"age": "20",
		"role": "Software Engineer",
		"color": "green",
	}
	fmt.Println(dev3)


	// Delete properties from the map
	delete(dev3, "color")
	fmt.Println(dev3)


	// Add properties
	dev3["color"] = "green"
	fmt.Println(dev3)

	// For each with mapping
	for key, val := range dev3 {
		fmt.Println(key, " : ", val)
	}


	// Check if properties exists or not in the map
	name, ok := dev3["name"]
	if ok {
		fmt.Println("Name property is available", name)
	}

	salary, ok := dev3["salary"]
	if !ok {
        fmt.Println("Salary property is not available", salary)
    }

}