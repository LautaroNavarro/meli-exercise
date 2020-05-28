package determination

// Dna contains a matrix []string which represents a dna
type Dna struct {
	Matrix []string `json:"dna" binding:"required"`
}
