package main

import (
	"fmt"
)

//Hoje vamos trabalhar com loops eu acho

func main() {
	q := NewQueue[string]() //q de queue, ou fila em francês/english

	for i := 0; i < 16; i++ {
		q.Enqueue(string(rune('A' + i))) // esse rune/char aqui é uma loucura: soma os valores da tabela ASCII e o String converte esses valores em string, então eu to somando o valor de A com o valor de i
	} //e assim vou preenchendo rapidin

	for q.items.Len() > 1 {
		timeA := q.Dequeue() //remove o time A, primeiro elemento da fila
		timeB := q.Dequeue() //remove o time B, q agr é o primeiro elemento da fila

		var gA, gB int //gols de A e gols de B
		fmt.Scan(&gA, &gB)
		if gA > gB { //se o time A fizer mais gol
			q.Enqueue(timeA) //time A volta pra fila
		} else {
			q.Enqueue(timeB) //empate ou vitória de B
		}
	}

	//vamo imprimir o vencedor, que foi quem sobrou
	fmt.Println(q.Dequeue())
}
