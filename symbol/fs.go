package symbol

import "os"

func IsFileExist(s string) bool {
	if _, err := os.Stat(s); os.IsNotExist(err) {
		return false
	}

	return true
}
