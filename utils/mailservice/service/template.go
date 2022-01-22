package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	htmltemplate "html/template"
	"io/ioutil"
	"os"
	"time"

	"github.com/vfluxus/mailservice/repository/entity"
)

type TemplateService struct{}

type iTemplateService interface {
	Validate(template *entity.Template) (html []byte, err error)
	Parse(template *entity.Template, vars []*entity.Variable) (html []byte, err error)
}

var templateSrv iTemplateService = new(TemplateService)

func GetTemplate() iTemplateService {
	return templateSrv
}

const dir = "validate"

func (t *TemplateService) Validate(template *entity.Template) (html []byte, err error) {
	filePath, err := t.makeTempFile(dir, template)
	if err != nil {
		return nil, err
	}

	tpl, err := htmltemplate.ParseFiles(filePath)
	if err != nil {
		return nil, err
	}

	// make map for data input
	var (
		dataMap = make(map[string]interface{})
		vars    []*entity.Variable
	)

	if err = json.Unmarshal(template.Variables, &vars); err != nil {
		return nil, err
	}

	for i := range vars {
		dataMap[vars[i].Name] = vars[i].Default
	}

	var buffers = new(bytes.Buffer)
	if err = tpl.Execute(buffers, dataMap); err != nil {
		return nil, err
	}

	return buffers.Bytes(), nil
}

func (t *TemplateService) Parse(template *entity.Template, vars []*entity.Variable) (html []byte, err error) {
	filePath, err := t.makeTempFile(dir, template)
	if err != nil {
		return nil, err
	}
	defer t.RemoveTempFile(filePath)

	tpl, err := htmltemplate.ParseFiles(filePath)
	if err != nil {
		return nil, err
	}

	// make map for data input
	var (
		dataMap = make(map[string]interface{})
	)

	for i := range vars {
		dataMap[vars[i].Name] = vars[i].Value
	}

	var buffers = new(bytes.Buffer)
	if err = tpl.Execute(buffers, dataMap); err != nil {
		return nil, err
	}

	return buffers.Bytes(), nil
}

func (t *TemplateService) makeTempFile(dir string, template *entity.Template) (filePath string, err error) {
	// pre-exec check
	if template == nil {
		return "", errors.New("Nil template")
	}

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}

	var (
		fileName = fmt.Sprintf("%s-%d.html", template.Name, time.Now().UnixNano())
	)
	filePath = fmt.Sprintf("%s/%s", dir, fileName)

	if err := ioutil.WriteFile(filePath, []byte(template.Content), os.ModePerm); err != nil {
		return "", err
	}

	return filePath, nil
}

func (t *TemplateService) RemoveTempFile(filePath string) (err error) {
	if err := os.Remove(filePath); err != nil {
		return err
	}
	return nil
}
