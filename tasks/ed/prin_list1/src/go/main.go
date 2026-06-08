package main

import (
	"fmt"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	//parece que só precisamos implementar duas funções
	//aqui percorre a lista
	res := "[ "
	for n := l.Front(); n != l.End(); n = n.next {
		if n == sword {
			res += fmt.Sprintf("%d> ", n.Value)
		} else {
			res += fmt.Sprintf("%d ", n.Value)
		}

	}
	res += "]"
	return res
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	//aqui é só dizer que quando chegar no fim é pra voltar do início
	if it.Next() == l.End() {
		return l.Front()
	}
	return it.Next()
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	//fmt.Println(qtd, chosen) isso aqui tava atrapalhando
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
