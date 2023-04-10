package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/athenianco/ghid"
)

var (
	fIn   = flag.String("i", "", "input file")
	fOut  = flag.String("o", "", "output file")
	fJson = flag.Bool("json", false, "use json format instead of text")
	fRepo = flag.Uint64("repo", 0, "use specified repo id for upgrading ids")
	fOrg  = flag.Uint64("org", 0, "use specified org id for upgrading ids")
)

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var r io.Reader
	if path := *fIn; path == "" || path == "-" {
		r = os.Stdin
	} else {
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		r = f
	}
	var (
		w io.Writer
		c io.Closer
	)
	if path := *fOut; path == "" || path == "-" {
		r = os.Stdout
	} else {
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		w, c = f, f
	}
	var err error
	if *fJson || strings.HasSuffix(*fOut, ".json") || strings.HasSuffix(*fIn, ".json") {
		err = processJSON(w, r)
	} else {
		err = processLines(w, r)
	}
	if err != nil {
		return err
	}
	return c.Close()
}

func processJSON(w io.Writer, r io.Reader) error {
	var ids []string
	if err := json.NewDecoder(r).Decode(&ids); err != nil {
		return err
	}
	type Resp struct {
		Old   string `json:"old,omitempty"`
		New   string `json:"new,omitempty"`
		Error string `json:"error,omitempty"`
	}
	out := make([]Resp, 0, len(ids))
	for _, id := range ids {
		if id2, err := convert(id); err != nil {
			fmt.Fprintln(os.Stderr, id, err)
			out = append(out, Resp{Old: id, Error: err.Error()})
		} else {
			out = append(out, Resp{Old: id, New: id2})
		}
	}
	return json.NewEncoder(w).Encode(out)
}

func convert(id string) (string, error) {
	return ghid.Upgrade(id, &ghid.UpgradeOpts{
		OrgID:  ghid.OrgID(*fOrg),
		RepoID: ghid.RepoID(*fRepo),
	})
}

func processLines(w io.Writer, r io.Reader) error {
	bw := bufio.NewWriter(w)
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		text := bytes.TrimSpace(sc.Bytes())
		if len(text) == 0 {
			continue
		}
		id := string(text)
		id2, err := convert(id)
		if err != nil {
			fmt.Fprintln(os.Stderr, id, err)
			if _, werr := fmt.Fprintln(bw, "error:", err.Error()); werr != nil {
				return werr
			}
			continue
		}
		if _, werr := bw.WriteString(id2); werr != nil {
			return werr
		}
		if werr := bw.WriteByte('\n'); werr != nil {
			return werr
		}
	}
	if err := sc.Err(); err != nil {
		return err
	}
	return bw.Flush()
}
