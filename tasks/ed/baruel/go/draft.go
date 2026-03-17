package main

import "fmt"

func main() {
	qtd_album := 0
	qtd_fig := 0
	fmt.Scan(&qtd_album, &qtd_fig)
	album := make([]int, qtd_fig) //criei um vetor que é o tamanho das figurinhas q cada album tem
	unicos := make(map[int]bool)  //criando um mapa
	repetidos := make([]int, 0, qtd_fig)
	for i := range album { //for do tamanho do album
		fmt.Scan(&album[i])
	}
	for _, fig := range album { //para cada figurinha no album
		if unicos[fig] {
			repetidos = append(repetidos, fig) //se meu mapa de unicos já possui essa figurinha, bota nos repetidos
		} else {
			unicos[fig] = true
		}
	}
	if len(repetidos) == 0 {
		fmt.Println("N")
	} else {
		for i, valor := range repetidos {
			if i != 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%v", valor)
		}
	}
	fmt.Println("") //criar um vetor de saida aqui ou coisa do tipo sla
	for i := 1; i <= qtd_album; i++ {
		if !unicos[i] {
			fmt.Printf("%v", i)
			fmt.Printf(" ")
		}
	}
	fmt.Print("\n")

}
