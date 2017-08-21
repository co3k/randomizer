package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a := os.Args
	l := len(a) - 1
	i := r.Intn(l) + 1

	fmt.Printf("%s\n", a[i])
}
