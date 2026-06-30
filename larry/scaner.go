package main

import (
	"os"
	"path/filepath"
)

// so the func scaner will scan the given dir and put the files in map  the map will be [old abspath]=>[new abdpath] the logique of chosing the path will be indside
func Scanner(dir string, cfg Conf) (map[string]string, error) {
	var files map[string]string
	files = make(map[string]string)
	all, err := os.ReadDir(dir)
	if err != nil {
		return files, err
	}
	for _, v := range all {
		if !v.IsDir() {
			name := filepath.Ext(v.Name())
			absdir, _ := filepath.Abs(dir)
			newdir, ok := cfg[name]
			if ok == false {
				newdir = "other"
			}

			newpath := filepath.Join(absdir, newdir, v.Name())
			oldpath := filepath.Join(absdir, v.Name())
			files[oldpath] = newpath
		}
	}
	return files, err

}
