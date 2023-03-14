package main;

import (
	"encoding/json"
	"fmt"
);

// Declare struct
type Dev struct {
	Name string `json: "name"`
	Age int `json: "age"`
}

func main()  {

	// Json string response
	sFrom := `{"name":"hemant","age":2}`
	//fmt.Println(sFrom)

	var dev_1 Dev;

	err := json.Unmarshal([]byte(sFrom), &dev_1);
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dev_1);

	sto, err := json.Marshal(dev_1);

	if err != nil {
		fmt.Println(err);
	}
	fmt.Println(sto)
	fmt.Printf("%T\n", sto);
	fmt.Printf("%s\n", sto)
}
