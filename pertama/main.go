package main

import (
	"fmt"
	"pertama/calculation"
)
	

func main(){
	fmt.Println("Halo, Belajar Perkalian Dengan Golang")
	result := calculation.Multiply(3,3)
	fmt.Println(result)
}