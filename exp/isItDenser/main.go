package main

import "fmt"

type Rock struct {
	Mass int
	Volume int
}

func (r Rock) Density() int {
	return r.Mass / r.Volume
}

func IsItDenser(a, b Rock) bool {
	return a.Density() > b.Density()
}

func main() {
	aDenseRock := Rock{Mass: 10, Volume: 50}
	bDenseRock := Rock{Mass: 12, Volume: 60}

	a := IsItDenser(aDenseRock, bDenseRock)
	fmt.Println(a)
	
}
