package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet() *Set {
	return &Set{
		data:     make([]int, 10),
		size:     0,
		capacity: 10,
	}
}

func (s *Set) Reserve(capacity int) {
	novoData := make([]int, capacity)
	for i := 0; i < s.size; i++ {
		novoData[i] = s.data[i]
	}
	s.data = novoData
	s.capacity = capacity
}

func (s *Set) String() string {
	res := "["
	for i := 0; i < s.size; i++ {
		if i > 0 {
			res += ", "
		}
		res += fmt.Sprintf("%d", s.data[i])
	}
	res += "]"
	return res
}

func (s *Set) BinarySearch(value int) int { //começa do meio, vamo indo pra direita ou esquerda até achar o valor
	inicio := 0
	fim := s.size - 1
	for inicio <= fim { //ok, essa aqui demorou pra descobrir e sacar
		meio := (inicio + fim) / 2
		if s.data[meio] == value {
			return meio //se o cara der exatamente o valor da meiokkk
		} else if s.data[meio] < value { //se estivermos num valor mais baixo vamos pra direita
			inicio = meio + 1 //assim já filtra bem
		} else if s.data[meio] > value {
			fim = meio - 1 //cortemo
		}
	}

	return -1 //não achou
}

func (s *Set) Insert(value int) {
	if s.BinarySearch(value) != -1 {
		//fmt.Printf("valor ja inserido bro\n")
		return
	}
	if s.size == s.capacity {
		s.Reserve(2 * s.capacity)
	}

	pos := 0
	for pos < s.size && s.data[pos] < value { //só para quando encontrar um elemento maior que esse valor
		pos++ //assim sabemos onde inserir
	}

	for i := s.size; i > pos; i-- {
		s.data[i] = s.data[i-1]
	}
	s.data[pos] = value
	s.size++

}

func (s *Set) Clear() {
	s.size = 0
}
func (s *Set) Contains(value int) bool {
	if s.BinarySearch(value) == -1 {
		fmt.Println("false")
		return false
	}
	fmt.Println("true")
	return true
}

func (s *Set) Erase(value int) bool {
	pos := s.BinarySearch(value)
	if pos != -1 {
		for i := pos; i < s.size; i++ {
			s.data[i] = s.data[i+1]
		}
		s.size--
		return true
	} else {
		fmt.Println("value not found")
		return false
	}
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet()
	for scanner.Scan() {
		fmt.Print("$")
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
			v = NewSet()
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v)
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			v.Erase(value)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			v.Contains(value)
		case "clear":
			v.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
