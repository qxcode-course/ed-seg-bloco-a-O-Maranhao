package main

import "fmt"

func main() {
	var h, p, f, d int //variaveis daora de entrada
	fmt.Scan(&h, &p, &f, &d)
	//lembro de fazer essa questao em C
	for { //acabei de descobrir que nao existe while in Golang (na real eu ja sabia e so esqueci)
		//qual a logica: nosso fugitivo anda até cair numa casa do helicoptero ou da policia, a direção define se é horarío ou antihorário
		f = (f + d + 16) % 16
		if f == h {
			fmt.Printf("S\n")
			break
		} else if f == p {
			fmt.Printf("N\n")
			break
		}
	}

}
