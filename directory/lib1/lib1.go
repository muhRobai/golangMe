package lib1

import (
	"fmt"
)

func Print(name string)  {
	fmt.Println("ini lib1")
	introduce(name)
}

func introduce(name string) {
    fmt.Println("nama saya", name)
}