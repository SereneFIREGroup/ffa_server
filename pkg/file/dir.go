package file

import "os"

// CheckDirExist check if dir exist
func CheckDirExist(dirPath string) bool {
	_, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// MkDir make dir
func MkDir(dirPath string) error {
	return os.MkdirAll(dirPath, 0755)
}

// CheckDirAndMkDir check if dir exist, if not, make dir
func CheckDirAndMkDir(dirPath string) error {
	if !CheckDirExist(dirPath) {
		return MkDir(dirPath)
	}
	return nil
}
