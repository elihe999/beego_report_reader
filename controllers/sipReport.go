package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type MscSipInfo struct {
	Desc      string `json:"destination"`
	Source    string `json:"source"`
	FirstLine string `json:"firstLine"`
	Content   string `json:"content"`
}

type SipController struct {
	web.Controller
}

func (this *SipController) Get() {
	mscbody := this.ReadSipReport()
	fmt.Println(mscbody)
	this.Data["mscbody"] = mscbody
	this.TplName = "sip/index.tpl"
	this.Render()
}

func (this *SipController) ReadSipReport() map[interface{}]interface{} {
	var details []string
	fileName := "sipReport"
	fi, err := os.Open("static\\" + fileName)
	if err != nil {
		fmt.Println("Failed to open")
		fmt.Println(err)
	}
	defer fi.Close()
	// read file
	buf := bufio.NewReader(fi)
	for {
		b, errR := buf.ReadBytes('\n')
		if errR != nil {
			if errR == io.EOF {
				break
			}
			fmt.Println(errR.Error())
		}
		details = append(details, string(b))
	}

	var MscContext MscSipInfo
	mscInterface := make(map[interface{}]interface{})
	for index, value := range details {
		if err := json.Unmarshal([]byte(value), &MscContext); err == nil {
			mscInterface[index] = MscContext
			fmt.Println(MscContext)
		} else {
			// fmt.Println("getResultText:", err)
		}
	}
	return mscInterface
}
