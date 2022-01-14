package internal

import (
	"sort"
	"strings"
)

func getTagList() (map[string]*Version, error) {
	output, err := execCommand("git", "tag")
	if err != nil {
		return nil, err
	}

	tmp := map[string]Versions{}
	for _, v := range strings.Split(output, "\n") {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		version, err := parseVersion(v)
		if err != nil {
			return nil, err
		}

		tmp[version.Module] = append(tmp[version.Module], version)
	}

	tmp2 := map[string]*Version{}
	for moduleName, versions := range tmp {
		sort.Sort(versions)

		if moduleName == "" {
			moduleName = "main-package"
		}
		tmp2[moduleName] = versions[len(versions)-1]
	}

	return tmp2, err
}
