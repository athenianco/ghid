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
		return OrgKey{ID: OrgID(id)}, nil
	})
	RegisterDecodeV2(TypeOrganization, func(key msgpack.RawMessage) (KeyV2, error) {
		id, err := decodeV2Uint(TypeOrganization, 0, key)
		if err != nil {
			return nil, err
		}
		return OrgKey{ID: OrgID(id)}, nil
	})
	RegisterDecodeV1(TypeTeam, func(key []byte) (KeyV1, error) {
		id, err := strconv.ParseUint(string(key), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to decode team id: %w", err)
		}
		return TeamKeyV1{ID: id}, nil
	})
	RegisterDecodeV2(TypeTeam, func(key msgpack.RawMessage) (KeyV2, error) {
		arr, err := decodeV2UintArr(TypeTeam, 0, 2, key)
		if err != nil {
			return nil, err
		}
		return TeamKeyV2{OrgID: OrgID(arr[0]), ID: arr[1]}, nil
	})
}

// OrgID corresponds to databaseId field of the Organization node.
//
// See https://docs.github.com/en/graphql/reference/objects#organization.
type OrgID uint64

var (
	_ KeyV1 = OrgKey{}
	_ KeyV2 = OrgKey{}
)

// OrgKey is a unique key for Organization nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#organization.
type OrgKey struct {
	ID OrgID // corresponds to databaseId
}

// Type implements Key.
func (r OrgKey) Type() string {
	return TypeOrganization
}

// GetOrgID implements KeyWithOrg.
func (r OrgKey) GetOrgID() OrgID {
	return r.ID
}

// KeyV1 implements KeyV1.
func (r OrgKey) KeyV1() string {
	return strconv.FormatUint(uint64(r.ID), 10)
}

// KeyV2 implements KeyV2.
func (r OrgKey) KeyV2() msgpack.RawMessage {
	return mustEncodeV2([]any{uint(0), uint(r.ID)})
}

var (
	_ KeyV1NoOrg = TeamKeyV1{}
	_ KeyV2      = TeamKeyV2{}
)

// TeamKeyV1 is a unique IDv1 key for Team nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#team.
type TeamKeyV1 struct {
	ID uint64 // corresponds to databaseId
}

// Type implements Key.
func (r TeamKeyV1) Type() string {
	return TypeTeam
}

// KeyV1 implements KeyV1.
func (r TeamKeyV1) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// WithOrgV2 implements KeyV1NoOrg.
func (r TeamKeyV1) WithOrgV2(org OrgID) KeyV2 {
	return TeamKeyV2{OrgID: org, ID: r.ID}
}

// TeamKeyV2 is a unique IDv2 key for Team nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#team.
type TeamKeyV2 struct {
	OrgID OrgID  // corresponds to organization.databaseId
	ID    uint64 // corresponds to databaseId
}

// Type implements Key.
func (r TeamKeyV2) Type() string {
	return TypeTeam
}

// GetOrgID implements KeyWithOrg.
func (r TeamKeyV2) GetOrgID() OrgID {
	return r.OrgID
}

// KeyV1 implements KeyV1.
func (r TeamKeyV2) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// KeyV2 implements KeyV2.
func (r TeamKeyV2) KeyV2() msgpack.RawMessage {
	return mustEncodeV2([]any{uint(0), uint(r.OrgID), uint(r.ID)})
}
