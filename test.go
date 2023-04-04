package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	type article struct {
		Name string `json:"name"`
		Description string `json:"description"`
	}
	data :=`{"name": "Tesla S", "description": "Lorem ipsum dolor sit amet"}`;
	newCar := &article{};
	json.Unmarshal([]byte(data), &newCar);
	fmt.Println(newCar.Name, newCar.Description);
}
