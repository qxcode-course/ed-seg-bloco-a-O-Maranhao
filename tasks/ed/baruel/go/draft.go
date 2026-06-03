package main

import "fmt"

func main() {
	qtd_album := 0
	qtd_fig := 0
	fmt.Scan(&qtd_album)
	fmt.Scan(&qtd_fig)

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

	//Primeira linha é pras repetidas.
	if len(repetidos) == 0 {
		fmt.Println("N")
	} else {
		for i, valor := range repetidos {
			if i != 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%v", valor)
		}
		fmt.Println() //só pra quebra
	}

	faltou := false
	for i := 1; i <= qtd_album; i++ {
		if !unicos[i] {
			if faltou {
				fmt.Print(" ")
			}
			fmt.Printf("%v", i)
			faltou = true
		}
	}

	if !faltou {
		fmt.Println("N")
	} else {
		fmt.Println()
	}

}
