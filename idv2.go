package ghid

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/vmihailenco/msgpack/v5"
)

// List of all known IDv2 prefixes. Without those mapping the library cannot infer type of IDv2.
var (
	idv2PrefToType = map[string]string{
		"A": TypeApp,
		"C": TypeCommit,
		"I": TypeIssue,
		"L": TypeLicense,
		"O": TypeOrganization,
		"R": TypeRepository,
		"T": TypeTeam,
		"U": TypeUser,

		"AE": TypeAssignedEvent,
		"CC": TypeCommitComment,
		"CE": TypeClosedEvent,
		"CR": TypeCheckRun,
		"CS": TypeCheckSuite,
		"DE": TypeDeployment,
		"IC": TypeIssueComment,
		"LA": TypeLabel,
		"LE": TypeLabeledEvent,
		"ME": TypeMergedEvent,
		"PE": TypePinnedEvent,
		"PR": TypePullRequest,
		"RE": TypeRelease,
		"RR": TypeReviewRequest,
		"SC": TypeStatusContext,
		"SE": TypeSubscribedEvent,
		"TA": TypeTag,
		"UE": TypeUnsubscribedEvent,

		"BOT": TypeBot,
		"CDE": TypeCommentDeletedEvent,
		"COE": TypeConnectedEvent,
		"CRE": TypeCrossReferencedEvent,
		"DEE": TypeDeployedEvent,
		"DES": TypeDeploymentStatus,
		"DIE": TypeDisconnectedEvent,
		"LAN": TypeLanguage,
		"LOE": TypeLockedEvent,
		"MEE": TypeMentionedEvent,
		"MIE": TypeMilestonedEvent,
		"PRO": TypeProject,
		"PRR": TypePullRequestReview,
		"PSH": TypePush,
		"RDE": TypeReviewDismissedEvent,
		"REE": TypeReopenedEvent,
		"REF": TypeRef,
		"RRE": TypeReviewRequestedEvent,
		"RTE": TypeRenamedTitleEvent,
		"SCR": TypeStatusCheckRollup,
		"STA": TypeStatus,
		"TRE": TypeTransferredEvent,
		"UNE": TypeUnassignedEvent,

		"AMDE": TypeAutoMergeDisabledEvent,
		"AMEE": TypeAutoMergeEnabledEvent,
		"AREE": TypeAutoRebaseEnabledEvent,
		"ASEE": TypeAutoSquashEnabledEvent,
		"ATPE": TypeAddedToProjectEvent,
		"BRCE": TypeBaseRefChangedEvent,
		"BRDE": TypeBaseRefDeletedEvent,
		"CTDE": TypeConvertToDraftEvent,
		"DEME": TypeDemilestonedEvent,
		"HRDE": TypeHeadRefDeletedEvent,
		"HRRE": TypeHeadRefRestoredEvent,
		"MADE": TypeMarkedAsDuplicateEvent,
		"PRRC": TypePullRequestReviewComment,
		"PRRT": TypePullRequestReviewThread,
		"PURC": TypePullRequestCommit,
		"REFE": TypeReferencedEvent,
		"RFPE": TypeRemovedFromProjectEvent,
		"RFRE": TypeReadyForReviewEvent,
		"RRRE": TypeReviewRequestRemovedEvent,
		"TREE": TypeTree,
		"UADE": TypeUnmarkedAsDuplicateEvent,
		"UNLE": TypeUnlabeledEvent,
		"UNPE": TypeUnpinnedEvent,

		"ABCFE": TypeAutomaticBaseChangeFailedEvent,
		"ABCSE": TypeAutomaticBaseChangeSucceededEvent,
		"BRFPE": TypeBaseRefForcePushedEvent,
		"CITDE": TypeConvertedToDiscussionEvent,
		"CNTIE": TypeConvertedNoteToIssueEvent,
		"HRFPE": TypeHeadRefForcePushedEvent,
		"MCIPE": TypeMovedColumnsInProjectEvent,
		"PRCCT": TypePullRequestCommitCommentThread,
		"UNLOE": TypeUnlockedEvent,
	}
	idv2TypeToPref = make(map[string]string)
)

func init() {
	for pref, typ := range idv2PrefToType {
		idv2TypeToPref[typ] = pref
	}
}

// PrefixesV2 returns all known IDv2 prefixes and their mapping to GraphQL types.
func PrefixesV2() map[string]string {
	m := make(map[string]string, len(idv2PrefToType))
	for k, v := range idv2PrefToType {
		m[k] = v
	}
	return m
}

// NewKeyV2 creates a custom KeyV2.
// Type and key can be arbitrary and are not verified by this function.
func NewKeyV2(typ string, key msgpack.RawMessage) KeyV2 {
	return rawKeyV2{typ: typ, key: string(key)}
}

// EncodeV2 encodes a IDv2-compatible node Key to a text format, as used for GitHub Node IDv2.
func EncodeV2(key KeyV2) (string, error) {
	pref, ok := idv2TypeToPref[key.Type()]
	if !ok {
		return "", fmt.Errorf("unsupported IDv2 type: %q", key.Type())
	}
	return EncodeV2Raw(pref, key.KeyV2()), nil
}

func mustEncodeV2(key any) msgpack.RawMessage {
	data, err := msgpack.Marshal(key)
	if err != nil {
		panic(err)
	}
	return data
}

// EncodeV2Obj encodes a node type and key payload object in IDv2 format, as used for GitHub Node IDv2.
//
// This method should only be used for testing. Use EncodeV2 with a specific key struct type instead.
func EncodeV2Obj(typePref string, key any) (string, error) {
	if bkey, ok := key.(msgpack.RawMessage); ok {
		return EncodeV2Raw(typePref, bkey), nil
	}
	data, err := msgpack.Marshal(key)
	if err != nil {
		return "", err
	}
	return EncodeV2Raw(typePref, data), nil
}

// EncodeV2Raw encodes a node type and key payload in IDv2 format, as used for GitHub Node IDv2.
//
// This method should only be used for testing. Use EncodeV2 with a specific key struct type instead.
func EncodeV2Raw(typePref string, key msgpack.RawMessage) string {
	enc := base64.RawURLEncoding
	buf := make([]byte, len(typePref)+1+enc.EncodedLen(len(key)))
	n := copy(buf, typePref)
	buf[n] = '_'
	enc.Encode(buf[n+1:], key)
	return string(buf)
}

// DecodeV2Func is a decoder function for IDv2 payload.
type DecodeV2Func func(key msgpack.RawMessage) (KeyV2, error)

var idv2Dec = make(map[string]DecodeV2Func)

// RegisterDecodeV2 registers an IDv2 payload decoder for a given type.
func RegisterDecodeV2(typ string, fnc DecodeV2Func) {
	if _, ok := idv2Dec[typ]; ok {
		panic("already registered")
	}
	idv2Dec[typ] = fnc
}

func asUint64(v any) (uint64, bool) {
	switch v := v.(type) {
	case int8:
		return uint64(v), true
	case uint8:
		return uint64(v), true
	case int16:
		return uint64(v), true
	case uint16:
		return uint64(v), true
	case int32:
		return uint64(v), true
	case uint32:
		return uint64(v), true
	case int64:
		return uint64(v), true
	case uint64:
		return v, true
	case int:
		return uint64(v), true
	case uint:
		return uint64(v), true
	}
	return 0, false
}

func isIDv2(id string) bool {
	return strings.Contains(id, "_")
}

func typeV2(id string) (string, error) {
	typ, _, err := decodeV2(id, false)
	if err != nil {
		return "", err
	}
	return typ, nil
}

// DecodeV2 decodes GitHub Node IDv2 and returns a comparable KeyV2.
//
// Returned Key may be IDv1-compatible.
func DecodeV2(id string) (KeyV2, error) {
	typ, data, err := decodeV2(id, false)
	if err != nil {
		return nil, err
	}
	if fnc, ok := idv2Dec[typ]; ok {
		return fnc(data)
	}
	return rawKeyV2{typ: typ, key: string(data)}, nil
}

// DecodeV2Raw decodes GitHub Node IDv2 and returns raw prefix and msgpack payload.
//
// It does not convert prefixes to types with PrefixesV2, as DecodeV2 does, allowing to recode event unsupported IDs.
func DecodeV2Raw(id string) (string, msgpack.RawMessage, error) {
	return decodeV2(id, true)
}

func decodeV2(id string, raw bool) (string, []byte, error) {
	i := strings.IndexByte(id, '_')
	if i < 0 {
		return "", nil, errors.New("no IDv2 type")
	}
	pref := id[:i]
	skey := id[i+1:]
	typ := pref
	if !raw {
		var ok bool
		typ, ok = idv2PrefToType[pref]
		if !ok {
			return "", nil, fmt.Errorf("unsupported IDv2 type prefix: %q", pref)
		}
	}
	data, err := base64.RawURLEncoding.DecodeString(skey)
	if err != nil {
		return typ, nil, fmt.Errorf("failed to decode IDv2 key: %w", err)
	}
	return typ, data, nil
}

func decodeV2UintArr(typ string, vers byte, sz int, data msgpack.RawMessage) ([]uint64, error) {
	var arr []uint64
	err := msgpack.Unmarshal(data, &arr)
	if err != nil {
		return nil, err
	}
	if len(arr) != sz+1 || arr[0] != uint64(vers) {
		return nil, fmt.Errorf("unsupported IDv2 %s key: %#v", typ, arr)
	}
	return arr[1:], nil
}

func decodeV2Uint(typ string, vers byte, data msgpack.RawMessage) (uint64, error) {
	arr, err := decodeV2UintArr(typ, vers, 1, data)
	if err != nil {
		return 0, err
	}
	return arr[0], nil
}

var (
	_ KeyV2 = rawKeyV2{}
)

// KeyV2 is an interface for IDv2-compatible unique node keys.
type KeyV2 interface {
	Key
	// KeyV2 returns IDv2 binary msgpack payload.
	KeyV2() msgpack.RawMessage
}

type rawKeyV2 struct {
	typ string
	key string // binary, msgpack.RawMessage
}

// Type implements Key.
func (id rawKeyV2) Type() string {
	return id.typ
}

// KeyV2 implements KeyV2.
func (id rawKeyV2) KeyV2() msgpack.RawMessage {
	return msgpack.RawMessage(id.key)
}
