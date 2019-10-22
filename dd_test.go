package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
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

		buf, err := ioutil.ReadFile(srcDest)
		if err != nil {
			t.Errorf(err.Error())
		}

		assert.Equal(t, "Happy New", string(buf))
	})

	t.Run("copying fails when one of the files is absent", func(t *testing.T) {
		srcPath := "./files/source.txt"
		srcDest := "./files/dest.txt"

		_, err := CopyFiles(&Args{
			From:   "xxx",
			To:     srcDest,
			Offset: 13,
			Limit:  9,
		})

		assert.NotNil(t, err)

		_, err = CopyFiles(&Args{
			From:   srcPath,
			To:     "xxx",
			Offset: 13,
			Limit:  9,
		})

		assert.NotNil(t, err)
	})
}

