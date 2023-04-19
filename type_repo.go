package ghid

import (
	"fmt"
	"strconv"

	"github.com/vmihailenco/msgpack/v5"
)

func init() {
	RegisterDecodeV1(TypeRepository, func(key []byte) (KeyV1, error) {
		id, err := strconv.ParseUint(string(key), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to decode repo id: %w", err)
		}
		return RepoKey{ID: RepoID(id)}, nil
	})
	RegisterDecodeV2(TypeRepository, func(key msgpack.RawMessage) (KeyV2, error) {
		id, err := decodeV2Uint(TypeRepository, 0, key)
		if err != nil {
			return nil, err
		}
		return RepoKey{ID: RepoID(id)}, nil
	})
}

// RepoID corresponds to databaseId field of the Repository node.
//
// See https://docs.github.com/en/graphql/reference/objects#repository.
type RepoID uint64

var (
	_ KeyV1 = RepoKey{}
	_ KeyV2 = RepoKey{}
)

// RepoKey is a unique key for Repository nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#repository.
type RepoKey struct {
	ID RepoID // corresponds to databaseId
}

// Type implements Key.
func (r RepoKey) Type() string {
	return TypeRepository
}

// GetRepoID implements KeyWithRepo.
func (r RepoKey) GetRepoID() RepoID {
	return r.ID
}

// KeyV1 implements KeyV1.
func (r RepoKey) KeyV1() string {
	return strconv.FormatUint(uint64(r.ID), 10)
}

// KeyV2 implements KeyV2.
func (r RepoKey) KeyV2() msgpack.RawMessage {
	return mustEncodeV2([]any{uint(0), uint(r.ID)})
}
