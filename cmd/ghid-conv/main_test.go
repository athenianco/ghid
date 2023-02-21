package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var cases = []struct {
	name   string
	input  string
	output string
	json   bool
}{
	{
		name:   "json ok",
		input:  `["MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==", "MDY6Q29tbWl0MjMwOTY5NTk6ZjIxYmUyZmRjNmYxYmVjZGJlZDE1OTJlYTBiMjQ1Y2RlZWRjNWFjOA=="]`,
		output: `[{"old":"MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==","new":"R_kgDOAWBufw"},{"old":"MDY6Q29tbWl0MjMwOTY5NTk6ZjIxYmUyZmRjNmYxYmVjZGJlZDE1OTJlYTBiMjQ1Y2RlZWRjNWFjOA==","new":"C_kwDOAWBuf9oAKGYyMWJlMmZkYzZmMWJlY2RiZWQxNTkyZWEwYjI0NWNkZWVkYzVhYzg"}]` + "\n",
		json:   true,
	},
	{
		name:   "json err",
		input:  `["MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==", "MDU6SXNzdWU1MTI3NTg2OA==", "MDQ6VXNlcjg1NjY5MTE="]`,
		output: `[{"old":"MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==","new":"R_kgDOAWBufw"},{"old":"MDU6SXNzdWU1MTI3NTg2OA==","error":"unsupported IDv1 -\u003e IDv2 conversion for type \"Issue\""},{"old":"MDQ6VXNlcjg1NjY5MTE=","new":"U_kgDOAIK4fw"}]` + "\n",
		json:   true,
	},
	{
		name: "text ok",
		input: `
MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==
MDY6Q29tbWl0MjMwOTY5NTk6ZjIxYmUyZmRjNmYxYmVjZGJlZDE1OTJlYTBiMjQ1Y2RlZWRjNWFjOA==
`,
		output: `R_kgDOAWBufw
C_kwDOAWBuf9oAKGYyMWJlMmZkYzZmMWJlY2RiZWQxNTkyZWEwYjI0NWNkZWVkYzVhYzg
`,
	},
	{
		name: "text err",
		input: `
MDEwOlJlcG9zaXRvcnkyMzA5Njk1OQ==
MDU6SXNzdWU1MTI3NTg2OA==
MDQ6VXNlcjg1NjY5MTE=
`,
		output: `R_kgDOAWBufw
error: unsupported IDv1 -> IDv2 conversion for type "Issue"
U_kgDOAIK4fw
`,
	},
}

func TestCLI(t *testing.T) {
	var buf bytes.Buffer
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			proc := processLines
			if c.json {
				proc = processJSON
			}
			buf.Reset()
			err := proc(&buf, strings.NewReader(c.input))
			require.NoError(t, err)
			require.Equal(t, c.output, buf.String())
		})
	}
}
