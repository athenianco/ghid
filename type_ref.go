package ghid

import (
	"fmt"
	"strconv"

	"github.com/vmihailenco/msgpack/v5"
)

func init() {
	RegisterDecodeV1(TypeRef, func(key []byte) (KeyV1, error) {
		repo, rest, err := decodeV1IDAndRest(TypeRef, key)
		if err != nil {
			return nil, err
		}
		return RefKey{RepoID: RepoID(repo), RefName: string(rest)}, nil
	})
	RegisterDecodeV2(TypeRef, func(key msgpack.RawMessage) (KeyV2, error) {
		var arr []any
		err := msgpack.Unmarshal(key, &arr)
		if err != nil {
			return nil, err
		}
		if len(arr) != 3 {
			return nil, fmt.Errorf("unsupported IDv2 ref key: %#v", arr)
		}
		if v, ok := asUint64(arr[0]); !ok || v != 0 {
			return nil, fmt.Errorf("unsupported IDv2 ref key: %#v", arr)
		}
		repo, ok := asUint64(arr[1])
		if !ok {
			return nil, fmt.Errorf("unsupported IDv2 ref key: %#v", arr)
		}
		ref, ok := arr[2].(string)
		if !ok {
			return nil, fmt.Errorf("unsupported IDv2 ref key: %#v", arr)
		}
		return RefKey{RepoID: RepoID(repo), RefName: ref}, nil
	})
}

var (
	_ KeyV1 = RefKey{}
	_ KeyV2 = RefKey{}
)

// RefKey is a unique key for Ref nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#ref.
type RefKey struct {
	RepoID  RepoID // corresponds to repository.databaseId
	RefName string // corresponds to prefix+name
}

// Type implements Key.
func (r RefKey) Type() string {
	return TypeRef
}

// KeyV1 implements KeyV1.
func (r RefKey) KeyV1() string {
	return strconv.FormatUint(uint64(r.RepoID), 10) + ":" + r.RefName
}

// KeyV2 implements KeyV2.
func (r RefKey) KeyV2() msgpack.RawMessage {
	return mustEncodeV2([]any{uint(0), uint(r.RepoID), r.RefName})
}
