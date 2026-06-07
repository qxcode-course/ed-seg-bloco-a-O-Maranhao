package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct { //padrao das outras questoes
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet { //padrao das outras questoes
	return &MultiSet{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

// STRING
func (m *MultiSet) String() string { //vamo fazendo as coisas que parecem serem iguais
	res := "["
	for i := 0; i < m.size; i++ {
		if i > 0 {
			res += ", "
		}
		res += fmt.Sprintf("%d", m.data[i])
	}
	res += "]"
	return res
}

// Reserve
func (m *MultiSet) Reserve(newCap int) {
	novoData := make([]int, newCap)
	for i := 0; i < m.size; i++ {
		novoData[i] = m.data[i]
	}
	m.data = novoData
	m.capacity = newCap
}

// INSERT
func (m *MultiSet) Insert(value int) {
	if m.size == m.capacity {
		m.Reserve(2 * m.capacity)
	}
	//vamo meter tipo o binary search aqui tbm
	pos := 0
	for pos < m.size && m.data[pos] <= value {
		pos++
	}
	for i := m.size; i > pos; i-- { //botei m.size-1 mas ai perde o último elemento
		m.data[i] = m.data[i-1]
	}
	m.data[pos] = value
	m.size++
}

// ERASE perto do insert pq é o contrário
func (m *MultiSet) Erase(value int) {
	found, _ := m.Search(value)
	if found {
		pos := 0
		for pos < m.size && m.data[pos] < value {
			pos++
		}
		for i := pos; i < m.size-1; i++ { //só fui ver esse erro no final q ódio
			m.data[i] = m.data[i+1]
		}
		m.size--
		return
	}
	fmt.Println("value not found")
}

// Search
func (m *MultiSet) Search(value int) (bool, int) {
	low := 0
	high := m.size
	for low < high {
		middle := (high + low) / 2
		if m.data[middle] <= value {
			low = middle + 1 //se for menor ou igual, o low vira o meio mais um, encurta a pesquisa
		} else if m.data[middle] > value {
			high = middle //que nem na ultima tarefa
		}
	}
	pos := low - 1 //apontando pra última ocorrência, pq o low aponta normalmente pra depois dela né (middle +1)
	if pos >= 0 && m.data[pos] == value {
		return true, pos
	}
	return false, low //não achou, retorna a última posição de isnerção
}

// Contains
func (m *MultiSet) Contains(value int) bool {
	found, _ := m.Search(value) //demorei pra lembrar sobre usar o _ pra ignorar uma respostakk
	if !found {
		fmt.Println("false")
		return found
	} else {
		fmt.Println("true")
		return found
	}
}

// Count
func (m *MultiSet) Count(value int) int {
	found, pos := m.Search(value)
	if !found {
		return 0 //se n achou, n tem
	}
	//contador
	contador := 0
	for i := pos; i >= 0 && m.data[i] == value; i-- { //começa da ultima ocorrencia e vai enquanto o valor for igual
		contador++
	}

	return contador
}

// Unique
func (m *MultiSet) Unique() int {
	if m.size == 0 {
		return 0
	}
	contador := 1
	for i := 1; i < m.size; i++ { //começando do 1, comecei do 0 e comparava se era igual o próximo, mas aí tava fora dos limiteskk
		if m.data[i] != m.data[i-1] {
			contador++
		}
	}
	return contador
}
func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func (m *MultiSet) Clear() {
	m.size = 0
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(1) // botei 0 mas não rodoukk, talvez pq o reserve multiplica por 2 e 2x0 é 0, então botei 1 de valor mínimo, no caso quando é 0 ele só fica rodando ad eterno

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms)
		case "erase":
			value, _ := strconv.Atoi(args[1])
			ms.Erase(value)
		case "contains":
			value, _ := strconv.Atoi(args[1])
			ms.Contains(value)
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
