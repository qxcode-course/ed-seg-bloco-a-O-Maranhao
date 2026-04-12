package main

import "fmt"

func main() {
	var n int //Número de pessoas da fila no início
	fmt.Scan(&n)
	var nn int               //identificador de número
	var m int                // quantidade de gente que não tá na fila
	fila := []int{}          //aqui um slice
	saidas := map[int]bool{} //Vamos usar mapa né, um mapa em go é map[KeyType]ValueType

	for i := 0; i < n; i++ {
		fmt.Scan(&nn)
		fila = append(fila, nn)
		// por enquanto tudo já é falso ,não precisa dessa linha de código aquisaidas[nn] = false //por enquanto não saiu, não é intuitivo
	}

	fmt.Scan(&m)

	for i := 0; i < m; i++ {
		fmt.Scan(&nn)
		saidas[nn] = true //marcando pra tirar depois
	}

	for _, valor := range fila { //não usei o i né, ent botei um _
		if !saidas[valor] {
			//se para esse valor ele for falso, ou seja, se não tá marcado para sair
			fmt.Printf("%d ", valor)
		}
	}
	fmt.Println()
}
