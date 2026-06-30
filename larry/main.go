package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	Bold   = "\033[1m"
)
const defaultConfig = `{
    "_comment": "Extensions must start with a dot and be separated by spaces.",
    "Images": ".jpg .jpeg .png .gif .webp .svg",
    "Documents": ".pdf .doc .docx .txt .odt",
    "Music": ".mp3 .wav .flac .ogg .aac",
    "Videos": ".mp4 .mkv .avi .mov .webm",
    "Archives": ".zip .rar .7z .tar .gz",
    "Code": ".go .c .cpp .h .hpp .java .py .js .ts .html .css .json .xml .yml .yaml"
    
}`

type Conf map[string]string

// in the conf the it will lock like the_newdir_name=>"ext1 ext2 ext3 ....."
func LoadConf() (Conf, error) {
	var rawcfg Conf
	var cfg Conf
	cfg = make(Conf)
	confpath, err := os.UserConfigDir()
	if err != nil {
		return rawcfg, err
	}

	confpath = filepath.Join(confpath, "larry", "config.json")
	data, err := os.ReadFile(confpath)

	if err != nil {
		if os.IsNotExist(err) {
			// Create the default config
			err = os.MkdirAll(filepath.Dir(confpath), 0755)
			if err != nil {
				return rawcfg, err
			}
			err = os.WriteFile(confpath, []byte(defaultConfig), 0644)
			if err != nil {
				return rawcfg, err
			}

		} else {

			return rawcfg, err
		}

	}
	data, err = os.ReadFile(confpath)
	if err != nil {
		return rawcfg, err
	}

	err = json.Unmarshal(data, &rawcfg)

	if err != nil {
		return rawcfg, err
	}
	//we will swap the map
	for k, v := range rawcfg {
		if k != "_comment" {
			exts := strings.Split(v, " ")
			for _, v2 := range exts {
				cfg[v2] = k
			}
		}
	}

	return cfg, err
}
func main() {
	if len(os.Args) < 2 {
		println(Red + "larry need you to give him a dir to orginize" + Reset)
		println(Blue + "usage:<directory>" + Reset)
		return
	}
	dir := os.Args[1]
	cfg, err := LoadConf()
	if err != nil {
		println(Red + "sorry your conf does not exist or malformated" + Reset)
		return
	}
	files, err := Scanner(dir, cfg)
	if err != nil {
		println(Yellow + "ammmm i think this directory deos not exist ,or i dont have the permission to see it" + Reset)
		return
	}
	if len(files) == 0 {
		println(Yellow + "no files to be orgnize here" + Reset)
		return
	}
	err = move(files)
	if err != nil {
		println(Red + "larry i think i dont have the permision to write the directory" + Reset)
		return
	}
}
