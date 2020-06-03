package storage

import (
	"context"
	"fmt"
	"testing"
)

func TestDna_SetHash(t *testing.T) {
	dna := Dna{
		Matrix: []string{
			"ATGCGA",
			"ACGTGC",
			"ATATAT",
			"TGAAGG",
			"CGCCTA",
			"TCACTG",
		},
	}
	dna.setHash()
	if dna.Hash == "" {
		t.Error("The hash must be not empty hash.")
	}
}

func TestDna_Save(t *testing.T) {
	dna := Dna{
		Matrix: []string{
			"ATGCGA",
			"ACGTGC",
			"ATATAT",
			"TGAAGG",
			"CGCCTA",
			"TCACTG",
		},
		Hash:     "alto hash",
		IsMutant: true,
	}
	expected := "{true alto hash [ATGCGA ACGTGC ATATAT TGAAGG CGCCTA TCACTG]}"
	mc := mongoClientMock{}
	saved, _ := dna.Save(context.Background(), &mc)

	strSaved := fmt.Sprintf("%v", saved)

	if strSaved != expected {
		t.Errorf("Saved %v , expected %v", strSaved, expected)
	}
}

func TestStoreDna(t *testing.T) {
	expected := "{true 1f5ec16cdac407ca358686c95ca66a66 [ATGCGA ACGTGC ATATAT TGAAGG CGCCTA TCACTG]}"
	mc := mongoClientMock{}
	matrix := []string{
		"ATGCGA",
		"ACGTGC",
		"ATATAT",
		"TGAAGG",
		"CGCCTA",
		"TCACTG",
	}

	saved, _ := StoreDna(&mc, matrix, true)

	strSaved := fmt.Sprintf("%v", saved)

	if strSaved != expected {
		t.Errorf("Saved %v , expected %v", strSaved, expected)
	}
}
