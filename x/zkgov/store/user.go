package store

import (
	"bytes"
	"context"
	"errors"

	cosmosstore "cosmossdk.io/core/store"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
)

func InitUsers(ctx context.Context, store cosmosstore.KVStore, proposalID uint64) error {
	usersKey := types.UsersStoreKey(proposalID)

	usersBytes := []byte{}

	return store.Set(usersKey, usersBytes)

}

func StoreUser(ctx context.Context, store cosmosstore.KVStore, proposalID uint64, user string) error {
	usersKey := types.UsersStoreKey(proposalID)
	usersBytes, err := store.Get(usersKey)
	if err != nil {
		return err
	}

	userBytes := []byte(user)

	// if user already stored, throw error
	for i := 0; i < len(usersBytes); i += types.USER_SIZE {
		storedUser := usersBytes[i*types.USER_SIZE : (i+1)*types.USER_SIZE]
		if bytes.Equal(storedUser, userBytes) {
			return errors.New("user is already registered")
		}
	}

	usersBytes = append(usersBytes, userBytes...)
	store.Set(usersKey, usersBytes)

	return nil
}
