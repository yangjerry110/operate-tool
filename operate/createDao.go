/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-13 18:47:16
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-15 17:18:23
 * @Description: create dao
 */
package operate

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

type (
	CreateDao struct{}

	CreateDaoInfo struct{}
)

func (c *CreateDao) Action(operateParam string) {
	/**
	 * @step
	 * @解析参数
	 **/
	appParams := strings.Split(operateParam, ".")

	/**
	 * @step
	 * @检查参数
	 **/
	err := c.CheckAppParams(appParams)
	if err != nil {
		fmt.Printf("CreateDao Err : %v", err.Error())
		fmt.Print("\r\n")
		return
	}

	/**
	 * @step
	 * @定义参数
	 **/
	projectName := appParams[0]
	daoName := appParams[1]
	modelName := appParams[2]
	authorName := appParams[3]

	if authorName == "" {
		authorName = "Jerry.Yang"
	}

	/**
	 * @step
	 * @获取fileContent
	 **/
	fileContent, err := c.CreateContent(daoName, modelName, projectName, authorName)
	if err != nil {
		fmt.Printf("Err : %v", err.Error())
		fmt.Print("\r\n")
		return
	}

	/**
	 * @step
	 * @创建文件
	 **/
	err = c.CreateFile(daoName, fileContent)
	if err != nil {
		fmt.Printf("Err : %v", err.Error())
		fmt.Print("\r\n")
		return
	}
	return
}

/**
 * @description: CreateContent
 * @param {string} daoName
 * @param {string} modelName
 * @author: Jerry.Yang
 * @date: 2022-09-14 16:29:08
 * @return {*}
 */
func (c *CreateDao) CreateContent(daoName string, modelName string, projectName string, authorName string) (string, error) {
	fileContent := fmt.Sprintf("package dao")
	fileContent += "\r\n"

	/**
	 * @step
	 * @import
	 **/
	fileContent += "import ( \r\n"
	fileContent += fmt.Sprintf("\"%s/internal\"\r\n", projectName)
	fileContent += fmt.Sprintf("\"%s/model\"\r\n", projectName)
	fileContent += ")\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @interface
	 **/
	fileContent += fmt.Sprintf("type %sInterface interface {\r\n", strings.Title(daoName))
	fileContent += fmt.Sprintf("GetInfo(%s *model.%s) (*model.%s, error)\r\n", modelName, strings.Title(modelName), strings.Title(modelName))
	fileContent += fmt.Sprintf("GetList(%s *model.%s) ([]*model.%s, error) \r\n", modelName, strings.Title(modelName), strings.Title(modelName))
	fileContent += fmt.Sprintf("Save(%sModel *model.%s) (bool, error)\r\n", modelName, strings.Title(modelName))
	fileContent += fmt.Sprintf("Delete(%s *model.%s) (bool, error)\r\n", modelName, strings.Title(modelName))
	fileContent += "}\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @struct
	 **/
	fileContent += fmt.Sprintf("type %s struct {}\r\n", strings.Title(daoName))
	fileContent += "\r\n"

	/**
	 * @step
	 * @getInfoById
	 **/
	infoContent := c.CreateDaoInfo(daoName, modelName, authorName)
	fileContent += infoContent
	fileContent += "\r\n"

	/**
	 * @step
	 * @getList
	 **/
	listContent := c.CreateDaoList(daoName, modelName, authorName)
	fileContent += listContent
	fileContent += "\r\n"

	/**
	 * @step
	 * @save
	 **/
	saveContent := c.CreateDaoSave(daoName, modelName, authorName)
	fileContent += saveContent
	fileContent += "\r\n"

	/**
	 * @step
	 * @delete
	 **/
	deleteContent := c.CreateDaoDeleted(daoName, modelName, authorName)
	fileContent += deleteContent
	fileContent += "\r\n"

	return fileContent, nil
}

/**
 * @description: CreateDaoInfo
 * @param {string} daoName
 * @param {string} modelName
 * @author: Jerry.Yang
 * @date: 2022-09-14 16:35:31
 * @return {*}
 */
func (c *CreateDao) CreateDaoInfo(daoName string, modelName string, authorName string) string {
	fileContent := fmt.Sprintf("/** \r\n * @description: GetInfo \r\n * @param {*model.%s} %s \r\n * @author: %s \r\n * @date: %s \r\n * @return {*} \r\n */ \r\n", strings.Title(modelName), modelName, authorName, time.Now().Format("2006-01-02 15:04:05"))
	fileContent += fmt.Sprintf("func (%s *%s) GetInfo(%s *model.%s) (*model.%s, error) {\r\n", daoName[0:1], strings.Title(daoName), modelName, strings.Title(modelName), strings.Title(modelName))
	fileContent += fmt.Sprintf("result := model.%s{}\r\n", strings.Title(modelName))
	fileContent += fmt.Sprintf("if err := internal.DbClient().Where(%s).First(&result).Error; err != nil { \r\n", modelName)
	fileContent += "internal.LoggorError(err)\r\n"
	fileContent += "return nil, err\r\n"
	fileContent += "}\r\n"
	fileContent += "return &result, nil\r\n"
	fileContent += "}\r\n"
	fileContent += "\r\n"
	return fileContent
}

/**
 * @description: CreateDaoList
 * @param {string} daoName
 * @param {string} modelName
 * @author: Jerry.Yang
 * @date: 2022-09-15 14:37:55
 * @return {*}
 */
func (c *CreateDao) CreateDaoList(daoName string, modelName string, authorName string) string {
	fileContent := fmt.Sprintf("/** \r\n * @description: GetList \r\n * @param {*model.%s} %s \r\n * @author: %s \r\n * @date: %s \r\n * @return {*} \r\n */ \r\n", strings.Title(modelName), modelName, authorName, time.Now().Format("2006-01-02 15:04:05"))
	fileContent += fmt.Sprintf("func (%s *%s) GetList(%s *model.%s) ([]*model.%s, error) {\r\n", daoName[0:1], strings.Title(daoName), modelName, strings.Title(modelName), strings.Title(modelName))
	fileContent += fmt.Sprintf("result := make([]*model.%s, 0)\r\n", strings.Title(modelName))
	fileContent += fmt.Sprintf("if err := internal.DbClient().Where(%s).Find(&result).Error; err != nil {\r\n", modelName)
	fileContent += "internal.LoggorError(err)\r\n"
	fileContent += "return nil, err\r\n"
	fileContent += "}\r\n"
	fileContent += "return result, nil\r\n"
	fileContent += "}\r\n"
	fileContent += "\r\n"
	return fileContent
}

/**
 * @description: CreateDaoSave
 * @param {string} daoName
 * @param {string} modelName
 * @author: Jerry.Yang
 * @date: 2022-09-15 14:38:46
 * @return {*}
 */
func (c *CreateDao) CreateDaoSave(daoName string, modelName string, authorName string) string {
	fileContent := fmt.Sprintf("/** \r\n * @description: Save \r\n * @param {*model.%s} %s \r\n * @author: %s \r\n * @date: %s \r\n * @return {*} \r\n */ \r\n", modelName, strings.Title(modelName), authorName, time.Now().Format("2006-01-02 15:04:05"))
	fileContent += fmt.Sprintf("func (%s *%s) Save(%s *model.%s) (bool, error) {\r\n", daoName[0:1], strings.Title(daoName), modelName, strings.Title(modelName))
	fileContent += "result := false\r\n"
	fileContent += fmt.Sprintf("if err := internal.DbClient().Save(%s).Error; err != nil {\r\n", modelName)
	fileContent += "internal.LoggorError(err)\r\n"
	fileContent += "return result, err\r\n"
	fileContent += "}\r\n"
	fileContent += "result = true\r\n"
	fileContent += "return result, nil\r\n"
	fileContent += "}\r\n"
	fileContent += "\r\n"
	return fileContent
}

/**
 * @description: CreateDaoDeleted
 * @param {string} daoName
 * @param {string} modelName
 * @author: Jerry.Yang
 * @date: 2022-09-15 14:42:04
 * @return {*}
 */
func (c *CreateDao) CreateDaoDeleted(daoName string, modelName string, authorName string) string {
	fileContent := fmt.Sprintf("/** \r\n * @description: Delete \r\n * @param {*model.%s} %s \r\n * @author: %s \r\n * @date: %s \r\n * @return {*} \r\n */ \r\n", strings.Title(modelName), modelName, authorName, time.Now().Format("2006-01-02 15:04:05"))
	fileContent += fmt.Sprintf("func (%s *%s) Delete(%s *model.%s) (bool, error) {\r\n", daoName[0:1], strings.Title(daoName), modelName, strings.Title(modelName))
	fileContent += fmt.Sprintf("result := false\r\n")
	fileContent += fmt.Sprintf("if err := internal.DbClient().Model(&model.%s{}).Where(%s).Update(\"is_deleted\", model.IS_DELETED).Error; err != nil {\r\n", strings.Title(modelName), modelName)
	fileContent += "internal.LoggorError(err)\r\n"
	fileContent += "return result, err\r\n"
	fileContent += "}\r\n"
	fileContent += "result = true\r\n"
	fileContent += "return result, nil\r\n"
	fileContent += "}\r\n"
	fileContent += "\r\n"
	return fileContent
}

/**
 * @description: CreateFile
 * @param {string} daoName
 * @param {string} fileContent
 * @author: Jerry.Yang
 * @date: 2022-09-14 16:25:45
 * @return {*}
 */
func (c *CreateDao) CreateFile(daoName string, fileContent string) error {

	/**
	 * @step
	 * @获取当前目录
	 **/
	dir, err := os.Getwd()
	if err != nil {
		return errors.New("获取目录错误!请重试!")
	}

	/**
	 * @step
	 * @定义文件名称
	 **/
	inputFile := fmt.Sprintf("%s/%s.go", dir, daoName)

	/**
	 * @step
	 * @判断文件是否存在
	 **/
	_, err = os.Stat(inputFile)
	if err == nil {
		return errors.New(fmt.Sprintf("文件存在! filename = %s", inputFile))
	} else {
		if os.IsExist(err) {
			return errors.New(fmt.Sprintf("文件存在! filename = %s", inputFile))
		}
	}

	/**
	 * @step
	 * @定义osFile
	 **/
	var osFile *os.File

	/**
	 * @step
	 * @创建input文件
	 **/
	osFile, err = os.Create(inputFile)
	if err != nil {
		return errors.New(fmt.Sprintf("创建文件错误! filename = %s", inputFile))
	}
	defer osFile.Close()

	/**
	 * @step
	 * @写入内容
	 **/
	_, err = io.WriteString(osFile, fileContent)
	if err != nil {
		return errors.New(fmt.Sprintf("写入文件错误! filename = %s", inputFile))
	}

	fmt.Printf("%s 创建成功!", inputFile)
	fmt.Print("\r\n")

	/**
	 * @step
	 * @格式化代码
	 **/
	cmd := exec.Command("gofmt", "-w", inputFile)
	//cmd := exec.Command("ls", "-al", "/Users/admin/go/src/go.test.app/test")
	if err := cmd.Run(); err != nil {
		fmt.Printf("格式化错误! Err : %v", err.Error())
		fmt.Print("\r\n")
	}
	return nil
}

/**
 * @description: CheckAppParams
 * @param {[]string} appParams
 * @author: Jerry.Yang
 * @date: 2022-09-14 16:11:57
 * @return {*}
 */
func (c *CreateDao) CheckAppParams(appParams []string) error {

	/**
	 * @step
	 * @判断参数
	 **/
	if len(appParams) == 0 {
		return errors.New("projectName缺失!")
	}

	if len(appParams) == 1 {
		return errors.New("daoName缺失!")
	}

	if len(appParams) == 2 {
		return errors.New("modelName缺失!")
	}
	return nil
}
