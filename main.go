package main

import (
	"fmt"

	"github.com/hellicopthecat/learngo/person"
)

func main() {
	choi := person.Person{}
	choi.SetDetails("choi", 30)
	fmt.Println(choi)
}
