package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/thoas/go-funk"
	"io/ioutil"
	"testing"
)

func TestCopy(t *testing.T) {
	t.Run("copying with offset and limit works", func(t *testing.T) {
		var destBuffer = new(bytes.Buffer)
		var srcBuffer = bytes.NewReader([]byte("Hello World! Happy New Year!"))

		written, err := Copy(srcBuffer, destBuffer, 13, 9)

		assert.Equal(t, 9, written)
		assert.Equal(t, "Happy New", destBuffer.String())
		assert.Nil(t, err)
	})
}

func updateFileContent(t *testing.T, path string, content string) {
	err := ioutil.WriteFile(path, []byte(content), 0755)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func getStringFromFile(t *testing.T, pathDest string) string {
	buf, err := ioutil.ReadFile(pathDest)
	if err != nil {
		t.Errorf(err.Error())
	}
	return string(buf)
}

const (
	PathSrc = "./files/source.txt"
	PathSrc1 = "./files/source1.txt"
	PathDest = "./files/dest.txt"
)

func TestCopyFiles(t *testing.T) {
	t.Run("copying files with offset and limit works", func(t *testing.T) {
		updateFileContent(t, PathSrc, "Hello World! Happy New Year!")

		written, err := CopyFiles(&Args{
			From:   PathSrc,
			To:     PathDest,
			Offset: 13,
			Limit:  9,
		})

		assert.Nil(t, err)
		assert.Equal(t, 9, written)
		assert.Equal(t, "Happy New", getStringFromFile(t, PathDest))
	})

	t.Run("copying fails when source is absent", func(t *testing.T) {
		_, err := CopyFiles(&Args{
			From:   "xxx",
			To:     PathDest,
			Offset: 0,
			Limit:  0,
		})

		assert.NotNil(t, err)
	})

	t.Run("copying works when destination file is absent", func(t *testing.T) {
		updateFileContent(t, PathSrc, "Hello!")
		pathDestNotExist := PathDest + funk.RandomString(16) + ".txt"

		_, err := CopyFiles(&Args{
			From:   PathSrc,
			To:     pathDestNotExist,
			Offset: 0,
			Limit:  0,
		})

		assert.Nil(t, err)
		assert.Equal(t, "Hello!", getStringFromFile(t, pathDestNotExist))
	})

	t.Run("reset destination file content before copying", func(t *testing.T) {
		updateFileContent(t, PathSrc, "Hello!")
		_, _ = CopyFiles(&Args{
			From:   PathSrc,
			To:     PathDest,
			Offset: 0,
			Limit:  0,
		})
		assert.Equal(t, "Hello!", getStringFromFile(t, PathDest))

		updateFileContent(t, PathSrc1, "Bye!")
		_, _ = CopyFiles(&Args{
			From:   PathSrc1,
			To:     PathDest,
			Offset: 0,
			Limit:  0,
		})
		assert.Equal(t, "Bye!", getStringFromFile(t, PathDest))
	})
}

