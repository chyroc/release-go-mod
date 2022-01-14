package internal

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Version struct {
	Module string
	Major  int
	Minor  int
	Patch  int
}

type Versions []*Version

func (r Versions) Len() int {
	return len(r)
}

func (r Versions) Less(i, j int) bool {
	if r[i].Major != r[j].Major {
		return r[i].Major < r[j].Major
	}
	if r[i].Minor != r[j].Minor {
		return r[i].Minor < r[j].Minor
	}
	return r[i].Patch < r[j].Patch
}

func (r Versions) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

var regVer = regexp.MustCompile(`(?m)^v?(\d+)\.(\d+)\.(\d+)$`)

func parseVersion(s string) (*Version, error) {
	res := new(Version)

	s = strings.TrimSpace(s)
	if strings.Count(s, "/") > 0 {
		res.Module = strings.TrimSpace(strings.Split(s, "/")[0])
		s = strings.TrimSpace(strings.Split(s, "/")[1])
	}

	match := regVer.FindStringSubmatch(s)
	if len(match) != 4 || match[1] == "" || match[2] == "" || match[3] == "" {
		return nil, fmt.Errorf("invalid: %q", s)
	}
	res.Major, _ = strconv.Atoi(match[1])
	res.Minor, _ = strconv.Atoi(match[2])
	res.Patch, _ = strconv.Atoi(match[3])
	return res, nil
}
