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
	RegisterDecodeV1(TypePullRequestCommit, func(key []byte) (KeyV1, error) {
		pr, rest, err := decodeV1IDAndRest(TypePullRequestCommit, key)
		if err != nil {
			return nil, err
		}
		return PRCommitKeyV1{PR: pr, SHA: string(rest)}, nil
	})
	RegisterDecodeV2(TypePullRequestCommit, func(key msgpack.RawMessage) (KeyV2, error) {
		var arr []any
		err := msgpack.Unmarshal(key, &arr)
		if err != nil {
			return nil, err
		}
		if len(arr) != 4 {
			return nil, fmt.Errorf("unsupported IDv2 PR commit key: %#v", arr)
		}
		if v, ok := asUint64(arr[0]); !ok || v != 0 {
			return nil, fmt.Errorf("unsupported IDv2 PR commit key: %#v", arr)
		}
		repo, ok := asUint64(arr[1])
		if !ok {
			return nil, fmt.Errorf("unsupported IDv2 PR commit key: %#v", arr)
		}
		pr, ok := asUint64(arr[2])
		if !ok {
			return nil, fmt.Errorf("unsupported IDv2 PR commit key: %#v", arr)
		}
		sha, ok := arr[3].(string)
		if !ok {
			return nil, fmt.Errorf("unsupported IDv2 PR commit key: %#v", arr)
		}
		return PRCommitKeyV2{RepoID: RepoID(repo), PR: pr, SHA: sha}, nil
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

// GetRepoID implements KeyWithRepo.
func (r PRKeyV2) GetRepoID() RepoID {
	return r.RepoID
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

// GetRepoID implements KeyWithRepo.
func (r PRReviewKeyV2) GetRepoID() RepoID {
	return r.RepoID
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

// GetRepoID implements KeyWithRepo.
func (r PRReviewThreadKeyV2) GetRepoID() RepoID {
	return r.RepoID
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

// GetRepoID implements KeyWithRepo.
func (r PRReviewCommentKeyV2) GetRepoID() RepoID {
	return r.RepoID
}

// KeyV1 implements KeyV1.
func (r PRReviewCommentKeyV2) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// KeyV2 implements KeyV2.
func (r PRReviewCommentKeyV2) KeyV2() msgpack.RawMessage {
	return mustEncodeV2([]any{uint(0), uint(r.RepoID), uint(r.ID)})
}

var (
	_ KeyV1NoRepo = PRCommitKeyV1{}
	_ KeyV2       = PRCommitKeyV2{}
)

// PRCommitKeyV1 is a unique IDv1 key for PullRequestCommit nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#pullrequestcommit.
type PRCommitKeyV1 struct {
	PR  uint64 // corresponds to pullRequest.databaseId
	SHA string
}

// Type implements Key.
func (r PRCommitKeyV1) Type() string {
	return TypePullRequestCommit
}

// KeyV1 implements KeyV1.
func (r PRCommitKeyV1) KeyV1() string {
	return strconv.FormatUint(r.PR, 10) + ":" + r.SHA
}

// WithRepoV2 implements KeyV1NoRepo.
func (r PRCommitKeyV1) WithRepoV2(repo RepoID) KeyV2 {
	return PRCommitKeyV2{RepoID: repo, PR: r.PR, SHA: r.SHA}
}

// PRCommitKeyV2 is a unique IDv2 key for PullRequestCommit nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#pullrequestcommit.
type PRCommitKeyV2 struct {
	RepoID RepoID // corresponds to repository.databaseId
	PR     uint64 // corresponds to pullRequest.databaseId
	SHA    string
}

// Type implements Key.
func (r PRCommitKeyV2) Type() string {
	return TypePullRequestCommit
}

// GetRepoID implements KeyWithRepo.
func (r PRCommitKeyV2) GetRepoID() RepoID {
	return r.RepoID
}

// KeyV1 implements KeyV1.
func (r PRCommitKeyV2) KeyV1() string {
	return strconv.FormatUint(r.PR, 10) + ":" + r.SHA
}

// KeyV2 implements KeyV2.
func (r PRCommitKeyV2) KeyV2() msgpack.RawMessage {
	// GitHub encodes commit SHA as string16, not string8, even though SHA length is only 40 (<256). Optimization?
	return mustEncodeV2([]any{uint(0), uint(r.RepoID), uint(r.PR), msgpackEncodeStr16(r.SHA)})
}
