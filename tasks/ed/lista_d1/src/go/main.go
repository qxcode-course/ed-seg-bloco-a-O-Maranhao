package main

//lista duplamente encadeada
//cada elemento é como um nó que aponta pro pŕoximo
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// cada nó tem
type Node struct {
	Value int
	next  *Node
	prev  *Node
}

type LList struct {
	root *Node //nó de marcação/sentinela
}

func NewLList() *LList {
	root := &Node{}
	root.next = root
	root.prev = root
	return &LList{root: root}
}

func (ll *LList) String() string {
	res := "["
	node := ll.root.next

	for node != ll.root {
		if node != ll.root.next {
			res += ", "
		}
		res += fmt.Sprintf("%d", node.Value)
		node = node.next
	}
	res += "]"
	return res
}

// PushFront
func (ll *LList) PushFront(value int) {
	novo := &Node{Value: value} //crianddo um novo nó
	novo.next = ll.root.next
	novo.prev = ll.root

	ll.root.next.prev = novo //o antigo primeiro agora aponta pro novo
	ll.root.next = novo      //root que é o 0 na minha cabeça aponta pro novo
}

// Size
func (ll *LList) Size() int {
	count := 0
	node := ll.root.next  //começa no primeiro
	for node != ll.root { //enquanto não der a volta
		count++
		node = node.next
	}
	return count //baita contador
}

// Clear
func (ll *LList) Clear() {
	ll.root.next = ll.root //só voltar ao 0
	ll.root.prev = ll.root
}

// Push Back aqui
func (ll *LList) PushBack(value int) { //adicionamo no final da lista
	novo := &Node{Value: value}
	novo.prev = ll.root.prev //apontando pro antigo último
	novo.next = ll.root      //novo aponta pra root
	ll.root.prev.next = novo //o antigo ultimo aponta pro novo
	ll.root.prev = novo      //o root aponta para o novo
}

//PopFront, só remove

func (ll *LList) PopFront() {
	if ll.root.next == ll.root {
		return
	} //lista vazia, não tem como né

	first := ll.root.next     //guarda o primeiro nó
	ll.root.next = first.next //root aponta pro segundo
	first.next.prev = ll.root //o segundo aponta pro root

}

func (ll *LList) PopBack() {
	if ll.root.next == ll.root {
		return //de novo
	}
	last := ll.root.prev
	ll.root.prev = last.prev
	last.prev.next = ll.root
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
