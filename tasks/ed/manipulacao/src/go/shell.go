package main

//parece que só vou precisar fazer as função
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	resultado := []int{}
	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			resultado = append(resultado, vet[i])
		}
	}
	return resultado
}

func getCalmWomen(vet []int) []int {
	resultado := []int{} //[] declara que é um slice, int é o tipo, {} é iniciando ele vazio
	for i := 0; i < len(vet); i++ {
		if vet[i] < 0 {
			stress := vet[i] * (-1)
			if stress < 10 {
				resultado = append(resultado, vet[i])
			}
		}
	}
	return resultado
}

func sortVet(vet []int) []int {
	ordenado := make([]int, len(vet)) //make cria um slice de tamanho fixo do tipo make(tipo de valor, tamanho)
	copy(ordenado, vet)               //(copiador, copiado)
	sort.Ints(ordenado)               //biblioteca sort
	return ordenado
}

func sortStress(vet []int) []int {
	ordenado := make([]int, len(vet))
	copy(ordenado, vet)

	sort.Slice(ordenado, func(i, j int) bool {
		stressI := ordenado[i]
		stressJ := ordenado[j]

		if stressI < 0 {
			stressI = -stressI
		}
		if stressJ < 0 {
			stressJ = -stressJ
		}

		return stressI < stressJ
	})
	return ordenado
}

func reverse(vet []int) []int {
	resultado := []int{} //engraçado, quando fiz resultado:=make([]int, len(vet)) deu errado
	for i := len(vet) - 1; i >= 0; i-- {
		resultado = append(resultado, vet[i])
	}
	return resultado
}

func unique(vet []int) []int {
	//vamo ter que verificar a unicidade, criamos um slice para "marcar" os valores já vistos
	unico := map[int]bool{} // mapa[valor da chave]valor que ela armazena
	resultado := []int{}
	for i := 0; i < len(vet); i++ {
		valor := vet[i]
		if !unico[valor] {
			resultado = append(resultado, valor)
			unico[valor] = true
		}
	}
	return resultado
}

func repeated(vet []int) []int {
	unico := map[int]bool{} //verificar unicidade
	resultado := []int{}

	for i := 0; i < len(vet); i++ {
		valor := vet[i] //valor que o index pega
		if unico[valor] {
			resultado = append(resultado, vet[i])
		} else {
			unico[valor] = true
		}
	}
	sort.Ints(resultado)

	return resultado
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
