package main

import ( //importando várias bibliotecas além do main
	"fmt"
	"math"
)

func main() {
	var a, b, c, p float64 //lados de um triângulo +  o semi perímetro

	fmt.Scanf("%f %f %f", &a, &b, &c) //Scanf igual ao C
	p = (a + b + c) / 2

	var area float64 //float pode ser ou de 32 ou de 64

	area = float64(math.Sqrt(p * (p - a) * (p - b) * (p - c))) //usei float64 pois o math.sqrt pede

	fmt.Printf("%f\n", area)

}
