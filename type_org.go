package ghid

import (
	"fmt"
	"strconv"

	"github.com/vmihailenco/msgpack/v5"
)

func init() {
	RegisterDecodeV1(TypeOrganization, func(key []byte) (KeyV1, error) {
		id, err := strconv.ParseUint(string(key), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to decode org id: %w", err)
		}
		return OrgKey{ID: id}, nil
	})
	RegisterDecodeV2(TypeOrganization, func(key msgpack.RawMessage) (KeyV2, error) {
		id, err := decodeV2Uint(TypeOrganization, 0, key)
		if err != nil {
			return nil, err
		}
		return OrgKey{ID: id}, nil
	})
}

var (
	_ KeyV1 = OrgKey{}
	_ KeyV2 = OrgKey{}
)

// OrgKey is a unique key for Organization nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#organization.
type OrgKey struct {
	ID uint64 // corresponds to databaseId
}

// Type implements Key.
func (r OrgKey) Type() string {
	return TypeOrganization
}

// KeyV1 implements KeyV1.
func (r OrgKey) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// KeyV2 implements KeyV2.
func (r OrgKey) KeyV2() msgpack.RawMessage {
	return mustEncodeV2([]any{uint(0), uint(r.ID)})
}
