package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

type Result struct {
	Type    string `json:"type"`
	Desc    string `json:"desc"`
	Info    string `json:"info"`
	Details string `json:"details"`
	Status  string `json:"status"`
	Time    string `json:"time"`
}

func (c *MainController) ReadFile(path string) (details []string, err error) {
	// open file
	fi, err := os.Open(path)
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

func (c *MainController) getResultText() {
	content, err := c.ReadFile("./static/Result")
	if err != nil {
		fmt.Println(err)
	}
	var someContext Result
	mapInterface := make(map[interface{}]interface{})

	for index, value := range content {
		if err := json.Unmarshal([]byte(value), &someContext); err == nil {
			// fmt.Println(someContext)
			if someContext.Type != "If" && someContext.Type != "While" && someContext.Type != "Var" && someContext.Type != "Foreach" {
				if someContext.Type == "Screenshot" {
					someContext.Details = "<img src=\"static/" + someContext.Details + "\" />"
				}
				mapInterface[index] = someContext
			}
			// json.Unmarshal(js, &m)
		} else {
			// fmt.Println("getResultText:", err)
		}
	}
	c.Data["json"] = mapInterface
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "ylhe@grandstream.cn"
	this.getResultText()
	this.TplName = "result/index.tpl"
	this.Render()
}
