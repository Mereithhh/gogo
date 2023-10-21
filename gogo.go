package gogo

import "fmt"

func Sprint(any interface{}) string {
	return superPrintf(any)
}

func Print(any interface{}) {
	fmt.Println(superPrintf(any))
}
