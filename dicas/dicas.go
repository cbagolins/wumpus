package dicas

import (
	"fmt"
	"strings"
)

type Dicas struct {
	perigos  [20]int
	conteudo []string
}

func (d *Dicas) New() {
	for i := range d.perigos {
		d.perigos[i] = -1
	}
	d.conteudo = make([]string, 0)
	d.conteudo = append(d.conteudo, "Seguro")
}

func (d *Dicas) IsPontoFinal(i int) bool {
	return d.perigos[i] > 0
}

func (d *Dicas) melhorAviso(a, b int) int {
	avisoA := d.conteudo[a]
	if a == 0 || b == 0 {
		return 0
	} else if a == b {
		return a
	} else if strings.Contains(avisoA, "/") {
		return b
	}
	return a
}

func (d *Dicas) MarcaDicas(avisoPerigo string, cavernas []int, posicaoAtual int) {
	/* você está em caverna segura, senão o jogo teria acabado */
	d.perigos[posicaoAtual] = 0
	/* se não houver perigo */
	if avisoPerigo == "" {
		for _, i := range cavernas {
			d.perigos[i] = 0
		}
	} else {
		pos := -1
		/* caso haja perigo, verifica se já encontrou este perigo */
		for i, a := range d.conteudo {
			if a == avisoPerigo {
				pos = i
				break
			}
		}
		/* primeira vez do perigo, acrescentar à lista */
		if pos == -1 {
			pos = len(d.conteudo)
			d.conteudo = append(d.conteudo, avisoPerigo)
		}
		/* marca a situação com o melhor aviso */
		for _, i := range cavernas {
			if d.perigos[i] == -1 {
				d.perigos[i] = pos
			} else {
				d.perigos[i] = d.melhorAviso(d.perigos[i], pos)
			}
		}
	}
}

func (d *Dicas) MostraDicas() {
	for i, s := range d.conteudo {
		imprimiu := false
		for j, k := range d.perigos {
			if i == k {
				if imprimiu == false {
					fmt.Printf("Aviso: %s", s)
					imprimiu = true
				}
				fmt.Printf(" %d", j)
			}
		}
		if imprimiu == true {
			fmt.Println()
		}
	}
}

// se o jogador atirar e errar apaga todas as dicas porque o wumpus se movimentou
func (d *Dicas) ResetaDicas() {
	d.New()
}
