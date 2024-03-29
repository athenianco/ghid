package ghid

import (
	"fmt"
	"strconv"

	"github.com/vmihailenco/msgpack/v5"
	"github.com/vmihailenco/msgpack/v5/msgpcode"
)

func init() {
	RegisterDecodeV1(TypeCommit, func(key []byte) (KeyV1, error) {
		repo, rest, err := decodeV1IDAndRest(TypeCommit, key)
		if err != nil {
			return nil, err
		}
		return CommitKey{RepoID: RepoID(repo), SHA: string(rest)}, nil
	})
	RegisterDecodeV2(TypeCommit, func(key msgpack.RawMessage) (KeyV2, error) {
		var arr []any
		err := msgpack.Unmarshal(key, &arr)
		if err != nil {
			return nil, err
		}
		if len(arr) != 3 {
			return nil, fmt.Errorf("unsupported IDv2 commit key: %#v", arr)
		}
		if v, ok := asUint64(arr[0]); !ok || v != 0 {
			return nil, fmt.Errorf("unsupported IDv2 commit key: %#v", arr)
		}
		repo, ok := asUint64(arr[1])
		if !ok {
			return nil, fmt.Errorf("unsupported IDv2 commit key: %#v", arr)
		}
		sha, ok := arr[2].(string)
		if !ok {
			return nil, fmt.Errorf("unsupported IDv2 commit key: %#v", arr)
		}
		return CommitKey{RepoID: RepoID(repo), SHA: sha}, nil
	})
	RegisterDecodeV1(TypeTag, func(key []byte) (KeyV1, error) {
		repo, rest, err := decodeV1IDAndRest(TypeTag, key)
		if err != nil {
			return nil, err
		}
		return TagKey{RepoID: RepoID(repo), SHA: string(rest)}, nil
	})
	RegisterDecodeV2(TypeTag, func(key msgpack.RawMessage) (KeyV2, error) {
		var arr []any
		err := msgpack.Unmarshal(key, &arr)
		if err != nil {
			return nil, err
		}
		if len(arr) != 3 {
			return nil, fmt.Errorf("unsupported IDv2 tag key: %#v", arr)
		}
		if v, ok := asUint64(arr[0]); !ok || v != 0 {
			return nil, fmt.Errorf("unsupported IDv2 tag key: %#v", arr)
		}
		repo, ok := asUint64(arr[1])
		if !ok {
			return nil, fmt.Errorf("unsupported IDv2 tag key: %#v", arr)
		}
		sha, ok := arr[2].(string)
		if !ok {
			return nil, fmt.Errorf("unsupported IDv2 tag key: %#v", arr)
		}
		return TagKey{RepoID: RepoID(repo), SHA: sha}, nil
	})
}

func msgpackEncodeStr16(s string) msgpack.RawMessage {
	n := len(s)
	buf := make([]byte, 1+2+n)
	buf[0] = msgpcode.Str16
	buf[1] = byte(n >> 8)
	buf[2] = byte(n)
	copy(buf[3:], s)
	return buf
}

var (
	_ KeyV1 = CommitKey{}
	_ KeyV2 = CommitKey{}
)

// CommitKey is a unique key for Commit nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#commit.
type CommitKey struct {
	RepoID RepoID // corresponds to repository.databaseId
	SHA    string // corresponds to oid
}

// Type implements Key.
func (r CommitKey) Type() string {
	return TypeCommit
}

// GetRepoID implements KeyWithRepo.
func (r CommitKey) GetRepoID() RepoID {
	return r.RepoID
}

// KeyV1 implements KeyV1.
func (r CommitKey) KeyV1() string {
	return strconv.FormatUint(uint64(r.RepoID), 10) + ":" + r.SHA
}

// KeyV2 implements KeyV2.
func (r CommitKey) KeyV2() msgpack.RawMessage {
	// GitHub encodes commit SHA as string16, not string8, even though SHA length is only 40 (<256). Optimization?
	return mustEncodeV2([]any{uint(0), uint(r.RepoID), msgpackEncodeStr16(r.SHA)})
}

var (
	_ KeyV1 = TagKey{}
	_ KeyV2 = TagKey{}
)

// TagKey is a unique key for Tag nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#tag.
type TagKey struct {
	RepoID RepoID // corresponds to repository.databaseId
	SHA    string // corresponds to oid
}

// Type implements Key.
func (r TagKey) Type() string {
	return TypeTag
}

// GetRepoID implements KeyWithRepo.
func (r TagKey) GetRepoID() RepoID {
	return r.RepoID
}

// KeyV1 implements KeyV1.
func (r TagKey) KeyV1() string {
	return strconv.FormatUint(uint64(r.RepoID), 10) + ":" + r.SHA
}

// KeyV2 implements KeyV2.
func (r TagKey) KeyV2() msgpack.RawMessage {
	// GitHub encodes commit SHA as string16, not string8, even though SHA length is only 40 (<256). Optimization?
	return mustEncodeV2([]any{uint(0), uint(r.RepoID), msgpackEncodeStr16(r.SHA)})
}
