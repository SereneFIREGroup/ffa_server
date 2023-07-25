package env

import "path/filepath"

var basePath = func() string {
	dir, err := filepath.Abs("./")
	if err != nil {
		panic(err)
	}
	return dir
}()

func BasePath() string {
	return basePath
}
