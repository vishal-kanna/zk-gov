package keeper

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math/big"
	"time"

	"math/rand"

	cosmosstore "cosmossdk.io/core/store"
	"github.com/consensys/gnark-crypto/accumulator/merkletree"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/accumulator/merkle"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/circuit"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
)

type Keeper struct {
	storeKey cosmosstore.KVStoreService
	cdc      codec.BinaryCodec
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey cosmosstore.KVStoreService,
) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}
func init() {

}
func (k *Keeper) StoreCommitment(ctx context.Context, commitment string) error {

	store := k.storeKey.OpenKVStore(ctx)
	seq, err := k.nextSequence(ctx, CommitmentSeqPrefix)
	if err != nil {
		return err
	}
	commitments := types.Commitment{
		Commitment:   commitment,
		CommitmentId: seq,
	}
	bz, err := k.cdc.Marshal(&commitments)
	if err != nil {
		return err
	}
	return store.Set(CommitmentStoreKey(seq), bz)
}
func (k *Keeper) RegisterUsr(ctx context.Context) error {
	store := k.storeKey.OpenKVStore(ctx)
	userId, err := k.nextSequence(ctx, UserSeqPrefix)
	if err != nil {
		return err
	}
	fmt.Println("The userid is>>>>>>>>>>>", userId)
	randomNumer := getRandomNumber()
	commitment, nullifier := createCommitmentAndNullifier(int64(userId), randomNumer)
	user := types.User{
		Userid:       userId,
		Commitment:   commitment,
		Nullifier:    nullifier,
		RandomNumber: uint64(randomNumer),
	}
	bz, err := k.cdc.Marshal(&user)
	if err != nil {
		return err
	}
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<")
	return store.Set(UserStoreKey(userId), bz)
}
func (k *Keeper) nextSequence(ctx context.Context, key []byte) (uint64, error) {
	store := k.storeKey.OpenKVStore(ctx)
	found, err := store.Has(key)
	fmt.Println("The found value is>>>>>>>>>>.", found)
	if err != nil {
		return 0, err
	}
	var seq uint64 = 1
	if found {
		pvBytes, err := store.Get(key)
		if err != nil {
			return 0, err
		}
		seq = binary.BigEndian.Uint64(pvBytes) + 1
	}
	seqBytes := uint64ToBytes(seq)
	store.Set(key, seqBytes)
	return seq, nil
}

func (k *Keeper) GetUserInfo(ctx context.Context, userid int64) (types.User, error) {
	store := k.storeKey.OpenKVStore(ctx)
	bz, err := store.Get(UserStoreKey(uint64(userid)))
	if err != nil {
		return types.User{}, err
	}
	fmt.Println("Im being called >>>>>>>>>>>>>>>>")
	var usr types.User
	err = k.cdc.Unmarshal(bz, &usr)
	if err != nil {
		return types.User{}, err
	}
	return usr, nil
}

func (k *Keeper) SetSequence(ctx context.Context, seq uint64) {
	store := k.storeKey.OpenKVStore(ctx)
	seqBytes := uint64ToBytes(seq)
	store.Set(UserSeqPrefix, seqBytes)
}
func uint64ToBytes(value uint64) []byte {
	seqBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(seqBytes, value)
	return seqBytes
}

// Generate a random 5-digit salt
func getRandomNumber() int64 {
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	return int64(rng.Intn(10000))
}

// To create a commitment and nullifier
func createCommitmentAndNullifier(userId, randomId int64) ([]byte, []byte) {
	hFunc := mimc.NewMiMC()

	// Create commitment
	hFunc.Write(big.NewInt(userId).Bytes())
	hFunc.Write(big.NewInt(randomId).Bytes())
	commitment := hFunc.Sum(nil)
	hFunc.Reset()

	// Create nullifier
	hFunc.Write(big.NewInt(randomId).Bytes())
	nullifier := hFunc.Sum(nil)

	return commitment, nullifier
}

// to get merkle proof from root+proof bytes
func GetMerkleProofFromBytes(rootBytes []byte, proofBytes [][]byte) merkle.MerkleProof {
	var merkleProof merkle.MerkleProof
	merkleProof.RootHash = rootBytes
	merkleProof.Path = make([]frontend.Variable, len(proofBytes))
	for i := 0; i < len(proofBytes); i++ {
		merkleProof.Path[i] = proofBytes[i]
	}
	return merkleProof
}

func (k Keeper) ProofGeneration(ctx context.Context, userid uint64) {
	store := k.storeKey.OpenKVStore(ctx)
	bz, err := store.Get(UserStoreKey(uint64(userid)))
	if err != nil {
		fmt.Println("Err", err)
	}
	var usr types.User
	err = k.cdc.Unmarshal(bz, &usr)
	if err != nil {
		fmt.Println("Err", err)
	}
	commitment := usr.Commitment
	nullifier := usr.Nullifier
	randomNumber := usr.RandomNumber
	var buf bytes.Buffer
	// build merkle proof
	dataSegments := 4
	proofIndex := 0
	dataSize := len(commitment)
	hFunc := mimc.NewMiMC()
	for j := byte(1); j <= byte(dataSegments); j++ {
		data := commitment
		hFunc.Reset()
		hFunc.Write(data)
		hash := hFunc.Sum(nil)

		_, err := buf.Write(hash)
		if err != nil {
			fmt.Println("failed to write hash", err)
		}
	}
	root, proof, numLeaves, err := merkletree.BuildReaderProof(&buf, hFunc, dataSize, uint64(proofIndex))
	verified := merkletree.VerifyProof(hFunc, root, proof, uint64(proofIndex), numLeaves)
	if verified {
		fmt.Println("Proof is generated and verified")
	}
	// Define the inputs
	assignment := circuit.Circuit{
		UniqueId1:   usr.Userid,
		ProofIndex:  0,
		UniqueId2:   randomNumber,
		Commitment:  commitment,
		Nullifier:   nullifier,
		MerkleProof: GetMerkleProofFromBytes(root, proof),
		MerkleRoot:  root,
		VoteOption:  1,
	}
	witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	publicWitness, err := witness.Public()
	fmt.Println("pub", publicWitness)
}
