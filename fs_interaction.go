package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func FindIps() map[string]string {
	if file, err := os.ReadFile("ips.txt"); err == nil {
		fmt.Println("Content", string(file))
		ipMap := make(map[string]string)
		for _, line := range strings.Split(string(file), "\n") {
			k, v := filterIpLine(strings.Split(line, " "))
			ipMap[k] = v
		}
		return ipMap
	}
	return nil
}

func FindIpList() []string {
	if file, err := os.ReadFile("ips.txt"); err == nil {
		ipMap := []string{}

		for _, line := range strings.Split(string(file), "\n") {
			k, _ := filterIpLine(strings.Split(line, " "))
			ipMap = append(ipMap, k)
		}

		return ipMap
	}
	return nil
}
func FindKeys() map[string]string {
	hDir, _ := os.UserHomeDir()
	dir, _ := os.ReadDir(hDir + "/.ssh/")

	keys := make(map[string]string)
	for _, d := range dir {
		fInfo, _ := fs.DirEntry.Info(d)
		if !fInfo.IsDir() && !strings.HasSuffix(fInfo.Name(), ".pub") {
			keys[fInfo.Name()] = hDir + "/.ssh/" + fInfo.Name()
		}
	}

	return keys
}

func FindKeyList() []string {
	hDir, _ := os.UserHomeDir()
	dir, _ := os.ReadDir(hDir + "/.ssh/")

	keys := []string{}
	for _, d := range dir {
		fInfo, _ := fs.DirEntry.Info(d)
		if !fInfo.IsDir() && !strings.HasSuffix(fInfo.Name(), ".pub") {
			keys = append(keys, fInfo.Name())
		}
	}

	return keys
}

func filterIpLine(l []string) (string, string) {
	res := []string{}
	for _, c := range l {
		if c != "" {
			res = append(res, c)
		}
	}
	return res[1], res[0]
}
