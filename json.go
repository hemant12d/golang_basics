package main

import (
	"encoding/json"
	"fmt"
)

type Dev struct {
	Name string `json: "name"`
	Age int `json: "age"`
}

func main()  {

	// Json string response
	strDevData := `{"name":"hemant","age":2}`
	//fmt.Println(sFrom)

	var devDetails Dev

	err := json.Unmarshal([]byte(strDevData), &devDetails)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(devDetails)

	sto, err := json.Marshal(devDetails)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sto)
	fmt.Printf("%T\n", sto)
	fmt.Printf("%s\n", sto)
}
