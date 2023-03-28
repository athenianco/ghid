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
	RegisterDecodeV1(TypeBot, func(key []byte) (KeyV1, error) {
		id, err := strconv.ParseUint(string(key), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to decode bot id: %w", err)
		}
		return BotKey{ID: id}, nil
	})
	RegisterDecodeV2(TypeBot, func(key msgpack.RawMessage) (KeyV2, error) {
		id, err := decodeV2Uint(TypeBot, 0, key)
		if err != nil {
			return nil, err
		}
		return BotKey{ID: id}, nil
	})
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

var (
	_ KeyV1 = BotKey{}
	_ KeyV2 = BotKey{}
)

// BotKey is a unique key for Bot nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#bot.
type BotKey struct {
	ID uint64 // corresponds to databaseId
}

// Type implements Key.
func (r BotKey) Type() string {
	return TypeBot
}

// KeyV1 implements KeyV1.
func (r BotKey) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// KeyV2 implements KeyV2.
func (r BotKey) KeyV2() msgpack.RawMessage {
	return mustEncodeV2([]any{uint(0), uint(r.ID)})
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
