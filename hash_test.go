package hash

import (
	"io/ioutil"
	"os"
	"testing"
)

const (
	testData = "example"
	testHash = "w0mcJylzCn+AfvuGdqkty2+KP48="
)

func TestBytes(t *testing.T) {
	have := Bytes([]byte(testData))
	if have != testHash {
		t.Errorf("hashes do not match: %s != %s", have, testHash)
	}
}

func TestFile(t *testing.T) {
	file := "/tmp/go-hash-test"
	ioutil.WriteFile(file, []byte(testData), 0644)
	defer func() {
		os.Remove(file)
	}()

	have, err := File(file)
	if err != nil {
		t.Errorf("failed to hash file %s", file)
	} else {
		if have != testHash {
			t.Errorf("hashes do not match: %s != %s", have, testHash)
		}
	}
}
