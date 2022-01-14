package internal

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func tryUpdateVersionFile(module string, current string, target string) (bool, error) {
	root := "."
	moduleGoMod := "go.mod"
	if module != "" {
		root = "./" + module
		moduleGoMod = module + "/go.mod"
	}
	_ = moduleGoMod
	expectGoModulePath := getGoModulePathOfDir(root)

	versionFile := ""
	_ = filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if versionFile != "" {
			return nil
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		if bs, _ := ioutil.ReadFile(path); strings.Contains(string(bs), current) {
			pathGoModule := getGoModulePath(path)
			if pathGoModule == expectGoModulePath {
				versionFile = path
			}
		}
		return nil
	})

	if versionFile == "" {
		return false, nil
	}

	content, err := ioutil.ReadFile(versionFile)
	if err != nil {
		return false, err
	}
	return true, ioutil.WriteFile(versionFile, []byte(strings.Replace(string(content), current, target, -1)), 0o644)
}

func getGoModulePath(path string) string {
	dir, _ := filepath.Split(path)
	return getGoModulePathOfDir(dir)
}

func getGoModulePathOfDir(dir string) string {
	var err error
	dir, err = filepath.Abs(dir)
	if err != nil {
		return ""
	}
	if dir == "/" {
		return ""
	}
	file := dir + "/go.mod"
	if _, err := ioutil.ReadFile(file); err == nil {
		return file
	}

	return getGoModulePathOfDir(filepath.Join(dir, "../"))
}
