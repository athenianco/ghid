package ghid

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var casesV1 = []struct {
	id  string
	typ string
	raw string
	key KeyV1
}{
	{
		id:  "MDEyOk9yZ2FuaXphdGlvbjQzMTQwOTI=", // golang
		typ: TypeOrganization,
		raw: "4314092",
		key: OrgKey{4314092},
	},
	{
		id:  "MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==", // golang/go
		typ: TypeRepository,
		raw: "23096959",
		key: RepoKey{23096959},
	},
	{
		id:  "MDU6SXNzdWU1MTI3NTg2OA==", // golang/go#1
		typ: TypeIssue,
		raw: "51275868",
		key: IssueKeyV1{51275868},
	},
	{
		id:  "MDExOlB1bGxSZXF1ZXN0MjU2NTg2NDA=", // golang/go#9219
		typ: TypePullRequest,
		raw: "25658640",
		key: PRKeyV1{25658640},
	},
	{
		id:  "MDM6UmVmMjMwOTY5NTk6cmVmcy9oZWFkcy9tYXN0ZXI=", // golang/go@master
		typ: TypeRef,
		raw: "23096959:refs/heads/master",
		key: RefKey{23096959, "refs/heads/master"},
	},
	{
		id:  "MDY6Q29tbWl0MjMwOTY5NTk6ZjIxYmUyZmRjNmYxYmVjZGJlZDE1OTJlYTBiMjQ1Y2RlZWRjNWFjOA==", // golang/go@f21be2fdc6f1becdbed1592ea0b245cdeedc5ac8
		typ: TypeCommit,
		raw: "23096959:f21be2fdc6f1becdbed1592ea0b245cdeedc5ac8",
		key: CommitKey{23096959, "f21be2fdc6f1becdbed1592ea0b245cdeedc5ac8"},
	},
	{
		id:  "MDc6UmVsZWFzZTYyNzAyNQ==", // Kubernetes v0.4
		typ: TypeRelease,
		raw: "627025",
		key: ReleaseKeyV1{627025},
	},
	{
		id:  "MDc6UHJvamVjdDM1OTM4MDY=", // golang, Proposals (old)
		typ: TypeProject,
		raw: "3593806",
		key: ProjectKeyV1{3593806},
	},
	{
		id:  "MDQ6VXNlcjg1NjY5MTE=", // gopherbot
		typ: TypeUser,
		raw: "8566911",
		key: UserKey{8566911},
	},
	{
		id:  "MDM6Qm90NDk2OTkzMzM=", // dependabot
		typ: TypeBot,
		raw: "49699333",
		key: BotKey{49699333},
	},
	{
		id:  "MDg6TGFuZ3VhZ2UxOTA=",
		typ: TypeLanguage,
		raw: "190",
		key: LangKey{190},
	},
	{
		id:  "MDc6TGljZW5zZTU=",
		typ: TypeLicense,
		raw: "5",
		key: LicenseKeyV1{5},
	},
}

func TestIDv1(t *testing.T) {
	for _, c := range casesV1 {
		c := c
		t.Run(c.typ, func(t *testing.T) {
			typ, rkey, err := decodeV1(c.id)
			require.NoError(t, err)
			require.Equal(t, c.typ, typ)
			require.Equal(t, c.raw, string(rkey))

			got := EncodeV1(rawKeyV1{typ: typ, key: string(rkey)})
			require.Equal(t, c.id, got)

			key, err := DecodeV1(c.id)
			require.NoError(t, err)
			require.Equal(t, c.key, key)

			got = EncodeV1(key)
			require.Equal(t, c.id, got)
		})
	}
}
