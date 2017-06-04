package p609

import "strings"

func getfile(path string) map[string]string {
	names := strings.Split(path, " ")
	dir := names[0]
	names = names[1:]
	dir = dir + "/"
	res := make(map[string]string)
	for _, name := range names {
		idx := strings.Index(name, "(")
		realname := name[:idx]
		content := name[idx+1 : len(name)-1]
		res[dir+realname] = content
	}
	return res
}

func findDuplicate(paths []string) [][]string {

	fs := make(map[string]map[string]struct{})

	for _, p := range paths {
		files := getfile(p)
		for k, v := range files {
			if fs[v] == nil {
				fs[v] = make(map[string]struct{})
			}
			fs[v][k] = struct{}{}
		}
	}

	res := make([][]string, 0)
	for _, files := range fs {
		if len(files) > 1 {
			dup := make([]string, 0, len(files))
			for f := range files {
				dup = append(dup, f)
			}
			res = append(res, dup)
		}
	}
	return res
}
