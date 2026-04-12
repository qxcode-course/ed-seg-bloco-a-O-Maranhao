package main

import "fmt"

func mostrar_jogadores(jogadores []bool, espada int) {
	fmt.Print("[")
	for i, valor := range jogadores {
		if valor == false {
			continue
		}
		fmt.Print(i + 1)
		if espada == i+1 {
			fmt.Print(">")
		}
		fmt.Print(" ")
	}
	fmt.Println("]")
}

func procurar_vivo(jogadores []bool, espada int) int {
	for {
		espada = (espada + 1) % len(jogadores)
		if espada == len(jogadores) {
			espada = 0
		}
		if jogadores[espada] == true {
			return espada
		}
	}
}

func main() {
	var qtd, espada int
	fmt.Scan(&qtd, &espada)
	jogadores := make([]bool, qtd)
	for i := range jogadores {
		jogadores[i] = true
	}

	for range qtd - 1 {
		mostrar_jogadores(jogadores, espada)
		vai_morrer := procurar_vivo(jogadores, espada)
		jogadores[vai_morrer] = false
		espada = procurar_vivo(jogadores, espada)
	}
	mostrar_jogadores(jogadores, espada)

}
