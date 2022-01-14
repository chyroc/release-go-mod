package internal

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Command(c *cli.Context) error {
	tagList, err := getTagList()
	if err != nil {
		return err
	}

	names := []string{}
	for k := range tagList {
		names = append(names, k)
	}

	idx, err := UISelect(fmt.Sprintf("请选择要升级哪个 sub-module"), names)
	if err != nil {
		return fmt.Errorf("fail to select sub-module: %v\n", err)
	}
	version := tagList[names[idx]]

	releaseMessage, err := UIInput("请输入发版详情")
	if err != nil {
		return fmt.Errorf("fail to input release message: %v\n", err)
	}

	currentVersion := fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch)
	currentVersionMajor := fmt.Sprintf("%d.%d.%d", version.Major+1, 0, 0)
	currentVersionMinor := fmt.Sprintf("%d.%d.%d", version.Major, version.Minor+1, 0)
	currentVersionPatch := fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch+1)
	items := []string{currentVersionPatch, currentVersionMinor, currentVersionMajor}
	idx, err = UISelect(fmt.Sprintf("当前版本是: %v, 请选择升级到哪个版本", currentVersion), items)
	if err != nil {
		return fmt.Errorf("Failed to select: %v\n", err)
	}
	newVersion := items[idx]
	fmt.Printf("从 %v 升级到 %v\n", currentVersion, newVersion)

	writeFile, err := tryUpdateVersionFile(version.Module, currentVersion, newVersion)
	if err != nil {
		return err
	}

	tag := fmt.Sprintf("%v/v%v", version.Module, newVersion)
	if version.Module == "" {
		tag = fmt.Sprintf("v%v", newVersion)
	}
	if writeFile {
		if err := execCommandStd("git", "commit", "-a", "-m", fmt.Sprintf("release %s", tag)); err != nil {
			return err
		}
	}
	if err := execCommandStd("git", "tag", tag, "-m", releaseMessage); err != nil {
		return err
	}

	fmt.Println("success")
	return nil
}
