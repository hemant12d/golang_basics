package main;

import (
	"errors"
	"fmt"
)

func main(){
	fmt.Println("Hello World");
	helloWorld();
	fmt.Println(sum(3, 4));
	fmt.Println(sumOfArrElem(1, 2, 3, 4, 5));
	fmt.Println(printAndCheckName("Hemant"));

	// Function result as well as error ( error return only if error available )
	result, error := printAndCheckName("Hemant");
	if error != nil {
		fmt.Println("Error: ", error);
	} else {
		fmt.Println(result);
	}

}


// Basic function
func helloWorld(){
    fmt.Println("Hello World")
}

// Function with parameter
func sum(a int, b int) int{
    return a + b;
}

// Function with list of integers
func sumOfArrElem(val ...int) int {
	var sum int;
	for _, v := range val {
		sum +=v;
	}
    return sum
}

// Function with multiple return result
func printAndCheckName(name string) (string, error){
	if len(name) == 0 {
		return "", errors.New("Name is not okay  :( ");
	}
    return "Name is : " + name, nil;
}