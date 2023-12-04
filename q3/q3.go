package q3

//Você recebe um tabuleiro retangular de M x N quadrados. Além disso, você tem um número ilimitado de peças de dominó
//padrão de 2 x 1 quadrados. Você pode girar as peças. Você deve colocar o máximo de peças de dominó possível no
//tabuleiro, seguindo as seguintes condições:
//
//1. Cada peça de dominó cobre completamente dois quadrados.
//2. Nenhuma duas peças de dominó se sobrepõem.
//3. Cada peça de dominó está completamente dentro do tabuleiro. É permitido que a peça toque as bordas do tabuleiro.
//
//Encontre o número máximo de peças de dominó que podem ser colocadas sob essas restrições.
//
//Se M ou N forem iguais ou menores que 0, a função deve retornar um erro.


import (
	"fmt"
)

func contarPecasDominos(M, N int) (int, error) {
	if M <= 0 || N <= 0 {
		return 0, fmt.Errorf("M e N devem ser maiores que 0")
	}

	numPecas := (M * N) / 2

	return numPecas, nil
}

func main() {
	M := 3
	N := 3

	numPecas, err := contarPecasDominos(M, N)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Número máximo de peças de dominó:", numPecas)
	}
}
package main

import (
	"fmt"
)

func contarPecasDominos(M, N int) (int, error) {
	if M <= 0 || N <= 0 {
		return 0, fmt.Errorf("M e N devem ser maiores que 0")
	}

	numPecas := (M * N) / 2

	return numPecas, nil
}

func main() {
	M := 3
	N := 3

	numPecas, err := contarPecasDominos(M, N)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Número máximo de peças de dominó:", numPecas)
	}
}
