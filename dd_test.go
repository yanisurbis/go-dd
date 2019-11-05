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

func updateFilesContent(srcPath string, destPath string, srcContent string) error {
	err := ioutil.WriteFile(srcPath, []byte(srcContent), 0755)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(destPath, []byte(""), 0755)
	if err != nil {
		return err
	}
	return nil
}

func getStringFromFile(t *testing.T, pathDest string) string {
	buf, err := ioutil.ReadFile(pathDest)
	if err != nil {
		t.Errorf(err.Error())
	}
	return string(buf)
}

func TestCopyFiles(t *testing.T) {
	t.Run("copying files with offset and limit works", func(t *testing.T) {
		srcPath := "./files/source.txt"
		srcDest := "./files/dest.txt"
		srcContent := "Hello World! Happy New Year!"

		err := updateFilesContent(srcPath, srcDest, srcContent)
		if err != nil {
			t.Errorf(err.Error())
		}

		written, err := CopyFiles(&Args{
			From:   srcPath,
			To:     srcDest,
			Offset: 13,
			Limit:  9,
		})

		assert.Nil(t, err)
		assert.Equal(t, 9, written)
		assert.Equal(t, "Happy New", getStringFromFile(t, srcDest))
	})

	t.Run("copying fails when source is absent", func(t *testing.T) {
		pathDest := "./files/dest.txt"

		_, err := CopyFiles(&Args{
			From:   "xxx",
			To:     pathDest,
			Offset: 13,
			Limit:  9,
		})

		assert.NotNil(t, err)
	})

	t.Run("copying works when destination file is absent", func(t *testing.T) {
		pathSrc := "./files/source.txt"
		pathDestNotExist := "./files/dest-" + funk.RandomString(16)

		written, err := CopyFiles(&Args{
			From:   pathSrc,
			To:     pathDestNotExist,
			Offset: 13,
			Limit:  9,
		})

		assert.Nil(t, err)
		assert.Equal(t, 9, written)
		assert.Equal(t, "Happy New", getStringFromFile(t, pathDestNotExist))
	})
}

