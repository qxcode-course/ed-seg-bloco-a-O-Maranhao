package main

import (
	"fmt"
	"math"
)

//usando a biblioteca math para o squareroot/raiz quadrada

// Ok, vamo usar algumas formulas matematicas
// tipo s = (a+b+c)/2 pro semiperímetro
// outra fórmula é a de área que é a = sqrt(s*(s-a)*(s-b)*(s-c))
func main() {
	//fmt.Println("Hello, World!") //esqueci de tirar o helloworld de primeira
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	//vamo calcular o semiperimetro
	s := (a + b + c) / 2 //quando uso o := é assumido como float64
	//agora a area
	area := math.Sqrt(s * (s - a) * (s - b) * (s - c))

	fmt.Printf("%.2f\n", area) //pra imprimir com duas casas decimais

	//Certo deu certin
}
