package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// esse problema me lembra a lenda da serpente de são luís lá do meu estado, da minha cidade
// em resumo, uma cobra tem suas partes dividas em cabeça, corpo e rabo separadas no subterrâneo da ilha
// e nunca para de crescer, quando a cabeça encostar no rabo, a ilha afunda
func (q *Queue[T]) Enqueue(value T) {
	novo := &Node[T]{Value: value}
	if q.tail == nil { //se a fila tá vazia
		q.head = novo
		q.tail = novo //esse novo vira o início, cabeça e cauda
	} else {
		q.tail.next = novo
		q.tail = novo
	}
	q.size++ //sempre incrementando
}

func (q *Queue[T]) Dequeue() (T, bool) { //agora remover do head
	if q.head == nil { //fila vazia de novo
		var zero T
		return zero, false
	}
	value := q.head.Value
	q.head = q.head.next
	if q.head == nil { //se o próximo valor for nulo, a nova cabeça no caso
		q.tail = nil // a cauda tbm é nula
	}

	q.size--
	return value, true
}

//fiz essas logo de uma vez pois são praticamente iguais

func (q *Queue[T]) Peek() (T, bool) {
	if q.head == nil {
		var zero T
		return zero, false
	}
	return q.head.Value, true
}

// func (q *Queue[T]) Size() int
// func (q *Queue[T]) IsEmpty() bool
// func (q *Queue[T]) Clear()

type Node[T any] struct {
	Value T
	next  *Node[T]
}

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) String() string {
	result := "["
	for n := q.head; n != nil; n = n.next {
		if n != q.head {
			result += ", "
		}
		result += fmt.Sprintf("%v", n.Value)
	}
	return result + "]"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	queue := NewQueue[int]()

	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println("$" + line)
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		switch args[0] {
		case "end":
			break
		case "show":
			fmt.Println(queue)
		case "push":
			for _, arg := range args[1:] {
				value, _ := strconv.Atoi(arg)
				queue.Enqueue(value)
			}
		case "pop":
			if _, ok := queue.Dequeue(); !ok {
				fmt.Println("falha: fila vazia")
			}
		case "peek":
			if value, ok := queue.Peek(); ok {
				fmt.Println(value)
			} else {
				fmt.Println("falha: fila vazia")
			}
		default:
			fmt.Println("Unknown command:", args[0])
		}
	}
}
