package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BetterSearch(slice []int, value int) (bool, int) { //lembrou o binarySearch
	high := len(slice) //agora não incluimos o último né //high é final
	low := 0           // é o começo
	for low < high {
		middle := (high + low) / 2
		if slice[middle] == value {
			return true, middle //achamo diretokk
		} else if value > slice[middle] {
			low = middle + 1
		} else if value < slice[middle] {
			high = middle //tinha botado -1, mas me toquei q não é a mesma coisa
		}
	}
	_, _ = slice, value //mexer aqui não
	return false, low   //basicamente é onde o bicho deveria ficar né CARALHO NEM FODENDO QUE FOI TÃO RÁPIDO
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	slice := []int{}
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	found, result := BetterSearch(slice, value)
	if found {
		fmt.Println("V", result)
	} else {
		fmt.Println("F", result)
	}
}
