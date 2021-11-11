package application

import (
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type FileHandler struct{}

func (f FileHandler) SaveFile(path string, reader io.Reader) error {
	var (
		dir = filepath.Dir(path)
	)

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	if err != nil {
		// cleanup if err
		os.RemoveAll(path)

		return err
	}

	return nil
}

func (f FileHandler) SaveFileMultipart(path string, mulHeader *multipart.FileHeader) error {
	if mulHeader == nil {
		return nil
	}

	file, err := mulHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	if err := f.SaveFile(path, file); err != nil {
		return err
	}

	return nil
}

func (f FileHandler) Cleanup(paths ...string) {
	for i := range paths {
		os.Remove(paths[i])
	}
}

func (f FileHandler) GetVideoDuration(path string) (float64, error) {
	// ffprobe -v error -show_entries format=duration -of default=noprint_wrappers=1:nokey=1 input.mp4
	var (
		duration float64
	)
	cmd := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1", path)
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	duration, err = strconv.ParseFloat(strings.TrimSpace(string(out)), 64)
	if err != nil {
		return 0, err
	}

	return duration, nil
}
