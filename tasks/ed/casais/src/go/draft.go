package main

import (
	"fmt"
)

func main() {

	var n int    //numero de animais
	fmt.Scan(&n) //
	//vamo trabalhar com vetores
	s := []int{} //s de solteiros
	var pares int = 0

	for i := 0; i < n; i++ {
		var animal int
		fmt.Scan(&animal)

		casados := false
		for j := 0; j < len(s); j++ {
			if s[j] == (animal * (-1)) {
				s[j] = 0
				casados = true
				pares++
				break
			}
		}
		if !casados {
			s = append(s, animal)
		}
	}

	fmt.Printf("%d\n", pares)
}
