package hash

import (
	"crypto/sha1"
	"encoding/base64"
	"io/ioutil"
)

// Return a hash for some data.
func Bytes(data []byte) string {
	hash := sha1.New()
	hash.Write(data)
	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}

// Return the hash of a file.
func File(path string) (string, error) {
	if data, err := ioutil.ReadFile(path); err == nil {
		return Bytes(data), nil
	} else {
		return "", err
	}
}
