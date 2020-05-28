package determination

import "testing"

func TestCalculateNotMutants(t *testing.T) {

	notMutants := [][]string{
		{
			"ATGCGA",
			"ACGTGC",
			"ATATAT",
			"TGAAGG",
			"CGCCTA",
			"TCACTG",
		},
		{
			"ATGCGA",
			"ACGTGC",
			"ATATAT",
			"TGAAGG",
			"CGCCTG",
			"TCACTG",
		},
		{
			"ATGCGA",
			"ACGTGC",
			"ATATGT",
			"TGAACG",
			"CGCCTT",
			"TCACTG",
		},
	}
	for _, notMutant := range notMutants {
		result, _ := IsMutant(notMutant)
		if result == true {
			t.Error("DNA expected to be not mutant. ", notMutant)
		}
	}

}

func TestCalculateVerticalsMutants(t *testing.T) {

	mutants := [][]string{
		{
			"ATGCGA",
			"ACGTGC",
			"ATATCT",
			"AGAAGG",
			"TACCTA",
			"TCACTG",
		},
		{
			"ATGCGA",
			"ACGTGC",
			"TCATCT",
			"ACAAGG",
			"TCCCTA",
			"TGACTG",
		},
		{
			"ATGCGA",
			"ACGTGC",
			"TCATCT",
			"AGAAGG",
			"TTACTA",
			"TGACTG",
		},
		{
			"ATGCGA",
			"ACGTGC",
			"TCATCG",
			"AGTAGG",
			"TTGCTG",
			"TGACTG",
		},
	}
	for _, mutant := range mutants {
		result, _ := IsMutant(mutant)
		if result == false {
			t.Error("DNA expected to be mutant. ", mutant)
		}
	}

}

func TestCalculateLeftDiagonalMutants(t *testing.T) {

	mutants := [][]string{
		{
			"ATGCGA",
			"GAGTGC",
			"ATATAT",
			"TGAAGG",
			"CGCCTA",
			"TCACTG",
		},
		{
			"ATGCGA",
			"ACGTGC",
			"ATATAT",
			"TGAAGG",
			"CGCCAA",
			"TCACTA",
		},
		{
			"ATGCGA",
			"ACGTGC",
			"GTATAT",
			"TGAAGG",
			"CGGCTA",
			"TCAGTG",
		},
		{
			"ATGCGA",
			"ACTGGC",
			"ATATGT",
			"TGAAGG",
			"CGCCTA",
			"TCACTG",
		},
	}
	for _, mutant := range mutants {
		result, _ := IsMutant(mutant)
		if result == false {
			t.Error("DNA expected to be mutant. ", mutant)
		}
	}

}

func TestCalculateRigthDiagonalMutants(t *testing.T) {

	mutants := [][]string{
		{
			"ATGCGA",
			"ACGTAC",
			"ATTAAT",
			"TGACGG",
			"CGCCTA",
			"TCACTG",
		},
		{
			"ATGCGA",
			"ACGTGC",
			"ATATAA",
			"TGGAAG",
			"CGCATG",
			"TCACTG",
		},
		{
			"ATGGGA",
			"ACGTGC",
			"AGATGT",
			"GGAACG",
			"CGCCTT",
			"TCACTG",
		},
		{
			"ATGTGT",
			"ACGTGC",
			"ATAGGT",
			"GTGACG",
			"CGCCTT",
			"GCACTG",
		},
	}
	for _, mutant := range mutants {
		result, _ := IsMutant(mutant)
		if result == false {
			t.Error("DNA expected to be mutant. ", mutant)
		}
	}

}

func TestCalculateHorizontallMutants(t *testing.T) {

	mutants := [][]string{
		{
			"AAAAGA",
			"ACGTGC",
			"ATATAT",
			"TGAAGG",
			"CGCCTA",
			"TCACTG",
		},
		{
			"ATGCGA",
			"ACGTGC",
			"ATATAT",
			"TGAAGG",
			"CGCCTG",
			"TCAAAA",
		},
		{
			"ATGCGA",
			"ACGTGC",
			"AAAAGT",
			"TGAACG",
			"CGCCTT",
			"TCACTG",
		},
	}
	for _, mutant := range mutants {
		result, _ := IsMutant(mutant)
		if result == false {
			t.Error("DNA expected to be mutant. ", mutant)
		}
	}
}

func TestValidateDna(t *testing.T) {

	type validationTest struct {
		dna   []string
		valid bool
	}

	validationTests := []validationTest{
		validationTest{
			dna: []string{
				"AAAAGA",
				"ACGTGC",
				"ATATAT",
				"TGAAGG",
				"CGCCTA",
				"TCACTG",
			},
			valid: true,
		},
		validationTest{
			dna: []string{
				"AAAAGA",
				"ACGTGC",
				"ATATAT",
				"TGAAGG",
				"CGCCTA",
			},
			valid: false,
		},
		validationTest{
			dna: []string{
				"AAAAG",
				"ACGTGC",
				"ATATAT",
				"TGAAGG",
				"CGCCTA",
				"TCACTG",
			},
			valid: false,
		},
		validationTest{
			dna: []string{
				"ZAAAGA",
				"ACGTGC",
				"ATATAT",
				"TGAAGG",
				"CGCCTA",
				"TCACTG",
			},
			valid: false,
		},
	}
	for _, vt := range validationTests {
		_, error := IsMutant(vt.dna)

		if error != nil && vt.valid == true {
			t.Error("DNA expected to be valid. ", vt.dna)
		}

		if error == nil && vt.valid == false {
			t.Error("DNA expected to be invalid. ", vt.dna)
		}
	}

}

func BenchmarkCalculatorWorstCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsMutant([]string{
			"ATGCGA",
			"ACGTGC",
			"ATATAT",
			"TGAAGG",
			"CGCCAA",
			"TCACTA",
		})
	}
}
