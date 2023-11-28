package directory

import "os"

type RestoreDirectoryFunc func()

// ChangeDirectory returns the current directory prior to changing and
// a function that can be used to restore the current directory.
func ChangeDirectory(dir string) (string, RestoreDirectoryFunc, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", nilRestoreDirectory, err
	}
	err = os.Chdir(dir)
	if err != nil {
		return currentDir, nilRestoreDirectory, err
	}
	return currentDir, getRestoreDirectory(currentDir), nil
}

func nilRestoreDirectory() {}

// getRestoreDirectory returns a function meant to be called by defer
func getRestoreDirectory(dir string) RestoreDirectoryFunc {
	return func() {
		_ = os.Chdir(dir)
	}
}
