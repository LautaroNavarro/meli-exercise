package determination

import (
	"errors"
	"fmt"
	"strings"
)

func validateDna(dna []string) error {

	fmt.Println("Length :", len(dna))
	if len(dna) != 6 {
		return errors.New("Invalid DNA. Dna must be a Matrix 6x6")
	}
	for i := 0; i < len(dna); i++ {
		if len(dna[i]) != 6 {
			return errors.New("Invalid DNA. Dna must be a Matrix 6x6")
		}
		for j := 0; j < len(dna[i]); j++ {
			valid := false
			for _, char := range []byte{65, 84, 67, 71} {
				if char == dna[i][j] {
					valid = true
				}
			}
			if !valid {
				return errors.New("Invalid DNA. Matrix elements must be one of [A, T, C, G]")
			}
		}
	}

	return nil
}

// IsMutant returns true if a given dna is mutant, else returns false
func IsMutant(dna []string) (bool, error) {

	if err := validateDna(dna); err != nil {
		fmt.Println(err)
		return false, err
	}

	if checkHorizontal(dna) {
		return true, nil
	}

	if checkVertical(dna) {
		return true, nil
	}

	if checkDiagonal(dna) {
		return true, nil
	}

	return false, nil
}

func stringIsMutant(str string) bool {
	for _, char := range []byte{65, 84, 67, 71} {
		counter := 0
		for i := 0; i < len(str); i++ {
			if str[i] == char {
				counter++
				if counter >= 4 {
					return true
				}
			} else {
				counter = 0
			}
		}
	}
	return false
}

func checkHorizontal(dna []string) bool {
	for _, row := range dna {
		if stringIsMutant(row) {
			return true
		}
	}
	return false
}

func checkDiagonal(dna []string) bool {
	var builder strings.Builder

	for x := 0; x <= 2; x++ {
		for y := 0; y <= 2; y++ {
			if x == 0 || y == 0 {
				x2 := x
				y2 := y
				for x2 != 6 && y2 != 6 {
					builder.WriteByte(dna[x2][y2])
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
					builder.WriteByte(dna[x2][y2])
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

func checkVertical(dna []string) bool {
	var builder strings.Builder
	for y := 0; y < 6; y++ {
		for x := 0; x < 6; x++ {
			builder.WriteByte(dna[x][y])
		}
		if stringIsMutant(builder.String()) {
			return true
		}
		builder.Reset()
	}
	return false
}
