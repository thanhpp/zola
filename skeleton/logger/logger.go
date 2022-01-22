package logger

import (
	"encoding/json"

	"github.com/pinezapple/LibraryProject20201/skeleton/model"

	log "github.com/sirupsen/logrus"
)

func MustGet(servName string) (lg *model.LogFormat) {
	logger := &model.LogFormat{ServiceName: servName}

	return logger
}

// LogInfo information logging
func LogInfo(lg *model.LogFormat, message string) {
	if lg == nil {
		return
	}

	lg.Action = lg.Action + " " + message
	js, _ := json.Marshal(lg)
	log.Infof("%s", js)
	lg.Action = ""
}

// LogErr error logging
func LogErr(lg *model.LogFormat, err error) {
	if err == nil {
		return
	}

	js, _ := json.Marshal(&model.LogFormat{Err: err.Error()})
	log.Errorf("%s", js)
	lg.Err = nil
}

// LogWarning warning logging
func LogWarning(lg *model.LogFormat) {
	if lg == nil {
		return
	}

	js, _ := json.Marshal(lg)
	log.Infof("%s", js)
}
