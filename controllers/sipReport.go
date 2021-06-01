package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var sipReportName = "sipReport"

type MscSipInfo struct {
	Dest      string `json:"destination"`
	Source    string `json:"source"`
	FirstLine string `json:"firstLine"`
	Content   string `json:"content"`
}

type MscSipInfoSlice struct {
	Data []MscSipInfo
}

type SipController struct {
	BaseController
}

func (this *SipController) Get() {
	mscbody := this.ReadSipReport()
	this.Data["mscbody"] = mscbody
	this.TplName = "sip/index.tpl"
	this.Render()
}

func (this *SipController) ReadSipReport() (result []string) {
	var details []string
	fi, err := os.Open("static\\" + sipReportName)
	if err != nil {
		fmt.Println("Failed to open")
		fmt.Println(err)
	}
	// read file
	fileContext, err := ioutil.ReadAll(fi)
	defer fi.Close()
	// put json to struct
	var MscContextBody MscSipInfoSlice
	// var sip_msc_children = "\"%s\"=>>\"%s\" [label=\"%s\", title=\"%s\", id=\"%d\", url=\"javascript:show('%d')\"];"
	nameList := make([]string, 0)
	if err := json.Unmarshal([]byte(fileContext), &MscContextBody); err == nil {
		for index, sipSignal := range MscContextBody.Data {
			// fmt.Println(fmt.Sprintf(sip_msc_children, sipSignal.Source, sipSignal.Dest, sipSignal.FirstLine, index, index))
			nameList = append(nameList, "\""+sipSignal.Source+"\"")
			nameList = append(nameList, "\""+sipSignal.Dest+"\"")
			details = append(details, "\""+sipSignal.Source+"\"=>>\""+sipSignal.Dest+"\" [label=\""+sipSignal.FirstLine+"\", title=\""+strconv.Itoa(index)+"\", id=\""+strconv.Itoa(index)+"\", url=\"javascript:show('"+strconv.Itoa(index)+"')\"];")
		}
	} else {
		// fmt.Println("getResultText:", err)
	}
	nameList = this.RemoveRepeatedElement(nameList)
	var tempNameStr string
	for _, name := range nameList {
		tempNameStr = tempNameStr + name + ","
	}
	tempNameStr = tempNameStr[0 : len(tempNameStr)-1]
	result = append(result, tempNameStr+";\n")
	result = append(result, "|||;\n")
	result = append(result, details...)
	result = append(result, "|||;\n")
	return
}

func (this *SipController) RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func (c *SipController) GetSipContextDetail() {
	sipContextId := c.Ctx.Input.Param(":id")
	fmt.Println(string(sipContextId))
	var target string
	fi, err := os.Open("static\\" + sipReportName)
	if err != nil {
		fmt.Println("Failed to open")
		fmt.Println(err)
	}
	// read file
	fileContext, err := ioutil.ReadAll(fi)
	defer fi.Close()
	// put json to struct
	var MscContextBody MscSipInfoSlice
	if err := json.Unmarshal([]byte(fileContext), &MscContextBody); err == nil {
		for index, sipSignal := range MscContextBody.Data {
			if strconv.Itoa(index) == sipContextId {
				target = string(sipSignal.Content)
				break
			}
			index = index + 1
		}
	} else {
		// fmt.Println("getResultText:", err)
	}
	c.SuccessJson(target)
}
