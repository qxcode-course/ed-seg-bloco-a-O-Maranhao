package main

import "fmt"

func main() {
	solteiros := make(map[int]int)
	//solteiros[5] = 3 //tenho tres animais do tipo 5
	//j'ai pris tout mon amour d l'interieur de la terre, tu me manque beaucoup, au revoir, petit
	qtd := 0
	fmt.Scan(&qtd)
	pares := 0
	for range qtd {
		animal := 0 // isso é o mesmo que var animal int = 0
		fmt.Scan(&animal)
		qtd, existe := solteiros[-animal]
		if existe && qtd > 0 {
			solteiros[-animal] = qtd - 1
			pares += 1
		} else {
			solteiros[animal] = solteiros[animal] + 1 // ou solteiros[animal]++
		}
	}
	fmt.Println(pares)
}
