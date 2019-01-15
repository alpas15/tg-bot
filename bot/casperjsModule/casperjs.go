package casperjsModule

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

type CasperJsConf struct {
	ExexFile string `json:"execution_file"`
}

const (
	templateName = "CasperJS"
)

type Casper struct {
	script   *os.File
	template *template.Template
	Output   string
}

type Url struct {
	Url string
}

func (u Url) Execute() {}

type CasperData interface {
	Execute()
}

type CasperTemplate struct {
	TemplateFile string
	Data         CasperData
}

func SetSettings(conf CasperJsConf) {
	os.Setenv("CASPERJS_EXEC", conf.ExexFile)
}

func (tpl *CasperTemplate) getPath() string {
	return filepath.FromSlash(tpl.TemplateFile)
}

func (c *Casper) create() error {

	//create new temp file inside temp folder. casperjs will ran this file.
	var err error
	c.script, err = ioutil.TempFile(os.TempDir(), "go_casperjs_")

	if err != nil {
		return errors.New("GoCasperjs:: Can't create temp file" + err.Error())
	}

	//create new template, throw erorr if can't
	c.template = template.New(templateName)

	return nil //return nil - no error.
}

func (c *Casper) loadTemplate(cTemplate CasperTemplate) error {

	//load template files.
	template, err := template.ParseFiles(cTemplate.getPath())
	if err != nil {
		return errors.New("GoCasperjs:: can't parse template files: " + err.Error())
	}

	//if data not supplied don't execute data.
	if cTemplate.Data != nil {
		cTemplate.Data.Execute()
	}

	//execute template and add into script.
	err = template.Execute(c.script, cTemplate.Data)
	if err != nil {
		return errors.New("GoCasperjs:: cant execute template with data: " + err.Error())
	}

	return nil
}

func (c *Casper) parseString(content string, data CasperData) error {
	var err error
	if len(content) > 0 && c.template != nil {
		c.template, err = c.template.Parse(content)
	} else {
		err = errors.New("GoCasperjs:: cant execute template with data: " + err.Error())
	}
	if err != nil {
		return errors.New("GoCasper:: can't parse string with data: " + err.Error())
	}

	err = c.template.Execute(c.script, data)
	if err != nil {
		return errors.New("GoCasperjs:: can't execute string with data: " + err.Error())
	}

	return nil
}

func (c *Casper) close() {
	c.script.Close()
	os.Remove(c.script.Name())
}

func (c *Casper) run() error {
	//TODO: вынести флаги в конфиг

	var err error
	var out []byte

	out, err = exec.Command("timeout", "180s",
		"casperjs",
		"--ignore-ssl-errors=true",
		c.script.Name()).Output()
	if err != nil {
		return errors.New("GoCasperjs:: failed to run casperjs: " + err.Error())
	}

	c.Output = string(out)

	return nil
}

func Loader(url, tplFile string) (text string) {
	casper := Casper{}
	tpl := CasperTemplate{
		TemplateFile: tplFile,
		Data:         Url{Url: url},
	}

	err := casper.create()

	if err == nil {
		defer casper.close()
	}

	if err != nil {
		fmt.Println(url, err)
		return
	}

	err = casper.loadTemplate(tpl)

	if err != nil {
		fmt.Println(url, err)
		return
	}

	err = casper.parseString(`casper.run();`, nil)

	if err != nil {
		fmt.Println(url, err)
		return
	}

	err = casper.run()

	if err != nil {
		fmt.Println(url, err)
		return
	}

	text = casper.Output

	return
}
