package p588

import (
	"bytes"
	"sort"
	"strings"
)

/**
Design an in-memory file system to simulate the following functions:

ls: Given a path in string format. If it is a file path, return a list that only contains this file's name. If it is a directory path, return the list of file and directory names in this directory. Your output (file and directory names together) should in lexicographic order.

mkdir: Given a directory path that does not exist, you should make a new directory according to the path. If the middle directories in the path don't exist either, you should create them as well. This function has void return type.

addContentToFile: Given a file path and file content in string format. If the file doesn't exist, you need to create that file containing given content. If the file already exists, you need to append given content to original content. This function has void return type.

readContentFromFile: Given a file path, return its content in string format.
**/

type node struct {
	name     string
	isDir    bool
	content  *bytes.Buffer
	children map[string]*node
}

type FileSystem struct {
	root *node
}

func (this *FileSystem) lookUp(path string, mk bool) *node {
	path = strings.TrimLeft(path, "/")
	paths := strings.Split(path, "/")
	if len(path) == 0 {
		paths = nil
	}
	cur := this.root
	for _, p := range paths {
		if cur.isDir && cur.children == nil {
			cur.children = make(map[string]*node)
		}
		if v, ok := cur.children[p]; ok {
			cur = v
		} else if mk {
			nd := &node{p, true, nil, nil}
			cur.children[p] = nd
			cur = nd
		} else {
			cur = nil
			break
		}
	}
	return cur
}

func Constructor() FileSystem {
	return FileSystem{root: &node{"", true, nil, make(map[string]*node)}}
}

func (this *FileSystem) Ls(path string) []string {

	nd := this.lookUp(path, false)
	var res []string
	if nd == nil {
		//pass
	} else if nd.isDir {
		res = make([]string, 0, len(nd.children))
		for p := range nd.children {
			res = append(res, p)
		}
		sort.Strings(res)
	} else {
		res = []string{nd.name}
	}
	return res
}

func (this *FileSystem) Mkdir(path string) {
	// paths := strings.Split(path, "/")
	nd := this.lookUp(path, true)
	if nd.children == nil {
		nd.children = make(map[string]*node)
	}
}

func (this *FileSystem) AddContentToFile(filePath string, content string) {
	nd := this.lookUp(filePath, true)
	nd.isDir = false
	if nd.content == nil {
		nd.content = new(bytes.Buffer)
	}
	nd.content.WriteString(content)
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
	nd := this.lookUp(filePath, false)
	if nd == nil || nd.isDir {
		return ""
	}
	return nd.content.String()
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ls(path);
 * obj.Mkdir(path);
 * obj.AddContentToFile(filePath,content);
 * param_4 := obj.ReadContentFromFile(filePath);
 */
