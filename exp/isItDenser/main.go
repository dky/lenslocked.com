package main

import "fmt"

type Rock struct {
	Mass int
	Volume int
}

type Geode struct {
	Rock
}

type Dense interface {
	Density() int
}

func (r Rock) Density() int {
	//fmt.Println(r.Mass)
	return r.Mass / r.Volume
}

func (g Geode) Density() int {
	return 100
}

//This calls the Density func above.
//func IsItDenser(a, b Rock) bool {
func IsItDenser(a, b Dense) bool {
	return a.Density() > b.Density()
}

func main() {
	/*
	aRock := Rock{Mass: 10, Volume: 50}
	bRock := Rock{Mass: 12, Volume: 60}
	//This dumps memory location
	//a := aRock.Density
	//fmt.Println(aRock.Density)
	//fmt.Println(&a) //memory address
	a := IsItDenser(aRock, bRock)
	fmt.Println(a)
	*/


	r := Rock{10, 1}
	g := Geode{Rock{2, 1}}
	fmt.Println(IsItDenser(g, r))

	rockDensity := r.Density
	fmt.Println(&rockDensity)
	
}
