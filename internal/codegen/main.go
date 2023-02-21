package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var (
	fTypes  = flag.String("types", "./github_schema.json.gz", "types cache")
	fSchema = flag.Bool("schema", false, "only fetch the schema")
	fOut    = flag.String("out", "types_gen.go", "output path")
	fToken  = flag.String("token", "", "GitHub access token to fetch the schema")
)

func main() {
	flag.Parse()
	if err := run(*fOut, *fTypes); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(out, tpath string) error {
	if _, err := os.Stat(tpath); os.IsNotExist(err) || *fSchema {
		fmt.Println("fetching the schema")
		if err := os.MkdirAll(filepath.Dir(tpath), 0755); err != nil {
			return err
		}
		if err := fetchSchemaTo(*fToken, tpath); err != nil {
			return err
		}
	}
	if *fSchema {
		return nil
	}
	return ProcessSchema(out, tpath)
}

type ObjectKind string

const (
	KindList      = ObjectKind("LIST")
	KindNonNull   = ObjectKind("NON_NULL")
	KindScalar    = ObjectKind("SCALAR")
	KindEnum      = ObjectKind("ENUM")
	KindObject    = ObjectKind("OBJECT")
	KindInterface = ObjectKind("INTERFACE")
	KindUnion     = ObjectKind("UNION")
)

type SchemaType struct {
	Kind          ObjectKind    `json:"kind"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	Fields        []SchemaField `json:"fields"`
	Interfaces    []TypeRef     `json:"interfaces"`
	PossibleTypes []TypeRef     `json:"possibleTypes"`
	EnumValues    []EnumValue   `json:"enumValues"`
	InputFields   []Arg         `json:"inputFields"`
}

type SchemaField struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Args        []Arg    `json:"args"`
	Type        *TypeRef `json:"type"`
	Deprecated  bool     `json:"isDeprecated"`
}

type Arg struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Type        *TypeRef `json:"type"`
	// TODO: defaultValue
}

type TypeRef struct {
	Kind   ObjectKind `json:"kind"`
	Name   string     `json:"name"`
	OfType *TypeRef   `json:"ofType"`
}

type EnumValue struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Deprecated  bool   `json:"isDeprecated"`
}

func ProcessSchema(out, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	zf, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer zf.Close()
	var r io.Reader = zf

	var resp struct {
		Data struct {
			Schema struct {
				Types []*SchemaType `json:"types"`
			} `json:"__schema"`
		} `json:"data"`
	}
	err = json.NewDecoder(r).Decode(&resp)
	if err != nil {
		return err
	}
	types := resp.Data.Schema.Types
	if len(types) == 0 {
		return errors.New("no types found")
	}
	var names []string
	for _, typ := range types {
		if typ.Name == "Node" && typ.Kind == KindInterface {
			for _, r := range typ.PossibleTypes {
				names = append(names, r.Name)
			}
		}
	}
	if len(names) == 0 {
		return errors.New("no Node types found")
	}
	sort.Strings(names)
	var buf bytes.Buffer
	buf.WriteString("package ghid\n\nconst (\n")
	for _, name := range names {
		fmt.Fprintf(&buf, `	// Type%s is constant for a type of %s node.
	//
	// See https://docs.github.com/en/graphql/reference/objects#%s.
	Type%s = %q

`,
			name, name, strings.ToLower(name),
			name, name,
		)
	}
	buf.WriteString(")\n")
	err = os.WriteFile(out, buf.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

func fetchSchema(githubToken string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/graphql", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "bearer "+githubToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("non-200 OK status code: %v body: %q", resp.Status, body)
	}
	return resp.Body, nil
}

func fetchSchemaTo(token, path string) error {
	if token == "" {
		if etok := os.Getenv("GITHUB_TOKEN"); etok != "" {
			token = etok
		} else {
			return errors.New("github token must be specified")
		}
	}
	rc, err := fetchSchema(token)
	if err != nil {
		return err
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(nil)
	err = json.Indent(buf, data, "", "\t")
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	zw := gzip.NewWriter(f)
	if _, err := buf.WriteTo(zw); err != nil {
		return err
	}
	if err = zw.Close(); err != nil {
		return err
	}
	return f.Close()
}
