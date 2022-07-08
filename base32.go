package ids

import "encoding/base32"

var (
	lowerCaseEncodeHex = "0123456789abcdefghijklmnopqrstuv"
	base32enc          = base32.NewEncoding(lowerCaseEncodeHex).WithPadding(base32.NoPadding)
)

func EncodeToBase32String(p []byte) string {
	return base32enc.EncodeToString(p)
}

func DecodeBase32String(str string) ([]byte, error) {
	return base32enc.DecodeString(str)
}
