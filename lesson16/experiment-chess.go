package main

import "fmt"

type pieces struct {
	king rune
	queen rune
	bishop rune
	knight rune
	rook rune
	pawn rune
}

func main() {
	/**
	r n b q k b n r | 114 110 98 113 107
	p p p p p p p p | 112


	P P P P P P P P | 80
	R N B Q K B N R | 82 78 66 81 75
	*/
	black := pieces{
		king: 107,
		queen: 113,
		bishop: 98,
		knight: 110,
		rook: 114,
		pawn: 112,
	}

	white := pieces{
		king: 75,
		queen: 81,
		bishop: 66,
		knight: 78,
		rook: 82,
		pawn: 80,
	}

	var board [8][8]rune

	board[0][0], board[0][7] = black.rook, black.rook
	board[0][1], board[0][6] = black.knight, black.knight
	board[0][2], board[0][5] = black.bishop, black.bishop
	board[0][3] = black.queen
	board[0][4] = black.king

	for column := range board[1] {
		board[1][column] = black.pawn
	}

	board[7][0], board[7][7] = white.rook, white.rook
	board[7][1], board[7][6] = white.knight, white.knight
	board[7][2], board[7][5] = white.bishop, white.bishop
	board[7][3] = white.queen
	board[7][4] = white.king

	for column := range board[6] {
		board[6][column] = white.pawn
	}

	fmt.Printf("%c\n", board[0])
	fmt.Printf("%c\n", board[1])
	fmt.Println(board[2])
	fmt.Println(board[3])
	fmt.Println(board[4])
	fmt.Println(board[5])
	fmt.Printf("%c\n", board[6])
	fmt.Printf("%c\n", board[7])
}
