package rotear

import (
	"fmt"
	"wumpus/dicas"
	"wumpus/dodecaedro"
)

type rota struct {
	passo      int
	pontoFinal bool
	anterior   int
}

type Rotear struct {
	caverna [20]rota
}

// input: número da caverna destino final desta rota
// output: boolean informando se existe movimento a ser feito
// calcula e mostra uma rota que leve a uma determinada caverna
func (r *Rotear) CalculaRota(origem int, destino int, d *dicas.Dicas) bool {
	// inicia criando base de dados para cálculo do melhor caminho
	for i := range r.caverna {
		r.caverna[i] = rota{passo: -1, pontoFinal: d.IsPontoFinal(i), anterior: -1}
	}
	// agora marca o passo 0
	r.caverna[origem].passo = 0
	i := 0
	for r.ExecutaPasso(i) {
		i++
	}

	if origem == destino {
		fmt.Println("Você já está na caverna ", destino)
		return false
	} else if r.caverna[destino].passo == -1 {
		fmt.Println("Não existe caminho conhecido para ", destino)
		return false
	} else {
		return true
	}
}

// input: quantidade de passos já efetuada
// output: boolean indicando se há necessidade de mais passos
// Este procedimento verifica para todas as cavernas que estão a
// "estePasso" da caverna onde se encontra o jogador, quais são
// as cavernas mais próximas
func (r *Rotear) ExecutaPasso(estePasso int) bool {
	novoPasso := false
	for i := range r.caverna {
		if r.caverna[i].passo != estePasso {
			continue
		}
		if r.caverna[i].pontoFinal {
			continue
		}
		lista := dodecaedro.Proximos(i)
		for _, k := range lista {
			if r.caverna[k].passo == -1 {
				r.caverna[k].passo = estePasso + 1
				r.caverna[k].anterior = i
				novoPasso = true
			}
		}
	}
	return novoPasso
}

// input: número da caverna de destino
// output: nenhum
// imprime recursivamente a rota para a caverna de "destino"
func (r *Rotear) ImprimeRota(destino int) {
	if r.caverna[destino].passo > 1 {
		r.ImprimeRota(r.caverna[destino].anterior)
	}
	fmt.Print(" -> ", destino)
}

// input: número da caverna de destino
// output: próximo número da caverna a caminho do destino
// procura passo a passo o caminho para algum destino desejado
func (r *Rotear) VeTrajeto(des int) int {
	cavAnterior := r.caverna[des].anterior
	for r.caverna[cavAnterior].passo > 1 {
		cavAnterior = r.caverna[cavAnterior].anterior
	}
	return cavAnterior
}
