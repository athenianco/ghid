package ghid

import (
	"fmt"
	"strconv"

	"github.com/vmihailenco/msgpack/v5"
)

func init() {
	RegisterDecodeV1(TypePullRequest, func(key []byte) (KeyV1, error) {
		id, err := strconv.ParseUint(string(key), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to decode pr id: %w", err)
		}
		return PRKeyV1{ID: id}, nil
	})
	RegisterDecodeV2(TypePullRequest, func(key msgpack.RawMessage) (KeyV2, error) {
		arr, err := decodeV2UintArr(TypePullRequest, 0, 2, key)
		if err != nil {
			return nil, err
		}
		return PRKeyV2{RepoID: RepoID(arr[0]), ID: arr[1]}, nil
	})
}

var (
	_ KeyV1 = PRKeyV1{}
	_ KeyV2 = PRKeyV2{}
)

// PRKeyV1 is a unique IDv1 key for PullRequest nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#pullrequest.
type PRKeyV1 struct {
	ID uint64 // corresponds to databaseId
}

// Type implements Key.
func (r PRKeyV1) Type() string {
	return TypePullRequest
}

// KeyV1 implements KeyV1.
func (r PRKeyV1) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// WithRepoV2 implements KeyV1NoRepo.
func (r PRKeyV1) WithRepoV2(repo RepoID) KeyV2 {
	return PRKeyV2{RepoID: repo, ID: r.ID}
}

// PRKeyV2 is a unique IDv2 key for PullRequest nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#pullrequest.
type PRKeyV2 struct {
	RepoID RepoID // corresponds to repository.databaseId
	ID     uint64 // corresponds to databaseId
}

// Type implements Key.
func (r PRKeyV2) Type() string {
	return TypePullRequest
}

// KeyV1 implements KeyV1.
func (r PRKeyV2) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// KeyV2 implements KeyV2.
func (r PRKeyV2) KeyV2() msgpack.RawMessage {
	return mustEncodeV2([]any{uint(0), uint(r.RepoID), uint(r.ID)})
}
