package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}

func Join(slice []int, sep string) string { //aqui não vamo mexer
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

// Status Status Status
func (v *Vector) Status() string { //v*Vector é equivalente ao This de java, engraçado
	return fmt.Sprintf("size:%d capacity:%d", v.size, v.capacity) //usando sprintf pra retornar, printf imprime no terminal mas se usar ele não funciona o return direito
}

// String String String
func (v *Vector) String() string { //saudades POO com Rubens quando eu só fazia isso no ToString
	res := "["
	for i := 0; i < v.size; i++ {
		if i > 0 {
			res += ", "
		}
		res += fmt.Sprintf("%d", v.data[i])
	}
	res += "]"
	return res
}

// RESERVA RESERVA RESERVA É MELHOR QUE TITULAR
func (v *Vector) Reserve(novaCapacidade int) {
	novoData := make([]int, novaCapacidade) // vamo redimensionar o vetor
	for i := 0; i < v.size; i++ {
		novoData[i] = v.data[i] //aqui é copiar tudo
	}
	v.data = novoData
	v.capacity = novaCapacidade
}

// PUSH ME, AND AFTER TOUCH ME, SO I CAN GET MY SATISFACTION
func (v *Vector) PushBack(value int) {
	if v.size == v.capacity {
		v.Reserve(2 * v.capacity) // aumenta a capacidade pelo dobro
	}
	v.data[v.size] = value //metendo
	v.size++               //aumentando
}

// GET GET GET
func (v *Vector) Get(index int) int {
	return v.data[index] //finalmente algo fácil
}

// AT AT AT
func (v *Vector) At(index int) (int, error) {
	if index < 0 || index >= v.size {
		return 0, fmt.Errorf("index out of range")
	}
	return v.data[index], nil
}

// /SET SET SET
func (v *Vector) Set(index int, value int) error {
	if index < 0 || index >= v.size {
		return fmt.Errorf("index out of range")
	}
	v.data[index] = value
	return nil
}

// LIMPANDO
func (v *Vector) Clear() {
	v.size = 0
}

// PopBack tipo Restart
func (v *Vector) PopBack() error {
	if v.size == 0 {
		return fmt.Errorf("vector is empty")
	}
	v.size--
	return nil
}

// INSERT INSERT INSERT
func (v *Vector) Insert(index int, value int) error {
	if index < 0 || index > v.size {
		return fmt.Errorf("index out of range")
	}
	if v.size == v.capacity {
		v.Reserve(2 * v.capacity)
	}

	for i := v.size; i > index; i-- { //essa logica aqui foi dor de cabeça tá, mas cheguei nela
		v.data[i] = v.data[i-1]
	}
	v.data[index] = value
	v.size++

	return nil
}

// ERASERHEAD (FILMAÇO)
func (v *Vector) Erase(index int) error {
	if index < 0 || index > v.size {
		return fmt.Errorf("index out of range")
	}
	for i := index; i < v.size-1; i++ {
		v.data[i] = v.data[i+1] //contrario do insert of, nem foi tão confuso
	}

	v.size--

	return nil
}

// INDEX OF INDEX OF
func (v *Vector) IndexOf(value int) int {
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			return i //se achar o valor, retorna o index
		}
	}
	return -1
}

// CONTAINS
func (v *Vector) Contains(value int) bool {
	return v.IndexOf(value) != -1 //ok, esse foi chato de pensar mesmo que seja uma resposta muito muito muito simples
}

// Slice tipo Slicers do Bioshock (Jogaço, saudades de jogar qualquer coisa); Isso é um pouco confuso mas é de boa
func (v *Vector) Slice(m int, n int) *Vector { //Vamos retornar um ponteiro, ia usar index1 e index2, mas resolvi usar m e n, pq não i e j? pq vou usar for
	if n < 0 { //n é o final
		n = v.size + n //só pra não dar errado aqui
	}
	novo := &Vector{}      //usando ponteiro né pai
	novo.data = v.data[m:] //tudo começa do começo
	novo.size = n - m      //final menos o começokk
	novo.capacity = n - m
	return novo
}

//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN//MAIN

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)
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
			value, _ := strconv.Atoi(parts[1])
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Println(v)
		case "status":
			fmt.Println(v.Status())
		case "pop":
			err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			// fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}

		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[2])
			slice := v.Slice(start, end)
			fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
