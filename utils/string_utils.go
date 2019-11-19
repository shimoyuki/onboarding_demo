package utils

func IsStringEmpty(str string) bool {
	bytes := []byte(str)
	strLen := len(str)
	if strLen == 0 {
		return true
	}
	for i := 0; i < strLen; i++ {
		if bytes[i] != ' ' && bytes[i] != '\n' && bytes[i] != '\t' {
			return false
		}
	}
	return true
}

func IsAnyStringEmpty(strs ...string) bool {
	for _, str := range strs {
		if IsStringEmpty(str) {
			return true
		}
	}
	return false
}
