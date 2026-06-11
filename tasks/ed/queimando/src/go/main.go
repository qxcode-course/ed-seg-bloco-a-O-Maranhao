package main

import (
	"bufio"
	"fmt"
	"os"
)

// Isso aqui é massa demais, lembro da aula sobre e lembro do Paint/Photoshop e outros apps q usam balde de tinta
// struct de posição que não existe ams vou fazer agora
type Pos struct {
	l, c int
}

func burnTrees(grid [][]rune, l, c int) {
	stack := NewStack[Pos]()
	stack.Push(Pos{l, c}) //adicionando

	// - enquanto a pilha não estiver vazia:
	for !stack.IsEmpty() {
		// retirar o elemento do topo
		pos := stack.Pop()
		l, c := pos.l, pos.c
		// vamo verifica se tá dentro da matriz e se é uma árvore no caso
		if l < 0 || l >= len(grid) || c < 0 || c >= len(grid[l]) {
			continue //pula pois tá fora da matriz
		}
		if grid[l][c] != '#' {
			continue //não é flor que se cheire
		}
		//se passar por tudo isso, QUEIMA!
		grid[l][c] = 'o'

		//agr vamo adicionar os 4 vizinhos
		stack.Push(Pos{l - 1, c}) //cima cima
		stack.Push(Pos{l + 1, c}) //baixo baixo
		stack.Push(Pos{l, c + 1}) //esquerda
		stack.Push(Pos{l, c - 1}) //direita esquerda direita
		//B A, Konami Code aí pra conseguir mais vidas no Contra do Nintendinho
	}
	// Essa função deve usar uma list como pilha
	// e marcar as árvores na matriz como queimados
	// Uma sugestão de como fazer isso é:
	// - adicionar a primeira posição na pilha
	// - enquanto a pilha não estiver vazia:
	//   - retirar o elemento do topo
	//   - se puder ser queimado, queime e adicione seus vizinhos à pilha

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
