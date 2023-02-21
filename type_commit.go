package ghid

import (
	"fmt"
	"strconv"

	"github.com/vmihailenco/msgpack/v5"
	"github.com/vmihailenco/msgpack/v5/msgpcode"
)

func init() {
	RegisterDecodeV1(TypeCommit, func(key []byte) (KeyV1, error) {
		repo, rest, err := decodeV1InRepo(TypeCommit, key)
		if err != nil {
			return nil, err
		}
		return CommitKey{RepoID: repo, SHA: string(rest)}, nil
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

// KeyV1 implements KeyV1.
func (r CommitKey) KeyV1() string {
	return strconv.FormatUint(uint64(r.RepoID), 10) + ":" + r.SHA
}

// KeyV2 implements KeyV2.
func (r CommitKey) KeyV2() msgpack.RawMessage {
	// GitHub encodes commit SHA as string16, not string8, even though SHA length is only 40 (<256). Optimization?
	n := len(r.SHA)
	buf := make([]byte, 1+2+n)
	buf[0] = msgpcode.Str16
	buf[1] = byte(n >> 8)
	buf[2] = byte(n)
	copy(buf[3:], r.SHA)
	return mustEncodeV2([]any{uint(0), uint(r.RepoID), msgpack.RawMessage(buf)})
}
