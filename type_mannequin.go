package ghid

import (
	"fmt"
	"strconv"

	"github.com/vmihailenco/msgpack/v5"
)

func init() {
	RegisterDecodeV1(TypeMannequin, func(key []byte) (KeyV1, error) {
		id, err := strconv.ParseUint(string(key), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to decode mannequin id: %w", err)
		}
		return MannequinKey{ID: id}, nil
	})
	RegisterDecodeV2(TypeMannequin, func(key msgpack.RawMessage) (KeyV2, error) {
		id, err := decodeV2Uint(TypeMannequin, 0, key)
		if err != nil {
			return nil, err
		}
		return MannequinKey{ID: id}, nil
	})
}

var (
	_ KeyV1 = MannequinKey{}
	_ KeyV2 = MannequinKey{}
)

// MannequinKey is a unique key Mannequin Bot nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#mannequin.
type MannequinKey struct {
	ID uint64 // corresponds to databaseId
}

// Type implements Key.
func (r MannequinKey) Type() string {
	return TypeMannequin
}

// KeyV1 implements KeyV1.
func (r MannequinKey) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// KeyV2 implements KeyV2.
func (r MannequinKey) KeyV2() msgpack.RawMessage {
	return mustEncodeV2([]any{uint(0), uint(r.ID)})
}
