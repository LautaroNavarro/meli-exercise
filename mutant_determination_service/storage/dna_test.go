package storage

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoIndexMock struct{}

func (mg mongoIndexMock) CreateOne(ctx context.Context, model mongo.IndexModel) (string, error) {
	return "", nil
}

type mongoClientMock struct{}

func (mc mongoClientMock) Database(name string, opts ...*options.DatabaseOptions) *mongo.Database {
	return &mongo.Database{}
}

func (mc mongoClientMock) Connect(ctx context.Context) error {
	return nil
}

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

// func TestDna_Save(t *testing.T) {
// 	dna := Dna{
// 		Matrix: []string{
// 			"ATGCGA",
// 			"ACGTGC",
// 			"ATATAT",
// 			"TGAAGG",
// 			"CGCCTA",
// 			"TCACTG",
// 		},
// 		Hash:     "alto hash",
// 		IsMutant: true,
// 	}
// 	expected := "[123 34 105 115 95 109 117 116 97 110 116 34 58 116 114 117 101 44 34 104 97 115 104 34 58 34 97 108 116 111 32 104 97 115 104 34 44 34 109 97 116 114 105 120 34 58 91 34 65 84 71 67 71 65 34 44 34 65 67 71 84 71 67 34 44 34 65 84 65 84 65 84 34 44 34 84 71 65 65 71 71 34 44 34 67 71 67 67 84 65 34 44 34 84 67 65 67 84 71 34 93 125]"
// 	mc := mongoClientMock{}
// 	saved, _ := dna.Save(context.Background(), mc)

// 	strSaved := fmt.Sprintf("%v", saved.InsertedID)

// 	if strSaved != expected {
// 		t.Errorf("Saved %v , expected %v", strSaved, expected)
// 	}
// }

// func TestStoreDna(t *testing.T) {
// 	expected := "[123 34 105 115 95 109 117 116 97 110 116 34 58 116 114 117 101 44 34 104 97 115 104 34 58 34 49 102 53 101 99 49 54 99 100 97 99 52 48 55 99 97 51 53 56 54 56 54 99 57 53 99 97 54 54 97 54 54 34 44 34 109 97 116 114 105 120 34 58 91 34 65 84 71 67 71 65 34 44 34 65 67 71 84 71 67 34 44 34 65 84 65 84 65 84 34 44 34 84 71 65 65 71 71 34 44 34 67 71 67 67 84 65 34 44 34 84 67 65 67 84 71 34 93 125]"
// 	mc := mongoClientMock{}
// 	matrix := []string{
// 		"ATGCGA",
// 		"ACGTGC",
// 		"ATATAT",
// 		"TGAAGG",
// 		"CGCCTA",
// 		"TCACTG",
// 	}

// 	saved, _ := StoreDna(mc, matrix, true)

// 	strSaved := fmt.Sprintf("%v", saved.InsertedID)

// 	if strSaved != expected {
// 		t.Errorf("Saved %v , expected %v", strSaved, expected)
// 	}
// }
