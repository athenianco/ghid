// Package ghid provides decoding and encoding utilities for GitHub GraphQL Node IDs.
//
// Ideally, ID in any kind of system must be unique. However, as the system scales, you may need to change primary keys
// in your DB, and thus change the unique identifier used in your API as well. The same happened with GitHub GraphQL API.
//
// To achieve a notion of Node IDs (a unique identifier, regardless of the object/node type), you likely don't want to
// use generic IDs like UUID - it's more efficient to encode a type of the node and some kind of primary key in the ID.
// This is exactly how GitHub Node IDs work: they encode some type-related information, so that the backend can quickly
// figure out which table/collection to use for a lookup, and an identifier specific for that node type.
//
// First version of GitHub Node IDs (aka IDv1) were text-based (base64-encoded text) which turned out to be inefficient.
// GitHub migrated to IDv2 format, which is binary-based (base64 over msgpack encoding). However, the migration was not
// handled well in our opinion.
//
// This library solves a problem: if you already have a huge dataset of cached info based on GitHub's GraphQL data,
// and you want to migrate it to the new ID format in offline mode (without hitting the actual API).
// This project provides a CLI and a Go library to handle these cases. As a bonus, it allows introspection of ID contents.
// When your client receives an ID, you can instantly tell which type it contains, what is the commit SHA,
// which repo it belongs to, etc. It can be useful as an optimization for some use cases.
//
// WARNING: This library is written with some assumptions inferred from reverse-engineering a bunch of GitHub Node IDs,
// which we've seen in the wild. It is NOT affiliated with GitHub. Generated output MAY be incorrect and break your system.
// Always design a fallback code path that actually hits GitHub API and verify output for your specific use case.
package ghid

import (
	"fmt"
)

// Key is an interface for underlying ID key.
// Implementations are comparable and can be used as a map key.
type Key interface {
	// Type returns a GitHub node type for this key.
	//
	// Type names directly correspond to Node implementation in GraphQL API:
	// https://docs.github.com/en/graphql/reference/interfaces#node.
	Type() string
}

// Type returns a GitHub node type for a given ID.
//
// Type names directly correspond to Node implementation in GraphQL API:
// https://docs.github.com/en/graphql/reference/interfaces#node.
func Type(id string) (string, error) {
	if isIDv2(id) {
		return typeV2(id)
	}
	return typeV1(id)
}

// Decode GitHub Node ID and return a unique node key. If ID version or type is unsupported, it returns an error.
//
// Some node keys may need to be upgraded before they can be compared. See Upgrade and UpgradeOpts.
func Decode(id string) (Key, error) {
	if isIDv2(id) {
		k, err := DecodeV2(id)
		if err != nil {
			return nil, err
		}
		return k, nil
	}
	k, err := DecodeV1(id)
	if err != nil {
		return nil, err
	}
	return k, nil
}

// UpgradeOpts provide
type UpgradeOpts struct {
	// RepoID will be used for ID upgrade, if it wasn't previously encoded in IDv1, but is required for IDv2.
	RepoID RepoID
}

// Upgrade tries to upgrade GitHub Node ID to the latest version.
// Upgrade always returns a non-empty valid ID, even in case of a failure.
// An error is returned if the upgrade is possible, but not currently supported.
//
// Some IDv2 contain additional information that is not available in corresponding IDv1.
// UpgradeOpts are used in this case and provide an additional context.
//
// For example, PRKeyV1 does not contain a repository ID, which is required for PRKeyV2.
// Thus, to upgrade this key type, it's required to set RepoID in UpgradeOpts.
func Upgrade(id string, opts *UpgradeOpts) (string, error) {
	if isIDv2(id) {
		return id, nil
	}
	key1, err := DecodeV1(id)
	if err != nil {
		return id, fmt.Errorf("cannot decode IDv1: %w", err)
	}
	key2, err := upgradeKey(key1, opts)
	if err != nil {
		return id, err
	}
	id2, err := EncodeV2(key2)
	if err != nil {
		return id, fmt.Errorf("cannot encode IDv2: %w", err)
	}
	return id2, nil
}

// UpgradeKey tries to upgrade node Key to the latest version.
// UpgradeKey always returns a non-nil Key, even in case of a failure.
// An error is returned if the upgrade is possible, but not currently supported.
//
// See Upgrade for more details.
func UpgradeKey(key Key, opts *UpgradeOpts) (Key, error) {
	key2, err := upgradeKey(key, opts)
	if err != nil {
		return key, err
	}
	return key2, nil
}

func upgradeKey(key Key, opts *UpgradeOpts) (KeyV2, error) {
	if key2, ok := key.(KeyV2); ok {
		return key2, nil
	}
	if ukey, ok := key.(KeyV1NoRepo); ok && opts != nil && opts.RepoID != 0 {
		return ukey.WithRepoV2(opts.RepoID), nil
	}
	return nil, fmt.Errorf("unsupported IDv1 -> IDv2 conversion for type %q", key.Type())
}
