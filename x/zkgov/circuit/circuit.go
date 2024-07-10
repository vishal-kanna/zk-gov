package circuit

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/accumulator/merkle"
	"github.com/consensys/gnark/std/math/emulated"
	"github.com/consensys/gnark/std/signature/ecdsa"
)

// Define the circuit
type PrivateVotingCircuit[T, S emulated.FieldParams] struct {
	SecretUniqueId1 frontend.Variable // randomly generated
	SecretUniqueId2 frontend.Variable // randomly generated

	Commitment frontend.Variable //  hash(secret1 + secret2 + pubkey)
	Nullifier  frontend.Variable `gnark:",public"` // hash(secret2 + pubkey)
	// VoteOption       frontend.Variable `gnark:",public"`
	Signature        ecdsa.Signature[S]
	Message          emulated.Element[S] `gnark:",public"` // hash(nullifier and vote option)
	Publickey        ecdsa.PublicKey[T, S]
	MerkleProofIndex frontend.Variable // commitment index in the list of commitments stored on chain
	MerkleProof      merkle.MerkleProof
	MerkleRoot       frontend.Variable `gnark:",public"`
}

func (circuit *PrivateVotingCircuit[T, S]) Define(api frontend.API) error {

	// TODO: circuit logic
	return nil
}

func PreparePublicWitness(nullifier string, voteOption uint64, merkleRoot string) witness.Witness {
	nullifierBytes := []byte(nullifier)
	merkleRootBytes := []byte(merkleRoot)

	voteOptionBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(voteOptionBytes, voteOption)
	message := append(nullifierBytes, merkleRootBytes...)
	message = append(message, voteOptionBytes...)
	messageHash := Sha256Hash(message)

	field := ecc.BN254.ScalarField()
	publicWitness, err := ConstructWitness(field, merkleRootBytes, nullifierBytes, messageHash)
	if err != nil {
		panic(err)
	}

	return publicWitness

}

// constructs new public witness using assignment's public inputs
func ConstructWitness(field *big.Int, merkleRootBytes []byte, nullifierBytes []byte, message []byte) (witness.Witness, error) {
	newWitness, err := witness.New(field)
	if err != nil {
		return nil, err
	}

	witnessChan := make(chan any)
	go passPubInputs(&witnessChan, merkleRootBytes, nullifierBytes, message)
	newWitness.Fill(3, 0, witnessChan)

	return newWitness, nil
}

// close the channel after passing the values to end the for loop over channel values
func passPubInputs(witnessChan *chan any, merkleRootBytes []byte, nullifierBytes []byte, message []byte) {
	*witnessChan <- nullifierBytes
	*witnessChan <- message
	*witnessChan <- merkleRootBytes

	fmt.Println("pulbic values sent via channel for witness construction...")
	close(*witnessChan)
}

func Sha256Hash(message []byte) []byte {
	shaFunc := sha256.New()
	shaFunc.Write(message)
	return shaFunc.Sum(nil)
}
