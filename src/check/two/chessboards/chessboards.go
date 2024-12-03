package chessboards

func isBlack(s string) bool {
	return ((s[1]-'1')^(s[0]-'a'))&1 == 1
}

func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
	return isBlack(coordinate1) == isBlack(coordinate2)
}
