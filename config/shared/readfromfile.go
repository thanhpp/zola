package shared

import (
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/spf13/viper"
)

var (
	ErrInvalidFormat = errors.New("invalid filepath, must .../<name>.<ext>")
	ErrNilExpected   = errors.New("nil expected")
)

func ReadFromFile(path string, expected interface{}) error {
	if expected == nil {
		return ErrNilExpected
	}

	file := filepath.Base(path)

	// 0: filename, 1: file ext
	fileSl := strings.Split(file, ".")
	if len(fileSl) != 2 {
		return ErrInvalidFormat
	}

	// read the config file using viper
	v := viper.New()
	v.SetConfigName(fileSl[0])
	v.SetConfigType(fileSl[1])
	v.AddConfigPath(".")     // current folder
	v.AddConfigPath("..")    // depth 1 - test file
	v.AddConfigPath("../..") // depth 2 - test file

	if err := v.ReadInConfig(); err != nil {
		return errors.WithMessage(err, "read config")
	}

	if err := v.Unmarshal(expected); err != nil {
		return errors.WithMessage(err, "unmarshal")
	}

	return nil
}
