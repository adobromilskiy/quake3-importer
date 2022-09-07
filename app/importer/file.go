package importer

import (
	"os"
)

func getFiles(path string) (res []string, err error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return res, err
	}

	for _, file := range files {
		if file.IsDir() {
			files2, _ := getFiles(path + "/" + file.Name())
			res = append(res, files2...)
		} else {
			res = append(res, path+"/"+file.Name())
		}
	}

	return res, nil
}
