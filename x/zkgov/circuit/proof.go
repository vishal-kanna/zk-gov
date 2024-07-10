package circuit

import (
	"encoding/json"

	"github.com/consensys/gnark/backend/groth16"
	bn254 "github.com/consensys/gnark/backend/groth16/bn254"
)

func UnMarshalZkProof(zkProofBytes []byte) (groth16.Proof, error) {

	// unmarshal sig into proof
	zkProof := new(bn254.Proof)
	err := json.Unmarshal(zkProofBytes, zkProof)
	if err != nil {
		return nil, err
	}

	return zkProof, nil
}
