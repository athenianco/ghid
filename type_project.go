package ghid

import (
	"fmt"
	"strconv"

	"github.com/vmihailenco/msgpack/v5"
)

func init() {
	RegisterDecodeV1(TypeProject, func(key []byte) (KeyV1, error) {
		id, err := strconv.ParseUint(string(key), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to decode project id: %w", err)
		}
		return ProjectKeyV1{ID: id}, nil
	})
	RegisterDecodeV2(TypeProject, func(key msgpack.RawMessage) (KeyV2, error) {
		arr, err := decodeV2UintArr(TypeIssue, 2, 2, key)
		if err != nil {
			return nil, err
		}
		return ProjectKeyV2{RepoID: RepoID(arr[0]), ID: arr[1]}, nil
	})
}

var (
	_ KeyV1 = ProjectKeyV1{}
	_ KeyV2 = ProjectKeyV2{}
)

// ProjectKeyV1 is a unique IDv1 key for Project nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#project.
type ProjectKeyV1 struct {
	ID uint64 // corresponds to databaseId
}

// Type implements Key.
func (r ProjectKeyV1) Type() string {
	return TypeProject
}

// KeyV1 implements KeyV1.
func (r ProjectKeyV1) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// WithRepoV2 implements KeyV1NoRepo.
func (r ProjectKeyV1) WithRepoV2(repo RepoID) KeyV2 {
	return ProjectKeyV2{RepoID: repo, ID: r.ID}
}

// ProjectKeyV2 is a unique IDv2 key for Project nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#project.
type ProjectKeyV2 struct {
	RepoID RepoID // corresponds to repository.databaseId
	ID     uint64 // corresponds to databaseId
}

// Type implements Key.
func (r ProjectKeyV2) Type() string {
	return TypeProject
}

// KeyV1 implements KeyV1.
func (r ProjectKeyV2) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// KeyV2 implements KeyV2.
func (r ProjectKeyV2) KeyV2() msgpack.RawMessage {
	return mustEncodeV2([]any{uint(2), uint(r.RepoID), uint(r.ID)})
}
