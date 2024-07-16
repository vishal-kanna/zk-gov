package circuit

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
)

type VoterInfo struct {
	ProposalID    uint64 `json:"proposal_id"`
	Commitment    string `json:"commitment"`
	Nullifier     string `json:"nullifier"`
	RandomSecret1 uint64 `json:"random_secret1"`
	RandomSecret2 uint64 `json:"random_secret2"`
	VoteOption    uint64 `json:"vote_option"`
}

// SaveInfo saves the voter info as JSON
func SaveInfo(proposalID uint64, commitment []byte, nullifier []byte, voteOption uint64) error {
	commitmentString := types.BytesToHexString(commitment)
	nullifierString := types.BytesToHexString(nullifier)
	info := VoterInfo{
		ProposalID: proposalID,
		Commitment: commitmentString,
		Nullifier:  nullifierString,
		VoteOption: voteOption,
	}
	data, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join("commitments", fmt.Sprint(proposalID)+".json"), data, 0666)
}

// FetchInfo fetches the voter info from JSON
func FetchInfo(proposalID string) (VoterInfo, error) {
	var info VoterInfo
	data, err := os.ReadFile(filepath.Join("commitments", proposalID+".json"))
	if err != nil {
		return info, err
	}
	err = json.Unmarshal(data, &info)
	return info, err
}

// func main() {
// 	// Example usage
// 	err := SaveInfo(12345, []byte("commitment123"), []byte("nullifier123"))
// 	if err != nil {
// 		fmt.Println("Error saving info:", err)
// 	}
// 	info, err := FetchInfo("12345")
// 	if err != nil {
// 		fmt.Println("Error fetching info:", err)
// 	} else {
// 		fmt.Printf("Fetched info: %+v\n", info)
// 	}
// }
