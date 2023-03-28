package ghid

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
)

// NewKeyV1 creates a custom KeyV1.
// Type and key can be arbitrary and are not verified by this function.
func NewKeyV1(typ, key string) KeyV1 {
	return rawKeyV1{typ: typ, key: key}
}

// EncodeV1 encodes a IDv1-compatible node Key to a text format, as used for GitHub Node IDv1.
func EncodeV1(key KeyV1) string {
	return EncodeV1Raw(key.Type(), key.KeyV1())
}

// EncodeV1Raw encodes a node type and key in IDv1 format, as used for GitHub Node IDv1.
//
// This method should only be used for testing. Use EncodeV1 with a specific key struct type instead.
func EncodeV1Raw(typ, key string) string {
	var buf bytes.Buffer
	buf.Grow(1 + 2 + 1 + len(typ) + len(key))
	buf.WriteByte('0')
	buf.WriteString(strconv.Itoa(len(typ)))
	buf.WriteByte(':')
	buf.WriteString(typ)
	buf.WriteString(key)
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

// DecodeV1Func is a decoder function for IDv1 payload.
type DecodeV1Func func(key []byte) (KeyV1, error)

var idv1Dec = make(map[string]DecodeV1Func)

// RegisterDecodeV1 registers an IDv1 payload decoder for a given type.
func RegisterDecodeV1(typ string, fnc DecodeV1Func) {
	if _, ok := idv1Dec[typ]; ok {
		panic("already registered")
	}
	idv1Dec[typ] = fnc
}

// DecodeV1 decodes GitHub Node IDv1 and returns a comparable KeyV1.
//
// Returned Key may be IDv2-compatible, or may require additional info for an upgrade. See Upgrade.
func DecodeV1(id string) (KeyV1, error) {
	typ, key, err := decodeV1(id)
	if err != nil {
		return nil, err
	}
	if fnc, ok := idv1Dec[typ]; ok {
		return fnc(key)
	}
	return rawKeyV1{typ: typ, key: string(key)}, nil
}

func typeV1(id string) (string, error) {
	typ, _, err := decodeV1(id)
	if err != nil {
		return "", err
	}
	return typ, nil
}

func decodeV1(id string) (string, []byte, error) {
	b, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		return "", nil, err
	} else if len(b) == 0 {
		return "", nil, errors.New("empty id data")
	}
	// 0<N>:<type name, N bytes><internal PK>
	if b[0] != '0' {
		return "", nil, fmt.Errorf("unsupported internal IDv1 version: %s", string(b[0]))
	}
	b = b[1:]
	i := bytes.Index(b, []byte(":"))
	if i < 0 {
		return "", nil, fmt.Errorf("unsupported or invalid IDv1 format: %q", string(b))
	}
	n, err := strconv.Atoi(string(b[:i]))
	if err != nil {
		return "", nil, fmt.Errorf("cannot parse type length: %w", err)
	}
	b = b[i+1:]
	if n > len(b) {
		return "", nil, fmt.Errorf("invalid type length: [%d:%d]", n, len(b))
	}
	typ := string(b[:n])
	key := b[n:]
	return typ, key, nil
}

func decodeV1IDAndRest(typ string, key []byte) (uint64, []byte, error) {
	i := bytes.IndexByte(key, ':')
	if i < 0 {
		return 0, nil, fmt.Errorf("invalid IDv1 key for %s: %q", typ, string(key))
	}
	repoID, err := strconv.ParseUint(string(key[:i]), 10, 64)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to decode %s repo id: %w", typ, err)
	}
	return repoID, key[i+1:], nil
}

var (
	_ KeyV1 = rawKeyV1{}
)

// KeyV1 is an interface for IDv1-compatible unique node keys.
type KeyV1 interface {
	Key
	// KeyV1 returns IDv1 text payload.
	KeyV1() string
}

type rawKeyV1 struct {
	typ string
	key string
}

// Type implements Key.
func (id rawKeyV1) Type() string {
	return id.typ
}

// KeyV1 implements KeyV1.
func (id rawKeyV1) KeyV1() string {
	return id.key
}

// KeyV1NoRepo is an optional interface for KeyV1 that allows upgrades to KeyV2, given RepoID.
type KeyV1NoRepo interface {
	KeyV1
	// WithRepoV2 uses RepoID to upgrade KeyV1 to KeyV2.
	WithRepoV2(repo RepoID) KeyV2
}
