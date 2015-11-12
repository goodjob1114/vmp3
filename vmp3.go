package vmp3

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func walkpath(path string, f os.FileInfo, err error) error {

	filetype := getFiletype(path)

	if isNotVideo(filetype) {
		return nil
	}

	mp3file := getFilename(path, filetype)

	cmd := fmt.Sprintf("ffmpeg -i '%s' -y '%s'", path, mp3file)
	fmt.Println("cmd::", cmd)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Convert::MP3::OK", string(out))

	return nil
}

func getFiletype(file string) string {
	f := strings.Split(file, ".")
	if len(f) == 1 {
		return ""
	}
	return f[len(f)-1]
}

func getFilename(file string, filetype string) string {
	f := strings.Split(file, fmt.Sprintf(".%s", filetype))
	return fmt.Sprintf("%s.mp3", f[0])
}

func isVideo(filetype string) bool {
	filetype = strings.ToLower(filetype)
	matched, _ := regexp.MatchString("mkv|mp4|avi", filetype)
	return matched
}

func isNotVideo(filetype string) bool {
	return !isVideo(filetype)
}

func GoMp3(root string) {
	err := filepath.Walk(root, walkpath)
	if err != nil {
		fmt.Println(err.Error())
	}
}
