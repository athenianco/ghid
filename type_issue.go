package ghid

import (
	"fmt"
	"strconv"

	"github.com/vmihailenco/msgpack/v5"
)

func init() {
	RegisterDecodeV1(TypeIssue, func(key []byte) (KeyV1, error) {
		id, err := strconv.ParseUint(string(key), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to decode issue id: %w", err)
		}
		return IssueKeyV1{ID: id}, nil
	})
	RegisterDecodeV1(TypeIssueComment, func(key []byte) (KeyV1, error) {
		id, err := strconv.ParseUint(string(key), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to decode issue comment id: %w", err)
		}
		return IssueCommentKeyV1{ID: id}, nil
	})
	RegisterDecodeV2(TypeIssue, func(key msgpack.RawMessage) (KeyV2, error) {
		arr, err := decodeV2UintArr(TypeIssue, 0, 2, key)
		if err != nil {
			return nil, err
		}
		return IssueKeyV2{RepoID: RepoID(arr[0]), ID: arr[1]}, nil
	})
	RegisterDecodeV2(TypeIssueComment, func(key msgpack.RawMessage) (KeyV2, error) {
		arr, err := decodeV2UintArr(TypeIssueComment, 0, 2, key)
		if err != nil {
			return nil, err
		}
		return IssueCommentKeyV2{RepoID: RepoID(arr[0]), ID: arr[1]}, nil
	})
}

var (
	_ KeyV1NoRepo = IssueKeyV1{}
	_ KeyV2       = IssueKeyV2{}
)

// IssueKeyV1 is a unique IDv1 key for Issue nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#issue.
type IssueKeyV1 struct {
	ID uint64 // corresponds to databaseId
}

// Type implements Key.
func (r IssueKeyV1) Type() string {
	return TypeIssue
}

// KeyV1 implements KeyV1.
func (r IssueKeyV1) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// WithRepoV2 implements KeyV1NoRepo.
func (r IssueKeyV1) WithRepoV2(repo RepoID) KeyV2 {
	return IssueKeyV2{RepoID: repo, ID: r.ID}
}

// IssueKeyV2 is a unique IDv2 key for Issue nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#issue.
type IssueKeyV2 struct {
	RepoID RepoID // corresponds to repository.databaseId
	ID     uint64 // corresponds to databaseId
}

// Type implements Key.
func (r IssueKeyV2) Type() string {
	return TypeIssue
}

// KeyV1 implements KeyV1.
func (r IssueKeyV2) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// KeyV2 implements KeyV2.
func (r IssueKeyV2) KeyV2() msgpack.RawMessage {
	return mustEncodeV2([]any{uint(0), uint(r.RepoID), uint(r.ID)})
}

var (
	_ KeyV1NoRepo = IssueCommentKeyV1{}
	_ KeyV2       = IssueCommentKeyV2{}
)

// IssueCommentKeyV1 is a unique IDv1 key for IssueComment nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#issuecomment.
type IssueCommentKeyV1 struct {
	ID uint64 // corresponds to databaseId
}

// Type implements Key.
func (r IssueCommentKeyV1) Type() string {
	return TypeIssueComment
}

// KeyV1 implements KeyV1.
func (r IssueCommentKeyV1) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// WithRepoV2 implements KeyV1NoRepo.
func (r IssueCommentKeyV1) WithRepoV2(repo RepoID) KeyV2 {
	return IssueCommentKeyV2{RepoID: repo, ID: r.ID}
}

// IssueCommentKeyV2 is a unique IDv2 key for IssueComment nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#issuecomment.
type IssueCommentKeyV2 struct {
	RepoID RepoID // corresponds to repository.databaseId
	ID     uint64 // corresponds to databaseId
}

// Type implements Key.
func (r IssueCommentKeyV2) Type() string {
	return TypeIssueComment
}

// KeyV1 implements KeyV1.
func (r IssueCommentKeyV2) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// KeyV2 implements KeyV2.
func (r IssueCommentKeyV2) KeyV2() msgpack.RawMessage {
	return mustEncodeV2([]any{uint(0), uint(r.RepoID), uint(r.ID)})
}
