package ghid

import (
	"fmt"
	"strconv"

	"github.com/vmihailenco/msgpack/v5"
)

func init() {
	RegisterDecodeV1(TypeLicense, func(key []byte) (KeyV1, error) {
		id, err := strconv.ParseUint(string(key), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to decode license id: %w", err)
		}
		return LicenseKeyV1{ID: id}, nil
	})
	RegisterDecodeV2(TypeLicense, func(key msgpack.RawMessage) (KeyV2, error) {
		var arr []any
		err := msgpack.Unmarshal(key, &arr)
		if err != nil {
			return nil, err
		}
		if len(arr) != 2 {
			return nil, fmt.Errorf("unsupported IDv2 license key: %#v", arr)
		}
		if v, ok := asUint64(arr[0]); !ok || v != 0 {
			return nil, fmt.Errorf("unsupported IDv2 license key: %#v", arr)
		}
		spdx, ok := arr[1].(string)
		if !ok {
			return nil, fmt.Errorf("unsupported IDv2 license key: %#v", arr)
		}
		return LicenseKeyV2{SpdxID: spdx}, nil
	})
}

var (
	_ KeyV1 = LicenseKeyV1{}
	_ KeyV2 = LicenseKeyV2{}
)

// LicenseKeyV1 is a unique IDv1 key for License nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#license.
type LicenseKeyV1 struct {
	ID uint64
}

// Type implements Key.
func (r LicenseKeyV1) Type() string {
	return TypeLicense
}

// KeyV1 implements KeyV1.
func (r LicenseKeyV1) KeyV1() string {
	return strconv.FormatUint(r.ID, 10)
}

// LicenseKeyV2 is a unique IDv2 key for License nodes.
//
// See https://docs.github.com/en/graphql/reference/objects#license.
type LicenseKeyV2 struct {
	SpdxID string // corresponds to ToLower(spdxId)
}

// Type implements Key.
func (r LicenseKeyV2) Type() string {
	return TypeLicense
}

// KeyV2 implements KeyV2.
func (r LicenseKeyV2) KeyV2() msgpack.RawMessage {
	return mustEncodeV2([]any{uint(0), r.SpdxID})
}
