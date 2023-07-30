package main

import "fmt"

// Object or product creation worker or function
type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}

// Product structure
type House struct {
	windowType string
	doorType   string
	numFloor   int
}

// Types of builders
type NormalBuilder struct {
	doorType   string
	windowType string
	numFloor   int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (nb *NormalBuilder) setWindowType() {
	nb.windowType = "normal window"
}
func (nb *NormalBuilder) setDoorType() {
	nb.doorType = "door type"
}

func (nb *NormalBuilder) setNumFloor() {
	nb.numFloor = 1
}

func (nb *NormalBuilder) getHouse() House {

	// Return the house, once it is build
	return House{
		doorType:   nb.doorType,
		windowType: nb.windowType,
		numFloor:   nb.numFloor,
	}
}

// Igloo builder struct
type IglooBuilder struct {
	windowType string
	doorType   string
	numFloor   int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *IglooBuilder) setNumFloor() {
	b.numFloor = 1
}

// Get the endproduct, when it's build
// You can also get the house, after one of the feature is complete
func (b *IglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		numFloor:   b.numFloor,
	}
}

// Here director comes into play
type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilders(b IBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func main() {
	//iglooBuilder := getBuilder("igloo")
	//directorForiglooBuilder := newDirector(iglooBuilder)

	normalBuilder := getBuilder("normal")
	directorForNormalBuilder := newDirector(normalBuilder)
	normalHouse := directorForNormalBuilder.buildHouse()
	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.numFloor)
}
