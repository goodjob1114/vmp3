package vmp3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFiletype(t *testing.T) {

	assert.Equal(t, "mkv", getFiletype("abc.mkv"), "they should be equal")

	assert.NotEqual(t, "mkv", getFiletype("abc.mp4"), "they should not be equal")

	assert.Equal(t, "", getFiletype("abcmp4"))
}

func TestGetFilename(t *testing.T) {

	assert.Equal(t, "/home/goodjob/tmp/test.mp3", getFilename("/home/goodjob/tmp/test.MKV", "MKV"))

	assert.NotEqual(t, "/home/foo.mp3", getFilename("/home/foo.mp4", "MKV"))
}

func TestIsVideo(t *testing.T) {

	assert.True(t, isVideo("mp4"), "this should be true")

	assert.False(t, isVideo("mp3"), "this should be false")
}

func TestIsNotVideo(t *testing.T) {

	assert.False(t, isNotVideo("mp4"), "this should be false")

	assert.True(t, isNotVideo("mp3"), "this should be true")
}
