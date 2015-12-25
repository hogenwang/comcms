package admin

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var allowImageType = "gif|jpeg|jpg|png|bmp"
var allowFileType = "rar|doc|docx|zip|pdf|txt|swf|mkv|avi|rm|rmvb|mpeg|mpg|ogg|mov|wmv|mp4|webm"
var website = "/"
var uploadfile = "attach/files/"
var uploadvideo = "attach/vedio/"
var uploadimage = "attach/images/"

var configJson []byte // 当客户端请求 /ueditor/go/controller?action=config 返回的json内容

func init() {
	file, err := os.Open("static/editor/ueditor/go/config.json")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer file.Close()
	buf := bytes.NewBuffer(nil)
	buf.ReadFrom(file)

	configJson = buf.Bytes()
}

type UEController struct {
	beego.Controller
}

//方法
func (this *UEController) UEditor() {
	/*	if id, _ := this.GetSession("adminid").(int); id == 0 {
		fmt.Println("Tip:no login")
		this.Ctx.WriteString("{\"state\":\"FAILED\"}")
		this.StopRun()
	}*/

	action := this.Ctx.Input.Query("action")
	switch {
	case action == "config": //获取配置文件
		Config(&this.Controller)
	case action == "uploadimage" || action == "uploadscrawl": //上传图片和涂鸦
		UploadImage(&this.Controller)
	case action == "uploadvideo": //上传视频
		UploadVedio(&this.Controller)
	case action == "uploadfile": //上传文件
		UploadFile(&this.Controller)
	case action == "listimage": //获取图片列表
		ImageManager(&this.Controller)
	case action == "catchimage": //远程抓图
		CatchImage(&this.Controller)
	case action == "listfile": //获取文件列表
		FilesManager(&this.Controller)
	}
}

func Config(this *beego.Controller) {
	//this.ServeJson(configJson)
	this.Ctx.WriteString(string(configJson))
	this.StopRun()
}

func UploadImage(this *beego.Controller) {
	//filename := this.Input().Get("Filename")
	f, h, _ := this.GetFile("upfile")
	filename := h.Filename
	f.Close() //关闭，减少缓存
	ext := filename[strings.LastIndex(filename, ".")+1:]
	//获取扩展名

	if !strings.Contains(allowImageType, ext) {
		fmt.Println(filename)
		this.Ctx.WriteString("{\"state\":\"FAILED\"}")
		this.StopRun()
	}
	newname := strconv.FormatInt(time.Now().Unix(), 10) + "_" + filename
	err := this.SaveToFile("upfile", uploadimage+newname)
	state := "SUCCESS"
	if err != nil {
		fmt.Println(err)
		state = "FAILED"
	}
	url := website + uploadimage + newname
	//this.Ctx.WriteString("{'original':'" + filename + "','url':'" + url + "','title':'" + this.Input().Get("pictitle") + "','state':'" + state + "'}")
	this.Ctx.WriteString("{\"state\": \"" + state + "\", \"url\": \"" + url + "\", \"title\": \"\",\"original\": \"" + filename + "\"}")
	this.StopRun()
}

//上传视频
func UploadVedio(this *beego.Controller) {
	//filename := this.Input().Get("Filename")
	f, h, _ := this.GetFile("upfile")
	filename := h.Filename
	f.Close() //关闭，减少缓存
	index := strings.LastIndex(filename, ".")
	filetype := ""
	if index == -1 {
		this.Ctx.WriteString("{\"state\":\"FAILED\"}")
		this.StopRun()
	}
	filetype = filename[index:]
	ext := filetype[1:]
	if !strings.Contains(allowFileType, ext) {
		this.Ctx.WriteString("{\"state\":\"FAILED\"}")
		this.StopRun()
	}
	newname := strconv.FormatInt(time.Now().Unix(), 10) + "_" + filename
	err := this.SaveToFile("upfile", uploadvideo+newname)
	state := "SUCCESS"
	if err != nil {
		fmt.Println(err)
		state = "FAILED"
	}
	url := website + uploadvideo + newname
	//this.Ctx.WriteString("{'url':'" + url + "','fileType':'" + filetype + "','state':'" + state + "','original':'" + filename + "'}")
	this.Ctx.WriteString("{\"state\": \"" + state + "\", \"url\": \"" + url + "\", \"title\": \"\",\"original\": \"" + filename + "\"}")
	this.StopRun()
}

func UploadFile(this *beego.Controller) {
	//filename := this.Input().Get("Filename")
	f, h, _ := this.GetFile("upfile")
	filename := h.Filename
	f.Close() //关闭，减少缓存

	index := strings.LastIndex(filename, ".")
	filetype := ""
	if index == -1 {
		this.Ctx.WriteString("{\"state\":\"FAILED\"}")
		this.StopRun()
	}
	filetype = filename[index:]
	ext := filetype[1:]
	if !strings.Contains(allowFileType, ext) {
		this.Ctx.WriteString("{\"state\":\"FAILED\"}")
		this.StopRun()
	}
	newname := strconv.FormatInt(time.Now().Unix(), 10) + "_" + filename
	err := this.SaveToFile("upfile", uploadfile+newname)
	state := "SUCCESS"
	if err != nil {
		fmt.Println(err)
		state = "FAILED"
	}
	url := website + uploadfile + newname

	this.Ctx.WriteString("{\"state\": \"" + state + "\", \"url\": \"" + url + "\", \"title\": \"\",\"original\": \"" + filename + "\"}")
	this.StopRun()
}

func ImageManager(this *beego.Controller) {
	strRet := ""
	callbackjson := "{\"state\": \"SUCCESS\",\"list\": ["
	total := 0
	err := filepath.Walk(uploadimage, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		ext := path[strings.LastIndex(path, ".")+1:]
		if strings.Contains(allowImageType, ext) {
			strRet += (path + "ue_separate_ue")
			callbackjson += "{\"url\": \"/" + uploadimage + f.Name() + "\"},"
			total++
			fmt.Println("allow:", path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	callbackjson += "],\"start\": 0,\"total\": " + strconv.Itoa(total) + "}"
	fmt.Println(strRet)
	this.Ctx.WriteString(callbackjson)
	this.StopRun()
}

//远程抓图
func CatchImage(this *beego.Controller) {
	//fmt.Println(this.Ctx.Request.Body)
	urls := this.GetStrings("source[]")
	//fmt.Println(urls)
	callbackjson := "{\"state\": \"SUCCESS\",\"list\": ["
	if len(urls) > 0 {
		for _, v := range urls {
			//去掉最后的!后面部分
			l := v
			//判断扩展名是否合法
			ext := l[strings.LastIndex(l, ".")+1:]

			if strings.Contains(allowImageType, ext) {
				//获取文件名
				filename := l[strings.LastIndex(l, "/")+1:]
				newname := strconv.FormatInt(time.Now().Unix(), 10) + "_" + filename
				res, err := http.Get(l)
				defer res.Body.Close()
				if err != nil {
					callbackjson += "{\"url\": \"\",\"source\": \"" + l + "\",\"state\": \"ERROR\"},"
					fmt.Println("Error:远程抓取失败；", err)
				} else {
					dst, err := os.Create(uploadimage + newname)
					if err != nil {
						callbackjson += "{\"url\": \"\",\"source\": \"" + l + "\",\"state\": \"ERROR\"},"
						fmt.Println("Error:保存失败；", err)
					} else {
						callbackjson += "{\"url\": \"" + uploadimage + newname + "\",\"source\": \"" + l + "\",\"state\": \"SUCCESS\"},"
						io.Copy(dst, res.Body)
					}
				}

			} else {
				callbackjson += "{\"url\": \"\",\"source\": \"" + l + "\",\"state\": \"ERROR\"},"
			}
			//fmt.Println(l)
		}
	}
	callbackjson += "]}"
	this.Ctx.WriteString(callbackjson)
	this.StopRun()
	return
}

//文件管理
func FilesManager(this *beego.Controller) {
	strRet := ""
	callbackjson := "{\"state\": \"SUCCESS\",\"list\": ["
	total := 0
	err := filepath.Walk(uploadfile, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		ext := path[strings.LastIndex(path, ".")+1:]
		if strings.Contains(allowFileType, ext) {
			strRet += (path + "ue_separate_ue")
			callbackjson += "{\"url\": \"/" + uploadfile + f.Name() + "\"},"
			total++
			fmt.Println("allow:", path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	callbackjson += "],\"start\": 0,\"total\": " + strconv.Itoa(total) + "}"
	fmt.Println(strRet)
	this.Ctx.WriteString(callbackjson)
	this.StopRun()
}
