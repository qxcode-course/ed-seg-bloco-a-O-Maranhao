package main

//parte boa de estar atrasado é que vi esse conteúdo junto com SO, então to sabendo do que é uma pilha
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack[T any] struct {
	data []T
}

// Função top de criar novo Stack
func NewStack[T any](cap int) *Stack[T] { // pilhas guardam qualquer tipo de dado
	return &Stack[T]{data: make([]T, 0, cap)} //tamanho 0, capacidade cap
}

func (s *Stack[T]) String() string {
	output := ""
	for i := range cap(s.data) {
		if i != 0 {
			output += ", "
		}
		if i < len(s.data) {
			output += fmt.Sprintf("%v", s.data[i])
		} else {
			output += "_"
		}
	}
	return output
}

// função de Push
func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value) //graças a deus tá facil
} //deu bom
// Função de Size pq acho que vai ser simples tbm
func (s *Stack[T]) Size() int {
	return len(s.data) //acho que é só isso msm
}

// Fnção de Clear
func (s *Stack[T]) Clear() {
	s.data = s.data[:0] //começando do 0 e terminando do 0 mesmo, botei antes 0: mas vi q ai n fazia muito sentido
} //tem q ter no máximo o 0

// Função Top, só retornar o último
func (s *Stack[T]) Peek() (T, error) {
	if len(s.data) == 0 {
		var zero T                                // na questão anterior, de encadeada, tive que fazer isso tbm pois tipo genérico é foda, GO não retorna nil pra tipos genéricos
		return zero, fmt.Errorf("stack is empty") //se tiver vazio
	}
	return s.data[len(s.data)-1], nil // aquela conta estranha é pra retornar o último
}

func (s *Stack[T]) Pop() error { // aqui é só pra remover
	if len(s.data) == 0 {
		return fmt.Errorf("stack is empty")
	}
	s.data = s.data[:len(s.data)-1] //só quero ver os valores até o penúltimo
	return nil
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewStack[int](10)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			cap, _ := strconv.Atoi(parts[1])
			v = NewStack[int](cap)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Push(value)
			}
		case "debug":
			fmt.Println(v)
		case "top":
			top, err := v.Peek()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(top)
			}
		case "size":
			fmt.Println(v.Size())
		case "pop":
			err := v.Pop()
			if err != nil {
				fmt.Println(err)
			}
		case "clear":
			v.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
