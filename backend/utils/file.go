package utils

import (
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"

	"gitee.com/lybbn/golyadmin/global"
	"go.uber.org/zap"
)

// GetFileSize 获取文件大小
func GetFileSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)
	return len(content), err
}

// GetFileExt 获取文件后缀
func GetFileExt(fileName string) string {
	return path.Ext(fileName)
}

// GetType 获取文件类型
func GetFileContentType(out *os.File) (string, error) {

	buff := make([]byte, 512)

	_, err := out.Read(buff)

	if err != nil {
		global.GL_LOG.Error("获取文件类型失败：", zap.Any(" error:", err))
		return "", err
	}

	filetype := http.DetectContentType(buff)
	return filetype, nil
}

// GetType 获取文件类型(根据文件路径)
func GetFileContentTypeByPath(p string) (string, error) {
	file, err := os.Open(p)

	if err != nil {
		global.GL_LOG.Error("获取文件类型失败,打开文件失败："+p, zap.Any(" error:", err))
		return "", err
	}
	defer file.Close()

	return GetFileContentType(file)
}

// 文件目录是否存在
func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 批量创建文件夹
func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.GL_LOG.Debug("create directory" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				global.GL_LOG.Error("create directory"+v, zap.Any(" error:", err))
				return err
			}
		}
	}
	return err
}
