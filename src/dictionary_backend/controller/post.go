package controller

import (
	"dictionary_backend/configure"
	"github.com/DataDog/go-python3"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"os"
)
import "fmt"

func PostInfo(c *gin.Context) {
	language := c.Param("language")
	str := c.Param("str")
	sign := c.Param("sign")

	file, _ := c.FormFile("file")

	path := configure.Path + "/" + language + "/" + str + "/" + sign + "/"
	//如果没有path文件目录就创建一个
	if _, err := os.Stat(path); err != nil {
		if !os.IsExist(err) {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				fmt.Println("mkdir failed! ", err)
				return
			}
		}
	}

	err1 := c.SaveUploadedFile(file, path+file.Filename)
	if err1 != nil {
		c.String(http.StatusBadRequest, err1.Error())
		fmt.Println("save file failed! ", err1)
		return
	}

	PyTest(path, file.Filename)

	database, err2 := sqlx.Connect("mysql", configure.ConnectSentence)
	if err2 != nil {
		fmt.Println("open mysql failed,", err2)
		return
	}

	configure.Db = database

	r, err3 := configure.Db.Exec("insert into "+configure.TableName+" (language, str, sign, location)values(?, ?, ?, ?)", language, str, sign, path+file.Filename)
	if err3 != nil {
		fmt.Println("exec failed, ", err3)
		return
	}
	id, err4 := r.LastInsertId()
	if err4 != nil {
		fmt.Println("exec failed, ", err4)
		return
	}

	fmt.Println("insert success! ", id)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func PyTest(filePath, fileName string) {
	python3.Py_Initialize()
	if !python3.Py_IsInitialized() {
		fmt.Println("Error initializing the python interpreter")
		os.Exit(1)
	}



	v := ImportModule("/home/lighthouse/python", "silence")
	if v == nil {
		fmt.Println("Import Module failed!")
		return
	}

	silenceRemove := v.GetAttrString("silenceremove")
	if silenceRemove == nil {
		fmt.Println("silence Remove is nil")
	}

	bArgs := python3.PyTuple_New(2)
	python3.PyTuple_SetItem(bArgs, 0, python3.PyUnicode_FromString(filePath))
	python3.PyTuple_SetItem(bArgs, 1, python3.PyUnicode_FromString(fileName))

	silenceRemove.Call(bArgs, python3.Py_None)

	python3.Py_Finalize()
}

func ImportModule(dir, name string) *python3.PyObject {
	sysModule := python3.PyImport_ImportModule("sys")
	path := sysModule.GetAttrString("path")

	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString(dir))

	return python3.PyImport_ImportModule(name)
}
