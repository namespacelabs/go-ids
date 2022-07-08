package ids

import "github.com/jxskiss/base62"

func EncodeToBase62String(p []byte) string {
	return base62.EncodeToString(p)
}

func DecodeBase62String(str string) ([]byte, error) {
	return base62.DecodeString(str)
}
