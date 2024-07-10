package circuit

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
)

func UnMarshalZkProof(zkProofBytes []byte) groth16.Proof {
	curveID := ecc.BN254
	zkProof := groth16.NewProof(curveID)

	// TODO: use proof bytes
	return zkProof
}
