package ghid

import (
	"fmt"
	"strconv"

	"github.com/vmihailenco/msgpack/v5"
)

func init() {
	RegisterDecodeV1(TypeLanguage, func(key []byte) (KeyV1, error) {
		id, err := strconv.ParseUint(string(key), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to decode lang id: %w", err)
		}
		return LangKey{ID: id}, nil
	})
	RegisterDecodeV2(TypeLanguage, func(key msgpack.RawMessage) (KeyV2, error) {
		id, err := decodeV2Uint(TypeLanguage, 0, key)
		if err != nil {
			return nil, err
		}
		return LangKey{ID: id}, nil
	})
}

var (
	_ KeyV1 = LangKey{}
	_ KeyV2 = LangKey{}
)

// LangKey is a unique key for Language nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#language.
type LangKey struct {
	ID uint64 // corresponds to databaseId
}

// Type implements Key.
func (r LangKey) Type() string {
	return TypeLanguage
}

// KeyV1 implements KeyV1.
func (r LangKey) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// KeyV2 implements KeyV2.
func (r LangKey) KeyV2() msgpack.RawMessage {
	return mustEncodeV2([]any{uint(0), uint(r.ID)})
}
