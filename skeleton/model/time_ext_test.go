package model

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeScan(t *testing.T) {
	assert := assert.New(t)

	ti := &Time{}

	// test with current time
	now := time.Now()
	_ = ti.Scan(now)

	// check conversion too
	x := ti.ConvertToStdTime()
	assert.NotNil(x)
	assert.True(x.Equal(now))

	// try to scan other type
	assert.Equal(ErrTimeExpectedType, ti.Scan(ti))
}

func TestNewTime(t *testing.T) {
	assert := assert.New(t)

	now := time.Now()
	ti := NewTime(now)

	val, err := ti.Value()
	assert.NoError(err)

	x := val.(time.Time)
	assert.True(x.Equal(now))

	// make ti not valid
	ti = &Time{}
	x1, x2 := ti.Value()
	assert.Nil(x1)
	assert.Nil(x2)

	// ti is not valid
	assert.Nil(ti.ConvertToStdTime())
}

func TestUnmarshalJSONTime(t *testing.T) {
	assert := assert.New(t)

	rd1 := []byte(`{"seconds": 12, "nanos":12}`)
	rd2 := []byte(`{"seconds": null, "nanos":12}`)
	rd3 := []byte(`{"nanos":12}`)
	rd4 := []byte(`{"seconds": 12, "nanos":null}`)
	rd5 := []byte(`{"seconds": 12 }`)
	rd6 := []byte(`{}`)
	rd7, err := json.Marshal(NewTime(time.Now()))
	assert.NoError(err)

	mTime := &Time{}
	for _, rd := range [][]byte{rd1, rd2, rd3, rd4, rd5, rd6, rd7} {
		assert.NoError(json.Unmarshal(rd, mTime))
	}

	rd8 := []byte(`{"time":"11/03/2019"}`)
	rd9 := []byte(`{"time":"11/3/2019"}`)
	rd10 := []byte(`{"time":"2019/03/11"}`)
	rd11 := []byte(`{"time": "2019/03/11T15:12:12" }`)

	a := &struct {
		T *Time `json:"time"`
	}{}

	for _, rd := range [][]byte{rd8, rd9, rd10, rd11} {
		assert.NoError(json.Unmarshal(rd, a))
	}
}
