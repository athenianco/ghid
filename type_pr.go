package ghid

import (
	"bytes"
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
	RegisterDecodeV1(TypePullRequestReview, func(key []byte) (KeyV1, error) {
		id, err := strconv.ParseUint(string(key), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to decode pr review id: %w", err)
		}
		return PRReviewKeyV1{ID: id}, nil
	})
	RegisterDecodeV1(TypePullRequestReviewThread, func(key []byte) (KeyV1, error) {
		i := bytes.IndexByte(key, ':')
		if i < 0 {
			return nil, fmt.Errorf("unsupported pr review thread id: %q", string(key))
		}
		id, err := strconv.ParseUint(string(key[:i]), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to decode pr review thread id: %w", err)
		}
		return PRReviewThreadKeyV1{ID: id, Vers: string(key[i+1:])}, nil
	})
	RegisterDecodeV1(TypePullRequestReviewComment, func(key []byte) (KeyV1, error) {
		id, err := strconv.ParseUint(string(key), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to decode pr review comment id: %w", err)
		}
		return PRReviewCommentKeyV1{ID: id}, nil
	})
	RegisterDecodeV2(TypePullRequest, func(key msgpack.RawMessage) (KeyV2, error) {
		arr, err := decodeV2UintArr(TypePullRequest, 0, 2, key)
		if err != nil {
			return nil, err
		}
		return PRKeyV2{RepoID: RepoID(arr[0]), ID: arr[1]}, nil
	})
	RegisterDecodeV2(TypePullRequestReview, func(key msgpack.RawMessage) (KeyV2, error) {
		arr, err := decodeV2UintArr(TypePullRequestReview, 0, 2, key)
		if err != nil {
			return nil, err
		}
		return PRReviewKeyV2{RepoID: RepoID(arr[0]), ID: arr[1]}, nil
	})
	RegisterDecodeV2(TypePullRequestReviewThread, func(key msgpack.RawMessage) (KeyV2, error) {
		arr, err := decodeV2UintArr(TypePullRequestReviewThread, 0, 2, key)
		if err != nil {
			return nil, err
		}
		return PRReviewThreadKeyV2{RepoID: RepoID(arr[0]), ID: arr[1]}, nil
	})
	RegisterDecodeV2(TypePullRequestReviewComment, func(key msgpack.RawMessage) (KeyV2, error) {
		arr, err := decodeV2UintArr(TypePullRequestReviewComment, 0, 2, key)
		if err != nil {
			return nil, err
		}
		return PRReviewCommentKeyV2{RepoID: RepoID(arr[0]), ID: arr[1]}, nil
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

var (
	_ KeyV1NoRepo = PRReviewKeyV1{}
	_ KeyV2       = PRReviewKeyV2{}
)

// PRReviewKeyV1 is a unique IDv1 key for PullRequestReview nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#pullrequestreview.
type PRReviewKeyV1 struct {
	ID uint64 // corresponds to databaseId
}

// Type implements Key.
func (r PRReviewKeyV1) Type() string {
	return TypePullRequestReview
}

// KeyV1 implements KeyV1.
func (r PRReviewKeyV1) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// WithRepoV2 implements KeyV1NoRepo.
func (r PRReviewKeyV1) WithRepoV2(repo RepoID) KeyV2 {
	return PRReviewKeyV2{RepoID: repo, ID: r.ID}
}

// PRReviewKeyV2 is a unique IDv2 key for PullRequestReview nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#pullrequestreview.
type PRReviewKeyV2 struct {
	RepoID RepoID // corresponds to repository.databaseId
	ID     uint64 // corresponds to databaseId
}

// Type implements Key.
func (r PRReviewKeyV2) Type() string {
	return TypePullRequestReview
}

// KeyV1 implements KeyV1.
func (r PRReviewKeyV2) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// KeyV2 implements KeyV2.
func (r PRReviewKeyV2) KeyV2() msgpack.RawMessage {
	return mustEncodeV2([]any{uint(0), uint(r.RepoID), uint(r.ID)})
}

var (
	_ KeyV1NoRepo = PRReviewThreadKeyV1{}
	_ KeyV2       = PRReviewThreadKeyV2{}
)

// PRReviewThreadKeyV1 is a unique IDv1 key for PullRequestReviewThread nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#pullrequestreviewthread.
type PRReviewThreadKeyV1 struct {
	ID   uint64
	Vers string
}

// Type implements Key.
func (r PRReviewThreadKeyV1) Type() string {
	return TypePullRequestReviewThread
}

// KeyV1 implements KeyV1.
func (r PRReviewThreadKeyV1) KeyV1() string {
	return strconv.FormatUint(r.ID, 10) + ":" + r.Vers
}

// WithRepoV2 implements KeyV1NoRepo.
func (r PRReviewThreadKeyV1) WithRepoV2(repo RepoID) KeyV2 {
	return PRReviewThreadKeyV2{RepoID: repo, ID: r.ID}
}

// PRReviewThreadKeyV2 is a unique IDv2 key for PullRequestReviewThread nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#pullrequestreviewthread.
type PRReviewThreadKeyV2 struct {
	RepoID RepoID // corresponds to repository.databaseId
	ID     uint64
}

// Type implements Key.
func (r PRReviewThreadKeyV2) Type() string {
	return TypePullRequestReviewThread
}

// KeyV2 implements KeyV2.
func (r PRReviewThreadKeyV2) KeyV2() msgpack.RawMessage {
	return mustEncodeV2([]any{uint(0), uint(r.RepoID), uint(r.ID)})
}

var (
	_ KeyV1NoRepo = PRReviewCommentKeyV1{}
	_ KeyV2       = PRReviewCommentKeyV2{}
)

// PRReviewCommentKeyV1 is a unique IDv1 key for PullRequestReviewComment nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#pullrequestreviewcomment.
type PRReviewCommentKeyV1 struct {
	ID uint64 // corresponds to databaseId
}

// Type implements Key.
func (r PRReviewCommentKeyV1) Type() string {
	return TypePullRequestReviewComment
}

// KeyV1 implements KeyV1.
func (r PRReviewCommentKeyV1) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// WithRepoV2 implements KeyV1NoRepo.
func (r PRReviewCommentKeyV1) WithRepoV2(repo RepoID) KeyV2 {
	return PRReviewCommentKeyV2{RepoID: repo, ID: r.ID}
}

// PRReviewCommentKeyV2 is a unique IDv2 key for PullRequestReviewComment nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#pullrequestreviewcomment.
type PRReviewCommentKeyV2 struct {
	RepoID RepoID // corresponds to repository.databaseId
	ID     uint64 // corresponds to databaseId
}

// Type implements Key.
func (r PRReviewCommentKeyV2) Type() string {
	return TypePullRequestReviewComment
}

// KeyV1 implements KeyV1.
func (r PRReviewCommentKeyV2) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// KeyV2 implements KeyV2.
func (r PRReviewCommentKeyV2) KeyV2() msgpack.RawMessage {
	return mustEncodeV2([]any{uint(0), uint(r.RepoID), uint(r.ID)})
}
