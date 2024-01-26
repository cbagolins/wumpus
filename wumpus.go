package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"wumpus/dicas"
	"wumpus/dodecaedro"
	"wumpus/jogo"
	"wumpus/lib"
	"wumpus/rotear"
)

func checaCavernaEExecuta(jogo jogo.Jogo, numeroCaverna int, invoca func()) {
	if jogo.VerificaDestino(numeroCaverna) == false {
		fmt.Println("Número inválido da caverna")
	} else {
		invoca()
	}
}

func main() {
	var comando string

	param := lib.ParametroJogo()
	lib.Inicializa()
	jogo := jogo.Jogo{
		Wumpus:  param[0],
		Jogador: param[1],
		Abismo:  param[2:],
	}
	jogo.Inicializa()
	dicas := dicas.Dicas{}
	dicas.New()
	// fmt.Println(jogo)
	for !jogo.FinalDoJogo() {
		jogo.MostraPosicao(&dicas)
		fmt.Println("Comando: [??] ou [C??]caminhar, [A??]atirar, [R??]rota, [T??]trajeto, [M]mapa ou [F]fim: ")
		fmt.Scanln(&comando)
		comando := strings.ToUpper(comando)

		/* Expressão Regular: \b[C|A|R|T|M|F]?[0-9]*\b */
		r, _ := regexp.Compile("\\b[C|A|R|T|M|F]?[0-9]*\\b")
		comando = r.FindString(comando)

		if comando == "" {
			fmt.Println("Comando não reconhecido!")
		} else if comando[:1] == "M" {
			jogo.MostraMapa()
		} else if comando[:1] == "C" {
			numeroCaverna, _ := strconv.Atoi(comando[1:])
			checaCavernaEExecuta(jogo, numeroCaverna, func() {
				jogo.MoveJogador(numeroCaverna)
			})
		} else if comando[:1] == "A" {
			numeroCaverna, _ := strconv.Atoi(comando[1:])
			checaCavernaEExecuta(jogo, numeroCaverna, func() {
				jogo.Atirar(numeroCaverna, &dicas)
			})
		} else if comando[:1] == "R" {
			numeroCaverna, _ := strconv.Atoi(comando[1:])
			if numeroCaverna < 0 || numeroCaverna >= dodecaedro.TamanhoDodecaedro() {
				fmt.Println("Número inválido da caverna")
			} else {
				r := rotear.Rotear{}
				if r.CalculaRota(jogo.Jogador, numeroCaverna, &dicas) {
					fmt.Println("Cuidado: pode haver perigos não mapeados!")
					fmt.Print("Rota ")
					r.ImprimeRota(numeroCaverna)
					fmt.Println()
				}
			}
		} else if comando[:1] == "T" {
			numeroCaverna, _ := strconv.Atoi(comando[1:])
			if jogo.VerificaDestino(numeroCaverna) {
				fmt.Println("Caverna próxima, use o comando: [C??]caminhar")
			} else {
				r := rotear.Rotear{}
				if r.CalculaRota(jogo.Jogador, numeroCaverna, &dicas) {
					jogo.MoveJogador(r.VeTrajeto(numeroCaverna))
				}
			}
		} else if comando[:1] == "F" {
			jogo.MostraPosicoesFinais()
		} else {
			numeroCaverna, _ := strconv.Atoi(comando)
			checaCavernaEExecuta(jogo, numeroCaverna, func() {
				jogo.MoveJogador(numeroCaverna)
			})
		}
	}
}
