package ghid

import (
	"fmt"
	"strconv"

	"github.com/vmihailenco/msgpack/v5"
)

func init() {
	RegisterDecodeV1(TypeUser, func(key []byte) (KeyV1, error) {
		id, err := strconv.ParseUint(string(key), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to decode user id: %w", err)
		}
		return UserKey{ID: id}, nil
	})
	RegisterDecodeV2(TypeUser, func(key msgpack.RawMessage) (KeyV2, error) {
		id, err := decodeV2Uint(TypeUser, 0, key)
		if err != nil {
			return nil, err
		}
		return UserKey{ID: id}, nil
	})
}

var (
	_ KeyV1 = UserKey{}
	_ KeyV2 = UserKey{}
)

// UserKey is a unique key for User nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#user.
type UserKey struct {
	ID uint64 // corresponds to databaseId
}

// Type implements Key.
func (r UserKey) Type() string {
	return TypeUser
}

// KeyV1 implements KeyV1.
func (r UserKey) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// KeyV2 implements KeyV2.
func (r UserKey) KeyV2() msgpack.RawMessage {
	return mustEncodeV2([]any{uint(0), uint(r.ID)})
}
