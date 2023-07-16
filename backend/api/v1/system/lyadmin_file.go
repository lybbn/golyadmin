package system

import (
	"path"

	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/utils"
	"gitee.com/lybbn/golyadmin/utils/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
// savePath = "media/uploadfile/"
)

var (
	//允许上传文件后缀白名单
	whiteFileExtMap = map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	allowMaxFileSize = 30 //允许最大文件大小，单位M
)

type FileResponse struct {
	Size int64  `json:"size"` //文件大小
	Path string `json:"path"` //文件路径(相对路径)
	Url  string `json:"url"`  //url地址
	Name string `json:"name"` //文件名
	Type string `json:"type"` //文件类型
}

type FileApi struct {
}

// @Tags      Menu
// @Summary   文件上传
// @Security  ApiKeyAuth
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 2000 {string} string	"{"code": 2000,"data":"fileurl", "msg": "上传成功"}"
// @Success 2000 {string} string	"{"code": 4000, "msg": "上传失败"}"
// @Router   /system/file/uploadFile [post]
func (f *FileApi) UploadFileLocal(c *gin.Context) {
	//限制 HTTP 请求中读取的最大字节数。这个函数会返回一个新的 Reader 对象，该对象会在读取请求的正文时自动检查字节数，如果超过指定的最大字节数，则会自动停止读取，返回错误
	// c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, int64(allowMaxFileSize<<20))
	urlPrefix := global.GL_CONFIG.System.UrlPrefix
	file, err := c.FormFile("file")
	if err != nil {
		global.GL_LOG.Error("上传的文件保存失败:", zap.Any(" error:", err))
		response.ErrorResponse("上传的上传失败", c)
		return
	}
	nday := utils.GetNowDayStr()
	dir := global.GL_CONFIG.System.UploadDir + nday
	err = utils.CreateDir(dir)
	if err != nil {
		response.ErrorResponse("初始化文件路径失败", c)
		return
	}
	ext := utils.GetFileExt(file.Filename)
	_, wte := whiteFileExtMap[ext]
	if !wte {
		response.ErrorResponse("上传的文件类型不合法", c)
		return
	}
	fileSize := file.Size
	if fileSize > int64(allowMaxFileSize*1024*1024) {
		response.ErrorResponse("上传的文件超过"+utils.FormatInt2String(allowMaxFileSize)+"M限制", c)
		return
	}
	srcf, err := file.Open()
	if err != nil {
		response.ErrorResponse("上传的文件错误", c)
		return
	}
	defer srcf.Close()
	fileName := utils.GenerateUniqueUUIDFileName() + ext
	// dstPathFileName := dir + fileName
	dstPathFile := path.Join(dir, fileName)
	err1 := c.SaveUploadedFile(file, dstPathFile)
	if err1 != nil {
		global.GL_LOG.Error("上传的文件保存失败:", zap.Any(" error:", err1))
		response.ErrorResponse("上传的文件保存失败", c)
		return
	}
	fileurl := urlPrefix + "/" + dstPathFile
	response.SuccessResponse(fileurl, "上传成功", c)
}
