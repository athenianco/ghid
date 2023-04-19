package ghid

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var casesV1 = []struct {
	id        string
	typ       string
	raw       string
	key       KeyV1
	v2        string
	v2org     OrgID
	v2repo    RepoID
	parseOnly bool
}{
	{
		id:  "MDEyOk9yZ2FuaXphdGlvbjQzMTQwOTI=", // golang
		typ: TypeOrganization,
		raw: "4314092",
		key: OrgKey{4314092},
		v2:  "O_kgDOAEHT7A",
	},
	{
		id:    "MDQ6VGVhbTM1NTQxNTc=",
		typ:   TypeTeam,
		raw:   "3554157",
		key:   TeamKeyV1{3554157},
		v2:    "T_kwDOA3oJoM4ANjtt",
		v2org: 58329504,
	},
	{
		id:  "MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==", // golang/go
		typ: TypeRepository,
		raw: "23096959",
		key: RepoKey{23096959},
		v2:  "R_kgDOAWBufw",
	},
	{
		id:     "MDU6SXNzdWU1MTI3NTg2OA==", // golang/go#1
		typ:    TypeIssue,
		raw:    "51275868",
		key:    IssueKeyV1{51275868},
		v2:     "I_kwDOAWBuf84DDmhc",
		v2repo: 23096959,
	},
	{
		id:     "MDExOlB1bGxSZXF1ZXN0MjU2NTg2NDA=", // golang/go#9219
		typ:    TypePullRequest,
		raw:    "25658640",
		key:    PRKeyV1{25658640},
		v2:     "PR_kwDOAWBuf84Bh4UQ",
		v2repo: 23096959,
	},
	{
		id:     "MDU6TGFiZWwxNTA4ODAyMDk=", // golang/go, Documentation
		typ:    TypeLabel,
		raw:    "150880209",
		key:    LabelKeyV1{150880209},
		v2:     "LA_kwDOAWBuf84I_j_R",
		v2repo: 23096959,
	},
	{
		id:  "MDU6TGFiZWw=",
		typ: TypeLabel,
		raw: "",
		key: LabelKeyV1{0},
	},
	{
		id:     "MDEyOklzc3VlQ29tbWVudDY2MDQ2Mjkz", // golang/go#1, first comment
		typ:    TypeIssueComment,
		raw:    "66046293",
		key:    IssueCommentKeyV1{66046293},
		v2:     "IC_kwDOAWBuf84D78lV",
		v2repo: 23096959,
	},
	{
		id:     "MDE3OlB1bGxSZXF1ZXN0UmV2aWV3NjMzNzc5", // kubernetes/kubernetes#24536, first review
		typ:    TypePullRequestReview,
		raw:    "633779",
		key:    PRReviewKeyV1{633779},
		v2:     "PRR_kwDOAToIks4ACauz",
		v2repo: 20580498,
	},
	{
		id:     "MDIzOlB1bGxSZXF1ZXN0UmV2aWV3VGhyZWFkNDAxOTEyMjQ6djI=", // kubernetes/kubernetes#24536, first thread
		typ:    TypePullRequestReviewThread,
		raw:    "40191224:v2",
		key:    PRReviewThreadKeyV1{40191224, "v2"},
		v2:     "PRRT_kwDOAToIks4CZUT4",
		v2repo: 20580498,
	},
	{
		id:     "MDI0OlB1bGxSZXF1ZXN0UmV2aWV3Q29tbWVudDYwNDgwODMy", // kubernetes/kubernetes#24536, first thread, first comment
		typ:    TypePullRequestReviewComment,
		raw:    "60480832",
		key:    PRReviewCommentKeyV1{60480832},
		v2:     "PRRC_kwDOAToIks4Dmt1A",
		v2repo: 20580498,
	},
	{
		id:  "MDM6UmVmMjMwOTY5NTk6cmVmcy9oZWFkcy9tYXN0ZXI=", // golang/go@master
		typ: TypeRef,
		raw: "23096959:refs/heads/master",
		key: RefKey{23096959, "refs/heads/master"},
		v2:  "REF_kwDOAWBuf7FyZWZzL2hlYWRzL21hc3Rlcg",
	},
	{
		id:  "MDY6Q29tbWl0MjMwOTY5NTk6ZjIxYmUyZmRjNmYxYmVjZGJlZDE1OTJlYTBiMjQ1Y2RlZWRjNWFjOA==", // golang/go@f21be2fdc6f1becdbed1592ea0b245cdeedc5ac8
		typ: TypeCommit,
		raw: "23096959:f21be2fdc6f1becdbed1592ea0b245cdeedc5ac8",
		key: CommitKey{23096959, "f21be2fdc6f1becdbed1592ea0b245cdeedc5ac8"},
		v2:  "C_kwDOAWBuf9oAKGYyMWJlMmZkYzZmMWJlY2RiZWQxNTkyZWEwYjI0NWNkZWVkYzVhYzg",
	},
	{
		id:     "MDE3OlB1bGxSZXF1ZXN0Q29tbWl0MjU2NTg2NDA6Zjk4MWJmOGUyMGMxMzk0NmM2MjdkYzg0MWNlZGJlMWE2OTQ5MTdkNg==",
		typ:    TypePullRequestCommit,
		raw:    "25658640:f981bf8e20c13946c627dc841cedbe1a694917d6",
		key:    PRCommitKeyV1{25658640, "f981bf8e20c13946c627dc841cedbe1a694917d6"},
		v2:     "PURC_lADOAWBuf84Bh4UQ2gAoZjk4MWJmOGUyMGMxMzk0NmM2MjdkYzg0MWNlZGJlMWE2OTQ5MTdkNg",
		v2repo: 23096959,
	},
	{
		id:  "MDM6VGFnMjA1ODA0OTg6MmQ5ZjM5NzI4NzYzMWVkMGFhMmVlM2JkNzU2NWY2ZGFhYTRhNmE3Ng==", // kubernetes@v0.4
		typ: TypeTag,
		raw: "20580498:2d9f397287631ed0aa2ee3bd7565f6daaa4a6a76",
		key: TagKey{20580498, "2d9f397287631ed0aa2ee3bd7565f6daaa4a6a76"},
		v2:  "TA_kwDOAToIktoAKDJkOWYzOTcyODc2MzFlZDBhYTJlZTNiZDc1NjVmNmRhYWE0YTZhNzY",
	},
	{
		id:     "MDc6UmVsZWFzZTYyNzAyNQ==", // Kubernetes v0.4
		typ:    TypeRelease,
		raw:    "627025",
		key:    ReleaseKeyV1{627025},
		v2:     "RE_kwDOAToIks4ACZFR",
		v2repo: 20580498,
	},
	{
		id:     "MDc6UHJvamVjdDM1OTM4MDY=", // golang, Proposals (old)
		typ:    TypeProject,
		raw:    "3593806",
		key:    ProjectKeyV1{3593806},
		v2:     "PRO_kwLOAWBuf84ANtZO",
		v2repo: 23096959,
	},
	{
		id:  "MDQ6VXNlcjg1NjY5MTE=", // gopherbot
		typ: TypeUser,
		raw: "8566911",
		key: UserKey{8566911},
		v2:  "U_kgDOAIK4fw",
	},
	{
		id:  "MDM6Qm90NDk2OTkzMzM=", // dependabot
		typ: TypeBot,
		raw: "49699333",
		key: BotKey{49699333},
		v2:  "BOT_kgDOAvZaBQ",
	},
	{
		id:  "MDk6TWFubmVxdWluNTc5MDY4NDE=",
		typ: TypeMannequin,
		raw: "57906841",
		key: MannequinKey{57906841},
		v2:  "M_kgDOA3OWmQ",
	},
	{
		id:  "MDg6TGFuZ3VhZ2UxOTA=",
		typ: TypeLanguage,
		raw: "190",
		key: LangKey{190},
		v2:  "LAN_kgDMvg",
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

			if c.v2 != "" {
				v2, err := Upgrade(c.id, &UpgradeOpts{
					OrgID:  c.v2org,
					RepoID: c.v2repo,
				})
				require.NoError(t, err)
				require.Equal(t, c.v2, v2)

				// Ensure we expose org/repo ID interfaces.
				key2, err := DecodeV2(v2)
				require.NoError(t, err)
				if c.v2org != 0 {
					_ = key2.(KeyWithOrg)
				}
				if c.v2repo != 0 {
					_ = key2.(KeyWithRepo)
				}
			}
		})
	}
}
