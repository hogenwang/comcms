package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"time"
)

//Webuploader
type WebuploadController struct {
	beego.Controller
}

func (this *WebuploadController) WebUploader() {
	action := this.Ctx.Input.Query("action")
	switch {
	case action == "UpLoadFile":
		WebUploadImage(&this.Controller)
	}
}

//上传图片
func WebUploadImage(this *beego.Controller) {
	//filename := this.Input().Get("Filename")
	f, h, _ := this.GetFile("Filedata")
	filename := h.Filename
	f.Close() //关闭，减少缓存
	ext := filename[strings.LastIndex(filename, ".")+1:]
	//获取扩展名

	if !strings.Contains(allowImageType, ext) {
		fmt.Println(filename)
		this.Ctx.WriteString("{\"state\":\"0\"}")
		this.StopRun()
	}
	newname := strconv.FormatInt(time.Now().Unix(), 10) + "_" + filename
	err := this.SaveToFile("Filedata", uploadimage+newname)
	state := "SUCCESS"
	if err != nil {
		fmt.Println(err)
		state = "0"
	}
	state = "1"
	url := website + uploadimage + newname
	//this.Ctx.WriteString("{'original':'" + filename + "','url':'" + url + "','title':'" + this.Input().Get("pictitle") + "','state':'" + state + "'}")
	//this.Ctx.WriteString("{\"state\": \"" + state + "\", \"url\": \"" + url + "\", \"title\": \"\",\"original\": \"" + filename + "\"}")
	this.Ctx.WriteString("{\"status\": " + state + ", \"msg\": \"上传文件成功！\", \"name\": \"1.jpg\", \"path\": \"" + url + "\", \"thumb\": \"" + url + "\", \"size\": 0, \"ext\": \"" + ext + "\"}")
	this.StopRun()
}
