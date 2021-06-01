package controllers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type UserInfoController struct {
	web.Controller
}

func (this *UserInfoController) Get() {
	file_test := this.getDeviceFileList("./static/")
	this.Data["Files"] = file_test
	this.TplName = "user/index.tpl"
	this.Render()
}

func (this *UserInfoController) GetSystemInfo(filePath string) (result []string) {
	fp, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()
	buff := make([]byte, 1024)
	for {
		len, err := fp.Read(buff)
		if err == io.EOF || len < 0 {
			break
		}
		result = append(result, string(buff[:len]))
	}
	return result
}

func (d *UserInfoController) getDeviceFileList(path string) []string {
	var temp []string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if strings.Contains(path, "SystemInfo") {
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
