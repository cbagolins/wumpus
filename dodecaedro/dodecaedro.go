package dodecaedro

import (
	"wumpus/lib"
)

var mapaDodecaedro = [20][3]int{{1, 4, 5}, {0, 2, 7}, {1, 3, 9}, {2, 4, 11}, {0, 3, 13},
	{0, 6, 14}, {5, 7, 15}, {1, 6, 8}, {7, 9, 16}, {2, 8, 10},
	{9, 11, 17}, {3, 10, 12}, {11, 13, 18}, {4, 12, 14}, {5, 13, 19},
	{6, 16, 19}, {8, 15, 17}, {10, 16, 18}, {12, 17, 19}, {14, 15, 18}}

func Proximos(posicao int) []int {
	// o [:] converte de array [3]int para slice [:]int
	return mapaDodecaedro[posicao][:]
}

func MoveProximo(posicao int) int {
	novo := lib.Aleatorio(3)
	return mapaDodecaedro[posicao][novo]
}

func TamanhoDodecaedro() int {
	return len(mapaDodecaedro)
}
