package day7

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type dir struct {
	totalSize int
}
type sys struct {
	cwd string
}

func updateSize(folder string, m map[string]dir, size int) {
	paths := strings.Split(folder, "/")
	count := len(paths)
	for count >= 1 {
		currPath := strings.Join(paths[:count], "/")
		if currPath == "" {
			currPath = "/"
		}
		dir := m[currPath]
		dir.totalSize += size
		m[currPath] = dir
		count--
	}
}
func Solution() {
	steps := strings.Split(utils.GetInputData(7), "\n")

	dirs := map[string]dir{}
	sys := sys{}
	for _, step := range steps {
		if strings.Contains(step, "$ ls") {
			continue
		}
		if strings.Contains(step, "dir ") {
			continue
		}
		if strings.Contains(step, "$ cd") {
			dirName := strings.Split(step, " ")[2]
			if dirName == "/" {
				continue
			}
			if dirName == ".." {
				paths := strings.Split(sys.cwd, "/")
				sys.cwd = strings.Join(paths[:len(paths)-1], "/")
			} else {
				sys.cwd += fmt.Sprintf("/%s", dirName)
			}
			continue
		}

		stdout := strings.Split(step, " ")
		fileSize := stdout[0]
		fileSizeInt, err := strconv.Atoi(fileSize)
		if err != nil {
			panic("Could not convert file size string to int")
		}
		updateSize(sys.cwd, dirs, fileSizeInt)

	}
	total := 0

	// Find the dirs with specified size
	unUsedSpace := 70000000 - dirs["/"].totalSize
	spaceNeeded := 30000000 - unUsedSpace
	dirToDelete := 40000000

	for k := range dirs {
		d := dirs[k]
		if d.totalSize <= 100000 {
			total += d.totalSize
		}

		if d.totalSize > spaceNeeded {
			if d.totalSize < dirToDelete {
				dirToDelete = d.totalSize
			}
		}
	}
	fmt.Println(dirToDelete)
	fmt.Println(total)
}
