package main

import (
	"bufio"
	"fmt"
	"os"
	"sort" //adicionei essa pra facilitar na hora de ordenar o mapa
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair { //Função de aparição né, vamo la
	//Primeiro vamos criar um contador
	contagem := map[int]int{} //o dentro do colchetes é a chave do map.
	for _, v := range vet {
		stress := v
		if stress < 0 {
			stress = -stress //isso aqui é pra gente usar só o valor absoluto
		}
		contagem[stress]++ //bacana
	}
	//vamo ordenar as chaves desse mapa agr
	chaves := []int{} //vamo ter que armazenar elas
	for chave := range contagem {
		chaves = append(chaves, chave)
	}
	sort.Ints(chaves)

	//agora vamos montar os Pairs na ordem certinha
	resultado := []Pair{}
	for _, chave := range chaves {
		resultado = append(resultado, Pair{One: chave, Two: contagem[chave]}) //bagulho bizarro, mas pairs são bacanas
	}

	//Retornando o resultado
	return resultado
}

func teams(vet []int) []Pair { //Segunda função a ser feita por causa dos testesss. Pelo que vi não precisa de ordenação, Vi que não é obrigatória então faço depoiskkk.
	_ = vet
	return nil
}

func mnext(vet []int) []int {
	_ = vet
	return nil
}

func alone(vet []int) []int {
	_ = vet
	return nil
}

func couple(vet []int) int {
	_ = vet
	return 0
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	_ = vet
	_ = seq
	_ = pos
	return false
}

func subseq(vet []int, seq []int) int {
	_ = vet
	_ = seq
	return -1
}

func erase(vet []int, posList []int) []int {
	_ = vet
	_ = posList
	return nil
}

func clear(vet []int, value int) []int {
	resultado := []int{}
	for _, valor := range vet {
		if valor != value {
			resultado = append(resultado, value)
		}
	}
	return resultado
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
