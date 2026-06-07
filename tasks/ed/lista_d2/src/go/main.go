package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// VAMO COMEÇAR PELO NODE que é mais tranquilo (eu acho)
type Node struct {
	Value int //botei minúsculo no começo mas corri
	next  *Node
	prev  *Node
	root  *Node
}

// Funções de Node
func (n *Node) Next() *Node {
	if n.next == n.root { //se próximo é o sentinela, joga fora
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}
	return n.prev
}

// Certin, agora vamo pra LLIST
type LList struct {
	root *Node
	size int
}

// Functions de LList
func NewLList() *LList {
	root := &Node{}
	root.next = root
	root.prev = root
	root.root = root
	return &LList{root: root}
}

// Funções que são praticamente iguais as outras
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
func (ll *LList) PushFront(value int) {
	novo := &Node{Value: value, root: ll.root}
	novo.next = ll.root.next
	novo.prev = ll.root
	ll.root.next.prev = novo
	ll.root.next = novo
	ll.size++ //isso é novo
}
func (ll *LList) PushBack(value int) {
	novo := &Node{Value: value, root: ll.root}
	novo.prev = ll.root.prev
	novo.next = ll.root
	ll.root.prev.next = novo
	ll.root.prev = novo
	ll.size++
}

func (ll *LList) Size() int {
	return ll.size
}

func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0 //isso é novo tbm
}

// Front e Back
func (ll *LList) Front() *Node {
	if ll.root.next == ll.root {
		return nil
	}
	return ll.root.next
}
func (ll *LList) Back() *Node {
	if ll.root.prev == ll.root {
		return nil
	}
	return ll.root.prev
}

// Search
func (ll *LList) Search(value int) *Node {
	for node := ll.Front(); node != nil; node = node.Next() { //DEMOREI PRA UM CARALHO PRA NÃO CONFUNDIR O MÉTODO PUBLICO COM O ATRIBUTO PRIVADO
		if node.Value == value {
			return node
		}
	}
	return nil
}

// Insert
func (ll *LList) Insert(node *Node, value int) {
	novo := &Node{Value: value, root: ll.root}
	novo.next = node //aponta pro nó atual
	novo.prev = node.prev
	node.prev.next = novo
	node.prev = novo
	ll.size++
}

// Remove
func (ll *LList) Remove(node *Node) *Node {
	node.prev.next = node.next
	node.next.prev = node.prev
	ll.size--
	return node.Next()
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
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.Value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
