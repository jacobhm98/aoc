package main

import (
	"aoc-22/pkg/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type dir struct {
	parent    *dir
	childDirs []*dir
	files     []file
	totalSize int
	name      string
}

func (parent *dir) addChildIfNotPresent(childName string) {
	if containsDir(parent.childDirs, childName) {
		return
	}
	newChild := dir{
		parent: parent,
		name:   childName,
	}
	parent.childDirs = append(parent.childDirs, &newChild)
}

func (parent *dir) addFileIfNotPresent(sizeStr string, name string) {
	if containsFile(parent.files, name) {
		return
	}
	size, _ := strconv.Atoi(sizeStr)
	newFile := file{
		parentDir: parent,
		size:      size,
		name:      name,
	}
	parent.files = append(parent.files, newFile)
}

func containsFile(files []file, fileName string) bool {
	for _, file := range files {
		if file.name == fileName {
			return true
		}
	}
	return false
}

func containsDir(children []*dir, childName string) bool {
	for _, child := range children {
		if child.name == childName {
			return true
		}
	}
	return false
}

type file struct {
	parentDir *dir
	size      int
	name      string
}

var root dir

func main() {
	lines := utils.ReadFile("day7/input.txt")
	dirTree := buildDirTree(lines)
	totalSize := populateSizes(dirTree)
	currUnusedSpace := 70000000 - totalSize
	limit := 30000000 - currUnusedSpace
	dirToBeDeleted := findSmallestDirLargerThanLimit(dirTree, limit)
	fmt.Println(dirToBeDeleted.totalSize)
}

func findSmallestDirLargerThanLimit(tree *dir, limit int) *dir {
	allDirs := getDirList(tree)
	sort.Slice(allDirs, func(i, j int) bool {
		return allDirs[i].totalSize < allDirs[j].totalSize
	})
	for _, dir := range allDirs {
		if dir.totalSize >= limit {
			return dir
		}
	}
	panic("could not find a dir larger than the limit")
}

func getDirList(tree *dir) []*dir {
	var dirList []*dir
	dirList = append(dirList, tree)
	for _, child := range tree.childDirs {
		dirList = append(dirList, getDirList(child)...)
	}
	return dirList
}

func totalSizeOfAllDirsLTE(tree *dir, limit int) int {
	totalSizeOfAllDirsLTELimit := 0
	if tree.totalSize <= limit {
		totalSizeOfAllDirsLTELimit += tree.totalSize
	}
	for _, child := range tree.childDirs {
		totalSizeOfAllDirsLTELimit += totalSizeOfAllDirsLTE(child, limit)
	}
	return totalSizeOfAllDirsLTELimit
}

func populateSizes(tree *dir) int {
	totalFileSize := sumOfFiles(tree.files)
	totalChildDirSize := 0
	for _, child := range tree.childDirs {
		totalChildDirSize += populateSizes(child)
	}
	tree.totalSize = totalFileSize + totalChildDirSize
	return tree.totalSize
}

func sumOfFiles(files []file) int {
	totalFileSize := 0
	for _, file := range files {
		totalFileSize += file.size
	}
	return totalFileSize
}

func buildDirTree(lines []string) *dir {
	root = dir{name: "/"}
	currPosition := &root
	for _, line := range lines {
		if isCd(line) {
			currPosition = cd(currPosition, line)
			continue
		}
		if isLs(line) {
			continue
		}
		addItem(currPosition, line)
	}
	return &root
}

func addItem(currPosition *dir, line string) {
	propertyAndName := strings.Split(line, " ")
	if propertyAndName[0] == "dir" {
		currPosition.addChildIfNotPresent(propertyAndName[1])
		return
	}
	currPosition.addFileIfNotPresent(propertyAndName[0], propertyAndName[1])
}

func isLs(line string) bool {
	return strings.Split(line, " ")[1] == "ls"
}

func cd(currPosition *dir, line string) *dir {
	newDir := strings.Split(line, " ")[2]
	if newDir == ".." {
		return currPosition.parent
	}
	if newDir == "/" {
		return &root
	}
	for _, child := range currPosition.childDirs {
		if child.name == newDir {
			return child
		}
	}
	panic("could not cd to " + newDir)
}

func isCd(line string) bool {
	return strings.Split(line, " ")[1] == "cd"
}
