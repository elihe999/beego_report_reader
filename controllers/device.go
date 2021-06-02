package controllers

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type DeviceController struct {
	web.Controller
}

func (this *DeviceController) Get() {
	file_test := this.getDeviceFileList("./static/")
	this.Data["Files"] = file_test
	this.TplName = "device/index.tpl"
	this.Render()
}

func (this *DeviceController) ListDeviceStep() {
	a := this.Ctx.Input.Param(":id")
	content, err := this.ReadFile(a)
	if err != nil {
		fmt.Println(err)
	}
	this.Data["Files"] = content
	this.TplName = "device/teststeps.tpl"
	this.Render()
}

func (d *DeviceController) getDeviceFileList(path string) []string {
	var temp []string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if strings.Contains(path, "testSteps.txt") {
			temppath := path[len("static\\"):]
			temp = append(temp, temppath)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	} else {
		return temp
	}
	return nil
}

func (c *DeviceController) ReadFile(path string) (details []string, err error) {
	// open file
	fi, err := os.Open("static\\" + path)
	if err != nil {
		fmt.Println("Failed to open")
		fmt.Println(err)
	}
	defer fi.Close()
	// read file
	buf := bufio.NewReader(fi)
	for {
		//遇到\n结束读取
		b, errR := buf.ReadBytes('\n')
		if errR != nil {
			if errR == io.EOF {
				break
			}
			fmt.Println(errR.Error())
		}
		// fmt.Println(string(b))
		details = append(details, string(b))
	}
	return details, err
}

func (c *DeviceController) FileDown() {
	name := c.Ctx.Input.Param(":id")
	c.Ctx.Output.Download("static/" + name)
}
