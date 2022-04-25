package tools

import "os"

func PathExists(path string) {
	_, err := os.Stat(path)
	if err == nil {
		return
	}else {
		os.Mkdir(path,0777)
	}
}
