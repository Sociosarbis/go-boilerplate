package white

func squareIsWhite(coordinates string) bool {
	if (coordinates[0]-'a')&1 == 0 {
		return !((coordinates[1]-'1')&1 == 0)
	} else {
		return (coordinates[1]-'1')&1 == 0
	}
}
