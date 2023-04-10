package ghid

import (
	"fmt"
	"strconv"

	"github.com/vmihailenco/msgpack/v5"
)

func init() {
	RegisterDecodeV1(TypeLabel, func(key []byte) (KeyV1, error) {
		id, err := strconv.ParseUint(string(key), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to decode label id: %w", err)
		}
		return LabelKeyV1{ID: id}, nil
	})
	RegisterDecodeV2(TypeLabel, func(key msgpack.RawMessage) (KeyV2, error) {
		arr, err := decodeV2UintArr(TypeLabel, 0, 2, key)
		if err != nil {
			return nil, err
		}
		return LabelKeyV2{RepoID: RepoID(arr[0]), ID: arr[1]}, nil
	})
}

var (
	_ KeyV1NoRepo = LabelKeyV1{}
	_ KeyV2       = LabelKeyV2{}
)

// LabelKeyV1 is a unique IDv1 key for Label nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#label.
type LabelKeyV1 struct {
	ID uint64
}

// Type implements Key.
func (r LabelKeyV1) Type() string {
	return TypeLabel
}

// KeyV1 implements KeyV1.
func (r LabelKeyV1) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// WithRepoV2 implements KeyV1NoRepo.
func (r LabelKeyV1) WithRepoV2(repo RepoID) KeyV2 {
	return LabelKeyV2{RepoID: repo, ID: r.ID}
}

// LabelKeyV2 is a unique IDv2 key for Label nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#label.
type LabelKeyV2 struct {
	RepoID RepoID // corresponds to repository.databaseId
	ID     uint64
}

// Type implements Key.
func (r LabelKeyV2) Type() string {
	return TypeLabel
}

// KeyV1 implements KeyV1.
func (r LabelKeyV2) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// KeyV2 implements KeyV2.
func (r LabelKeyV2) KeyV2() msgpack.RawMessage {
	return mustEncodeV2([]any{uint(0), uint(r.RepoID), uint(r.ID)})
}
