package lib

import (
	"math/rand"
	"time"
)

/*
 * Verifiquei que o método rand.Intn(x) produz um número pseudo-aleatório
 * inteiro de valor entre 0 (inclusive) e x (exclusive) o que é exatamente
 * aquilo de que eu preciso!
 *
 * Antes é necessário criar uma semente aleatória por meio do método
 * rand.Seed(time.Now().UnixNano()) como se vê em inicializa
 */

func Aleatorio(abaixoDe int) int {
	return rand.Intn(abaixoDe)
}

// input: dif slice de int, nov número inteiro de uma caverna
// output: true se nov for um dos elementos do slice, false caso contrário
func DentroDe(dif []int, nov int) bool {
	for i := range dif {
		if dif[i] == nov {
			return true
		}
	}
	return false
}

// input: nenhum
// output: slice com 5 int aleatórios
// gera os números das cavernas do Wumpus, do Jogador e dos 3 abismos
func ParametroJogo() []int {
	diferentes := make([]int, 0)
	for len(diferentes) < 5 {
		novo := Aleatorio(20)
		if !DentroDe(diferentes, novo) {
			diferentes = append(diferentes, novo)
		}
	}
	return diferentes
}

func Inicializa() {
	rand.Seed(time.Now().UnixNano())
}
