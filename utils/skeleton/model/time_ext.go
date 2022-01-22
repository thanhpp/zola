package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// ErrTimeExpectedType expected type is time.Time but got wrong type
var ErrTimeExpectedType = fmt.Errorf("Expected type is time.Time")

// ErrTimeScanNull nil time scan error
var ErrTimeScanNull = fmt.Errorf("Value is nil that can not be scanned")

var ErrTimeInvalidJSON = fmt.Errorf("Cannot parse JSON message as Time")

// Scan implements the Scanner interface.
func (nt *Time) Scan(value interface{}) (err error) {
	if value == nil {
		return
	}

	t, ok := value.(time.Time)
	if !ok {
		return ErrTimeExpectedType
	}

	// get sec and nano sec
	seconds := t.Unix()
	nanos := int32(t.Sub(time.Unix(seconds, 0)))

	// assign
	nt.Seconds, nt.Nanos = seconds, nanos
	return
}

// UnmarshalJSON unmarshal time from json message
// NOTE: this function is essential for parsing time from engine fraudscore
func (nt *Time) UnmarshalJSON(rd []byte) (err error) {
	// no-op
	if rd == nil {
		return
	}

	// Try to parse as a normal Time struct
	var tmp map[string]*json.RawMessage
	if err = json.Unmarshal(rd, &tmp); err == nil {

		seconds, ok := tmp["seconds"]
		if ok && seconds != nil {
			if err = json.Unmarshal(*seconds, &nt.Seconds); err != nil {
				return ErrTimeInvalidJSON
			}
		}

		nanos, ok := tmp["nanos"]
		if ok && nanos != nil {
			if err = json.Unmarshal(*nanos, &nt.Nanos); err != nil {
				return ErrTimeInvalidJSON
			}
		}

		return
	}
	err = nil

	// Try to parse as other kind of time format
	var t time.Time
	supportedFormat := []string{
		"02/01/2006",
		"02/1/2006",
		"2006-01-02",
		"2006/01/02",
		"2006/01/02T15:04:05",
		"2006-01-02T15:04:05",
		"2006-01-02T15:04:05Z",
	}

	str := string(rd[1 : len(rd)-1])
	for _, format := range supportedFormat {
		if t, err = time.Parse(format, str); err == nil {
			err = nt.Scan(t)
			return
		}
	}

	return fmt.Errorf("%s %s", ErrTimeInvalidJSON.Error(), str)
}

// Value implements the driver Valuer interface.
func (nt Time) Value() (driver.Value, error) {
	if nt.Seconds == 0 && nt.Nanos == 0 {
		return nil, nil
	}
	return time.Unix(int64(nt.Seconds), int64(nt.Nanos)), nil
}

// ConvertToStdTime convert to standard time
func (nt *Time) ConvertToStdTime() *time.Time {
	if nt == nil || nt.Seconds == 0 && nt.Nanos == 0 {
		return nil
	}
	x := time.Unix(int64(nt.Seconds), int64(nt.Nanos))
	return &x
}

// NewTime new Time from standard time
func NewTime(t time.Time) *Time {
	ts := &Time{}
	_ = ts.Scan(t)

	return ts
}
