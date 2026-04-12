package main

import "fmt"

func main() {
	var n int
	var e int //dono da espada
	fmt.Scan(&n, &e)
	pessoas := []int{} //colchetes dizem o tamanho, quando vazios significa que é um slide;
	for i := 1; i <= n; i++ {
		pessoas = append(pessoas, i) //assim que se adiciona em Go, não é tipo um método como em Java
	}

	var espada int                      // ou espada := 0, essa espada é o indíce de E
	for i := 0; i < len(pessoas); i++ { //estou atrabalho aqui com os índices, não com os valores
		if pessoas[i] == e {
			espada = i //espada é um índice
			break
		}
	}

	for len(pessoas) > 1 { //Enquanto tiver gente
		//Esse vai ser o for da impressão
		fmt.Printf("[ ")
		for i, valor := range pessoas { //Um for para ver quem tem a espada, i é o índice e o valor é o valor que tá na posição sendo visa
			fmt.Print(valor) //coloca o valor e imprime um espaço.
			if i == espada {
				fmt.Printf(">")
			}
			fmt.Print(" ")
		}
		fmt.Println("]") //só pra quebrar linha
		//vamos descobrir a vítima, no caso é (quem tem a espada + 1)%len(pessoas)
		vitima := (espada + 1) % len(pessoas)                     //essa é a vítima, o índice dela no caso né. Essa é a posição da vítima
		pessoas = append(pessoas[:vitima], pessoas[vitima+1:]...) //Aqui, eu to passando um slice como se fosse um argumento, o : significa "Tudo que vem antes ou depois", e as reticiências é para passar um array dinamico como argumentos
		espada = vitima % len(pessoas)                            //o problema agora é que a espada não tá ficando com o sobrevivente
	}
	fmt.Printf("[ %d> ]\n", pessoas[0])

}
