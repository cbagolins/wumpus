package jogo

import (
	"fmt"
	"wumpus/dicas"
	"wumpus/dodecaedro"
	"wumpus/lib"
)

type Jogo struct {
	Wumpus      int
	Jogador     int
	Abismo      []int
	fimJogo     bool
	percorridos []int
	// descobertas [20]int
	// rota        [20]veRota
}

// input: alvo como um int
// atira, se acertar no wumpus, informa que ganhou o jogo
// se errou o alvo, o wumpus se move e pode te papar
func (j *Jogo) Atirar(alvo int, d *dicas.Dicas) {
	if j.Wumpus == alvo {
		fmt.Println("Acertou no Wumpus ...... Ganhou!")
		j.MostraPosicoesFinais()
		j.fimJogo = true
	} else {
		fmt.Println("Errou o tiro. Wumpus se move ......")
		d.ResetaDicas()
		j.Wumpus = dodecaedro.MoveProximo(j.Wumpus)
		if j.Jogador == j.Wumpus {
			fmt.Println("O Wumpus te papou ...... Perdeu!")
			j.MostraPosicoesFinais()
			j.fimJogo = true
		}
	}
}

// apenas sinaliza se o jogo acabou
// output: true ou false
func (j *Jogo) FinalDoJogo() bool {
	return j.fimJogo
}

// inicio do jogo, coloca a caverna onde está o jogador
// na lista de cavernas percorridas
func (j *Jogo) Inicializa() {
	j.fimJogo = false
	j.percorridos = append(j.percorridos, j.Jogador)
}

// mostra as conexões de todas as cavernas para permitir que o jogador
// planeje para onde vai
func (j *Jogo) MostraMapa() {
	for i := 0; i < 20; i++ {
		cavernasProximas := dodecaedro.Proximos(i)
		fmt.Print("Caverna: ", i)
		fmt.Print(" vai a ", cavernasProximas)
		if i%2 == 0 {
			fmt.Print(" - ")
		} else {
			fmt.Println()
		}
	}
}

// mostra a posição do jogador e informa o que está nas proximidades
// sem, naturalmente, dizer em que local estão os perigos.
// é possível sentir o vazio do abismo, bem como sentir o cheiro do wumpus
func (j *Jogo) MostraPosicao(d *dicas.Dicas) {
	perigo := false
	avisoPerigo := ""
	cavernasProximas := dodecaedro.Proximos(j.Jogador)

	for _, i := range j.Abismo {
		if lib.DentroDe(cavernasProximas, i) {
			perigo = true
			if avisoPerigo == "" {
				avisoPerigo = "Abismo"
			} else {
				avisoPerigo += "/Abismo"
			}
		}
	}

	if lib.DentroDe(cavernasProximas, j.Wumpus) {
		perigo = true
		if avisoPerigo == "" {
			avisoPerigo = "Wumpus"
		} else {
			avisoPerigo += "/Wumpus"
		}
	}

	d.MarcaDicas(avisoPerigo, cavernasProximas, j.Jogador)
	fmt.Printf("Sua posição atual é na caverna: %d\n", j.Jogador)
	fmt.Println("As cavernas próximas são: ", cavernasProximas)
	fmt.Println("As cavernas visitadas são: ", j.percorridos)
	d.MostraDicas()

	if perigo {
		fmt.Printf("ATENÇÃO: %s\n", avisoPerigo)
	} else {
		fmt.Println("ATENÇÃO: Nenhum perigo nas proximidades!")
	}
}

// input: nenhum
// output: nenhum
// mostra as posições dos abismos, bem como a posição final do wumpus
func (j *Jogo) MostraPosicoesFinais() {
	fmt.Print("Wumpus: ", j.Wumpus)
	fmt.Println(", Abismos:", j.Abismo)
	j.fimJogo = true
}

// move o jogador para a caverna des
// se cair num abismo, informa
// se entrar na caverna do wumpus, também informa
func (j *Jogo) MoveJogador(des int) {
	j.Jogador = des
	j.percorridos = append(j.percorridos, des)
	for _, i := range j.Abismo {
		if i == j.Jogador {
			fmt.Println("Caiu no abismoooooooooo ...... Perdeu!")
			j.MostraPosicoesFinais()
			j.fimJogo = true
		}
	}
	if j.Wumpus == j.Jogador {
		fmt.Println("O Wumpus te papou ...... Perdeu!")
		j.MostraPosicoesFinais()
		j.fimJogo = true
	}
}

// input: número da caverna de destino do jogador
// output: true ou false se o jogador pode se mover àquela caverna
// verifica se o jogador pode se mover a caverna des
func (j *Jogo) VerificaDestino(destino int) bool {
	cavernasProximas := dodecaedro.Proximos(j.Jogador)
	return lib.DentroDe(cavernasProximas, destino)
}
