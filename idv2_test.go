package ghid

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vmihailenco/msgpack/v5"
)

var casesV2 = []struct {
	id  string
	typ string
	raw any
	key KeyV2
}{
	{
		id:  "O_kgDOAEHT7A", // golang
		typ: TypeOrganization,
		raw: []any{int8(0), uint32(4314092)},
		key: OrgKey{4314092},
	},
	{
		id:  "R_kgDOAWBufw", // golang/go
		typ: TypeRepository,
		raw: []any{int8(0), uint32(23096959)},
		key: RepoKey{23096959},
	},
	{
		id:  "I_kwDOAWBuf84DDmhc", // golang/go#1
		typ: TypeIssue,
		raw: []any{int8(0), uint32(23096959), uint32(51275868)},
		key: IssueKeyV2{23096959, 51275868},
	},
	{
		id:  "PR_kwDOAWBuf84Bh4UQ", // golang/go#9219
		typ: TypePullRequest,
		raw: []any{int8(0), uint32(23096959), uint32(25658640)},
		key: PRKeyV2{23096959, 25658640},
	},
	{
		id:  "REF_kwDOAWBuf7FyZWZzL2hlYWRzL21hc3Rlcg", // golang/go@master
		typ: TypeRef,
		raw: []any{int8(0), uint32(23096959), "refs/heads/master"},
		key: RefKey{23096959, "refs/heads/master"},
	},
	{
		id:  "C_kwDOAWBuf9oAKGYyMWJlMmZkYzZmMWJlY2RiZWQxNTkyZWEwYjI0NWNkZWVkYzVhYzg", // golang/go@f21be2fdc6f1becdbed1592ea0b245cdeedc5ac8
		typ: TypeCommit,
		raw: []any{int8(0), uint32(23096959), "f21be2fdc6f1becdbed1592ea0b245cdeedc5ac8"},
		key: CommitKey{23096959, "f21be2fdc6f1becdbed1592ea0b245cdeedc5ac8"},
	},
	{
		id:  "RE_kwDOAToIks4ACZFR", // Kubernetes v0.4
		typ: TypeRelease,
		raw: []any{int8(0), uint32(20580498), uint32(627025)},
		key: ReleaseKeyV2{20580498, 627025},
	},
	{
		id:  "PRO_kwLOAWBuf84ANtZO", // golang, Proposals (old)
		typ: TypeProject,
		raw: []any{int8(2), uint32(23096959), uint32(3593806)},
		key: ProjectKeyV2{23096959, 3593806},
	},
	{
		id:  "U_kgDOAIK4fw", // gopherbot
		typ: TypeUser,
		raw: []any{int8(0), uint32(8566911)},
		key: UserKey{8566911},
	},
	{
		id:  "BOT_kgDOAvZaBQ", // dependabot
		typ: TypeBot,
		raw: []any{int8(0), uint32(49699333)},
		key: BotKey{49699333},
	},
	{
		id:  "LAN_kgDMvg",
		typ: TypeLanguage,
		raw: []any{int8(0), uint8(190)},
		key: LangKey{190},
	},
	{
		id:  "L_kgCsYnNkLTMtY2xhdXNl",
		typ: TypeLicense,
		raw: []any{int8(0), "bsd-3-clause"},
		key: LicenseKeyV2{"bsd-3-clause"},
	},
}

func unmarshalV2(t testing.TB, data []byte) []any {
	var rdata []any
	err := msgpack.Unmarshal(data, &rdata)
	require.NoError(t, err)
	return rdata
}

func TestSortPrefixes(t *testing.T) {
	var prefs [][2]string
	for pref, typ := range idv2PrefToType {
		prefs = append(prefs, [2]string{
			pref, typ,
		})
	}
	sort.Slice(prefs, func(i, j int) bool {
		p1, p2 := prefs[i][0], prefs[j][0]
		if len(p1) < len(p2) {
			return true
		} else if len(p1) > len(p2) {
			return false
		}
		return p1 < p2
	})
	var buf bytes.Buffer
	prev := 1
	for _, v := range prefs {
		if len(v[0]) != prev {
			buf.WriteByte('\n')
			prev = len(v[0])
		}
		fmt.Fprintf(&buf, "\t%q: Type%s,\n", v[0], v[1])
	}
	t.Logf("\n%s", buf.String())
}

func TestIDv2(t *testing.T) {
	for _, c := range casesV2 {
		c := c
		t.Run(c.typ, func(t *testing.T) {
			typ, rkey, err := decodeV2(c.id, false)
			require.NoError(t, err)
			require.Equal(t, c.typ, typ)
			rdata := unmarshalV2(t, rkey)
			require.Equal(t, c.raw, rdata)

			got, err := EncodeV2(rawKeyV2{typ: typ, key: string(rkey)})
			require.NoError(t, err)
			require.Equal(t, c.id, got)

			key, err := DecodeV2(c.id)
			require.NoError(t, err)
			require.Equal(t, c.key, key)

			got, err = EncodeV2(key)
			require.NoError(t, err)
			if c.id != got {
				_, rkey1, err := decodeV2(c.id, true)
				require.NoError(t, err)
				rdata1 := unmarshalV2(t, rkey1)

				_, rkey2, err := decodeV2(got, true)
				require.NoError(t, err)
				rdata2 := unmarshalV2(t, rkey2)

				require.Equal(t, rdata1, rdata2)
				require.Equal(t, rkey1, rkey2,
					"exp: %s\ngot: %s\nuse https://sugendran.github.io/msgpack-visualizer/ to debug",
					base64.StdEncoding.EncodeToString(rkey1), hex.EncodeToString(rkey1),
					base64.StdEncoding.EncodeToString(rkey2), hex.EncodeToString(rkey2),
				)
				require.Equal(t, c.id, got)
			}
		})
	}
}
