package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func move(files map[string]string) error {
	var err error
	println(Blue + "larry will move those files" + Reset)

	for oldpath, newpath := range files {
		fmt.Printf(Bold+"%-30s => %s\n"+Reset, filepath.Base(oldpath), filepath.Dir(newpath))
	}
	fmt.Printf(Blue+"total moved files %v"+Reset, len(files))
	println(Blue + "do you want those changes? [y/n]" + Reset)
	var a string
	fmt.Scan(&a)
	if a == "n" || a == "N" {
		fmt.Println(Blue + "ok larry will stop working" + Reset)
		return err
	}
	for oldpath, newpath := range files {
		err = os.MkdirAll(filepath.Dir(newpath), 0755)
		if err != nil {
			return err
		}
		err = os.Rename(oldpath, newpath)
		if err != nil {
			return err
		}
	}
	println(Green + "larry finshed and ask you for tip" + Reset)
	println(` _                          
| |    __ _ _ __ _ __ _   _ 
| |   / _` + "`" + ` | '__| '__| | | |
| |__| (_| | |  | |  | |_| |
|_____\__,_|_|  |_|   \__, |
                       |___/
The file organizer`)
	return err
}
