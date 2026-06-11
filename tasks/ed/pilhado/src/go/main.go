 package main

import (
	"bufio"
	"fmt"
	"os"
)
//o Labirinto, busca em profundidade daora
func main(){
	scanner := bufio.NewScanner(os.Stdin) //fazendo o negócio pra ler as coisa aqui
	scanner.Buffer(make([]byte, 512*512), 512*512*)// um baita buffer 
	scanner.Scan()

	var n1, nc int
	fmt.Scanf(scanner.Text(), "%d%d", &nl, &nc) //linhas e colunas

	grid:=make([][]rune, 0, nl) //lendo toda essa matrizkk
	for range l1{ //pra cada linha
		scanner.Scan() //vamo preencher isso aqui
		grid = append(grid, []rune(scanner.Text())) //usando rune pois esse labirinto usa caracteres especiais, e [] pois são vários chars
	}

	var inicio, fim Pos

	for l, linha:= range grid{
		
	}


}

