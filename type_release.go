package ghid

import (
	"fmt"
	"strconv"

	"github.com/vmihailenco/msgpack/v5"
)

func init() {
	RegisterDecodeV1(TypeRelease, func(key []byte) (KeyV1, error) {
		id, err := strconv.ParseUint(string(key), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to decode release id: %w", err)
		}
		return ReleaseKeyV1{ID: id}, nil
	})
	RegisterDecodeV2(TypeRelease, func(key msgpack.RawMessage) (KeyV2, error) {
		arr, err := decodeV2UintArr(TypeRelease, 0, 2, key)
		if err != nil {
			return nil, err
		}
		return ReleaseKeyV2{RepoID: RepoID(arr[0]), ID: arr[1]}, nil
	})
}

var (
	_ KeyV1 = ReleaseKeyV1{}
	_ KeyV2 = ReleaseKeyV2{}
)

// ReleaseKeyV1 is a unique IDv1 key for Release nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#release.
type ReleaseKeyV1 struct {
	ID uint64 // corresponds to databaseId
}

// Type implements Key.
func (r ReleaseKeyV1) Type() string {
	return TypeRelease
}

// KeyV1 implements KeyV1.
func (r ReleaseKeyV1) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// WithRepoV2 implements KeyV1NoRepo.
func (r ReleaseKeyV1) WithRepoV2(repo RepoID) KeyV2 {
	return ReleaseKeyV2{RepoID: repo, ID: r.ID}
}

// ReleaseKeyV2 is a unique IDv2 key for Release nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#release.
type ReleaseKeyV2 struct {
	RepoID RepoID // corresponds to repository.databaseId
	ID     uint64 // corresponds to databaseId
}

// Type implements Key.
func (r ReleaseKeyV2) Type() string {
	return TypeRelease
}

// GetRepoID implements KeyWithRepo.
func (r ReleaseKeyV2) GetRepoID() RepoID {
	return r.RepoID
}

// KeyV1 implements KeyV1.
func (r ReleaseKeyV2) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// KeyV2 implements KeyV2.
func (r ReleaseKeyV2) KeyV2() msgpack.RawMessage {
	return mustEncodeV2([]any{uint(0), uint(r.RepoID), uint(r.ID)})
}
