package main

import (
	"fmt"
)

//vou usar pra calcular valor absoluto

func main() {
	var n int //numero de competidores n
	fmt.Scan(&n)
	melhor := 1000000 //essa aqui é a melhor pontuação atual, que é infinito praticamente esse numero é grande bagarai
	vencedor := -1    //esse é o vencedor atual, mentira não tem nenhum, mas esse é um valor negativo que se for preservado no final, significa que não tem ganhador

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		if a >= 10 && b >= 10 {
			pontuacao := a - b
			if pontuacao < 0 {
				pontuacao = -pontuacao
			}
			if pontuacao < melhor { //a logica que vou usar aquikk
				melhor = pontuacao
				vencedor = i
			}
		}
	}
	//Agora vendo quem venceu
	if vencedor == -1 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Printf("%d\n", vencedor)
	}
}
