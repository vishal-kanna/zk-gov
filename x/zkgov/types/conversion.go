package types

import (
	"encoding/hex"
	"fmt"
)

func HexStringToBytes(data string) ([]byte, error) {
	bytes, err := hex.DecodeString(data)
	if err != nil {
		return nil, fmt.Errorf("invalid hex string: %v", err)
	}
	return bytes, nil
}

func BytesToHexString(data []byte) string {
	return hex.EncodeToString(data)
}
