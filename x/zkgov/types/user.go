package types

import (
	"bytes"
	"context"
	"errors"

	cosmosstore "cosmossdk.io/core/store"
)

func StoreUser(ctx context.Context, store cosmosstore.KVStore, proposalID uint64, user string) error {
	usersKey := UsersStoreKey(proposalID)
	usersBytes, err := store.Get(usersKey)
	if err != nil {
		return err
	}

	userBytes := []byte(user)

	// if user already stored, throw error
	for i := 0; i < len(usersBytes); i += USER_SIZE {
		storedUser := usersBytes[i*USER_SIZE : (i+1)*USER_SIZE]
		if bytes.Equal(storedUser, userBytes) {
			return errors.New("user is already registered")
		}
	}

	usersBytes = append(usersBytes, userBytes...)
	store.Set(usersKey, usersBytes)

	return nil
}
