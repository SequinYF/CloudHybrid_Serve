package model

import (
	"ch_server/src/util"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func RegisterHandler(c *gin.Context) {
	name := c.PostForm("name")
	passwd := c.PostForm("passwd")
	dir := c.PostForm("dir")
	fmt.Println("debug",name, passwd, dir)

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.Required(name, "passwd").Message("密码不能为空")
	valid.MaxSize(name, 100, "passwd").Message("密码最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(name, "dir").Message("目录不能为空")
	valid.MaxSize(name, 200, "dir").Message("目录路径最长为200字符")
	code := util.INVALID_PARAMS
	if !valid.HasErrors() {
		if !ExistByName(name) {
			code = util.SUCCESS
			AddUser(name, passwd, dir)
			BuildDir(name, passwd)
		} else {
			code = util.ERROR_EXIST_NAME
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  util.GetMsg(code),
		"data": make(map[string]string),
	})
}


func ListHandler(c *gin.Context) {
	name := c.Query("name")
	name = "/Users/"+name
	fileInfoArr, err := ioutil.ReadDir(name)
	code := util.SUCCESS
	if err != nil {
		code = util.ERROR_NOT_EXIST_NAME
		log.Println("open file wrong", err)
	}
	locals := make(map[string]interface{})
	var files []string
	for _, fileInfo := range fileInfoArr {
		files = append(files, fileInfo.Name())
	}
	locals["files"] = files

	c.HTML(http.StatusOK, "../src/views/list.html", gin.H{
		"code": code,
		"msg":  util.GetMsg(code),
		"files": locals,
	})

}


func BuildDir(name, passwd string)  {
	command := `/Users/sequin_yf/go/src/ch_server/src/model/register.sh`
	cmd := exec.Command("/bin/bash", command, name, passwd)
	output, err := cmd.Output()
	if err != nil {
		log.Printf("Execute Shell:%s failed with error:%s\n", command, err)
		return
	}

	log.Printf("Execute Shell:%s finished with output:\n%s\n", command, string(output))
}