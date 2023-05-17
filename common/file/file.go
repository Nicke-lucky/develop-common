package file

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

// 读取json文件
func ReadJsonFile(FilePath, FileName string) []byte {
	fr, err := os.Open(FilePath + "/" + FileName)
	if err != nil {
		log.Errorln("Open file failed ERROR:", fmt.Sprint(err.Error()))
		return nil
	}
	defer func() {
		_ = fr.Close()
	}()
	byteValue, _ := ioutil.ReadAll(fr)
	return byteValue
}

// 写入json文件
func WriteJsonFile(FilePath, FileName string, data []byte) {

	// 创建文件
	fw, err := os.Create(FilePath + "/" + FileName)
	if err != nil {
		log.Errorln("WriteJsonFile Create file failed", err.Error())
		return
	}
	defer func() {
		_ = fw.Close()
	}()

	// 带JSON缩进格式写文件
	//data, err := json.MarshalIndent(Info, "  ", "  ")
	//if err != nil {
	//	log.Errorln("Encoder failed", err.Error())
	//	return
	//}

	_, fwerr := fw.Write(data)
	if fwerr != nil {
		log.Errorln("WriteJsonFile fw.Write failed", fwerr.Error())
	}
}

// 创建xml文件
func CreateFile(path, name string, output []byte) error {
	//dir, _ := os.Getwd()
	// check
	if _, err := os.Stat(path); err != nil {
		err = os.MkdirAll(path, 0711)
		if err != nil {
			log.Errorln("Error creating directory,err:", err)
			return err
		}
	}
	fw, f_werr := os.Create(path + name)
	defer func() {
		_ = fw.Close()
	}()
	if f_werr != nil {
		log.Errorln("os.Create error:", f_werr)
		return f_werr
	}
	_, ferr := fw.Write((output))
	if ferr != nil {
		return ferr
	}
	return nil
}
