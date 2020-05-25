package utils

import (
	"errors"
	"strings"
)

func validateAdn(adn []string) error {

	if len(adn) != 6 {
		return errors.New("Invalid adn length")
	}
	for i := 0; i < len(adn); i++ {
		if len(adn[i]) != 6 {
			return errors.New("Invalid adn length")
		}
		for j := 0; j < len(adn[i]); j++ {
			valid := false
			for _, char := range []byte{65, 84, 67, 71} {
				if char == adn[i][j] {
					valid = true
				}
			}
			if !valid {
				return errors.New("Invalid adn char")
			}
		}
	}

	return nil
}

// IsMutant returns true if a given adn is mutant, else returns false
func IsMutant(adn []string) (bool, error) {

	if err := validateAdn(adn); err != nil {
		return false, err
	}

	if checkHorizontal(adn) {
		return true, nil
	}

	if checkVertical(adn) {
		return true, nil
	}

	if checkDiagonal(adn) {
		return true, nil
	}

	return false, nil
}

func stringIsMutant(str string) bool {
	for _, searchString := range []string{"AAAA", "CCCC", "TTTT", "GGGG"} {
		if strings.Count(str, searchString) >= 1 {
			return true
		}
	}
	return false
}

func checkHorizontal(adn []string) bool {
	for _, row := range adn {
		if stringIsMutant(row) {
			return true
		}
	}
	return false
}

func checkDiagonal(adn []string) bool {
	var builder strings.Builder

	for x := 0; x <= 2; x++ {
		for y := 0; y <= 2; y++ {
			if x == 0 || y == 0 {
				x2 := x
				y2 := y
				for x2 != 6 && y2 != 6 {
					builder.WriteByte(adn[x2][y2])
					x2++
					y2++
				}
				if stringIsMutant(builder.String()) {
					return true
				}
				builder.Reset()
			}
		}
	}

	for x := 0; x <= 2; x++ {
		for y := 3; y <= 5; y++ {
			if x == 0 || y == 5 {
				x2 := x
				y2 := y
				for x2 != 6 && y2 != -1 {
					builder.WriteByte(adn[x2][y2])
					x2++
					y2--
				}
				if stringIsMutant(builder.String()) {
					return true
				}
				builder.Reset()
			}
		}
	}

	return false
}

func checkVertical(adn []string) bool {
	var builder strings.Builder
	for y := 0; y < 6; y++ {
		for x := 0; x < 6; x++ {
			builder.WriteByte(adn[x][y])
		}
		if stringIsMutant(builder.String()) {
			return true
		}
		builder.Reset()
	}
	return false
}
