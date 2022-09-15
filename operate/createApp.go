/*
 * @Author: yangjie04@qutoutiao.net
 * @Date: 2022-09-07 14:34:01
 * @LastEditors: yangjie04@qutoutiao.net
 * @LastEditTime: 2022-09-14 16:08:30
 * @Description: create app
 */
package operate

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type (
	CreateAppInterface interface{}

	CreatedApp struct{}

	CreateAppInputVO struct {
		FileName string
	}

	CreateAppOutputVO struct {
		FileName string
	}

	CreateAppRoute struct {
		FileName string
	}

	CreateAppService struct {
		FileName string
	}
)

/**
 * @description: Action
 * @param {string} operateParam
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-07 14:37:39
 * @return {*}
 */
func (c *CreatedApp) Action(operateParam string) {

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
		fmt.Printf("createApp Err : %v", err.Error())
		fmt.Print("\r\n")
		return
	}

	/**
	 * @step
	 * @定义参数
	 **/
	projectName := appParams[0] // project
	appName := appParams[1]     // appName
	method := appParams[2]      // method

	/**
	 * @step
	 * @获取当前目录
	 **/
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取目录错误!请重试!")
		fmt.Print("\r\n")
		return
	}

	/**
	 * @step
	 * @判断目录是否存在
	 **/
	needCreateAppDir := fmt.Sprintf("%s/%s", dir, appName)
	_, err = os.Stat(needCreateAppDir)
	if os.IsNotExist(err) {

		/**
		 * @step
		 * @创建目录
		 **/
		err = os.Mkdir(needCreateAppDir, 0777)
		if err != nil {
			fmt.Printf("创建App目录失败!请重试!")
			fmt.Print("\r\n")
			return
		}
	}

	/**
	 * @step
	 * @获取fileContent inputVO
	 **/
	createAppInputVO := CreateAppInputVO{FileName: "InputVO"}
	createAppInputVOFileContent := createAppInputVO.CreateContent(needCreateAppDir, appName, createAppInputVO.FileName)
	err = c.CreateFile(needCreateAppDir, appName, createAppInputVO.FileName, createAppInputVOFileContent)
	if err != nil {
		fmt.Printf("createApp Err : %v", err.Error())
		fmt.Print("\r\n")
	}

	/**
	 * @step
	 * @获取fileContent OutputVO
	 **/
	createAppOutputVO := CreateAppOutputVO{FileName: "OutputVO"}
	createAppOutputVOFileContent := createAppOutputVO.CreateContent(needCreateAppDir, projectName, appName, createAppOutputVO.FileName)
	err = c.CreateFile(needCreateAppDir, appName, createAppOutputVO.FileName, createAppOutputVOFileContent)
	if err != nil {
		fmt.Printf("createApp Err : %v", err.Error())
		fmt.Print("\r\n")
	}

	/**
	 * @step
	 * @获取fileContent route
	 **/
	createAppRoute := CreateAppRoute{FileName: "Route"}
	createAppRouteFileContent := createAppRoute.CreateContent(needCreateAppDir, projectName, appName, createAppRoute.FileName, method)
	err = c.CreateFile(needCreateAppDir, appName, createAppRoute.FileName, createAppRouteFileContent)
	if err != nil {
		fmt.Printf("createApp Err : %v", err.Error())
		fmt.Print("\r\n")
	}

	/**
	 * @step
	 * @获取fileContent service
	 **/
	createAppService := CreateAppService{FileName: "Service"}
	CreateAppServiceFileContent := createAppService.CreateContent(needCreateAppDir, projectName, appName, createAppService.FileName, method)
	err = c.CreateFile(needCreateAppDir, appName, createAppService.FileName, CreateAppServiceFileContent)
	if err != nil {
		fmt.Printf("createApp Err : %v", err.Error())
		fmt.Print("\r\n")
	}

	return
}

/**
 * @description: CreateAppInput
 * @param {string} appDir
 * @param {string} appName
 * @param {string} fileName
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-07 14:52:41
 * @return {*}
 */
func (c *CreateAppInputVO) CreateContent(appDir string, appName string, fileName string) string {

	/**
	 * @step
	 * @定义写入内容
	 **/
	fileContent := fmt.Sprintf("package %s", appName)
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @写入interface
	 **/
	interfaceContent := fmt.Sprintf("type %s%sInterface interface {}", strings.Title(appName), strings.Title(fileName))
	fileContent += interfaceContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @写入struct
	 **/
	structContent := fmt.Sprintf("type %s%s struct{}", strings.Title(appName), strings.Title(fileName))
	fileContent += structContent
	return fileContent
}

/**
 * @description: CreateContent
 * @param {string} appDir
 * @param {string} projectName
 * @param {string} appName
 * @param {string} fileName
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-07 16:04:00
 * @return {*}
 */
func (c *CreateAppOutputVO) CreateContent(appDir string, projectName string, appName string, fileName string) string {
	fileContent := fmt.Sprintf("package %s", appName)
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @import
	 **/
	importContent := fmt.Sprintf("import (\r\n")
	importContent += "	\"github.com/gin-gonic/gin\"\r\n"
	importContent += fmt.Sprintf("	\"%s/internal\"\r\n", projectName)
	importContent += ")\r\n"
	importContent += "\r\n"
	fileContent += importContent

	/**
	 * @step
	 * @interface
	 **/
	interfaceContent := fmt.Sprintf("type %s%sInterface interface {}", strings.Title(appName), strings.Title(fileName))
	fileContent += interfaceContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @structContent
	 **/
	structContent := fmt.Sprintf("type %s%s struct{\r\n	HttpStatus int\r\n RetCode int\r\n RetMsg string\r\n RetResult interface{}\r\n}", strings.Title(appName), strings.Title(fileName))
	fileContent += structContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @RenderOutputVOSimple
	 **/
	renderOutputVOSimpleContent := fmt.Sprintf("func RenderOutputVOSimple(ctx *gin.Context,retCode int,retMsg string,httpStatus ...int) error {\r\n")
	renderOutputVOSimpleContent += fmt.Sprintf("%s%s := %s%s{ \r\n", strings.Title(appName), strings.Title(fileName), strings.Title(appName), strings.Title(fileName))
	renderOutputVOSimpleContent += "RetCode : retCode,\r\n"
	renderOutputVOSimpleContent += "RetMsg : retMsg,\r\n"
	renderOutputVOSimpleContent += "} \r\n"
	renderOutputVOSimpleContent += "\r\n"
	renderOutputVOSimpleContent += "if len(httpStatus) == 0 {\r\n"
	renderOutputVOSimpleContent += fmt.Sprintf("%s%s.RenderOutputVO(ctx)\r\n", strings.Title(appName), strings.Title(fileName))
	renderOutputVOSimpleContent += "return nil\r\n"
	renderOutputVOSimpleContent += "}\r\n"
	renderOutputVOSimpleContent += "\r\n"
	renderOutputVOSimpleContent += fmt.Sprintf("%s%s.HttpStatus = httpStatus[0]\r\n", strings.Title(appName), strings.Title(fileName))
	renderOutputVOSimpleContent += fmt.Sprintf("%s%s.RenderOutputVO(ctx)\r\n", strings.Title(appName), strings.Title(fileName))
	renderOutputVOSimpleContent += "return nil\r\n"
	renderOutputVOSimpleContent += "}\r\n"
	fileContent += renderOutputVOSimpleContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @RenderOutputVO
	 **/
	renderOutputVOContent := fmt.Sprintf("func (%s%s *%s%s) RenderOutputVO(ctx *gin.Context) error {\r\n", appName, fileName, strings.Title(appName), strings.Title(fileName))
	renderOutputVOContent += fmt.Sprintf("internalOutput := internal.Output{}\r\n")
	renderOutputVOContent += fmt.Sprintf("internalOutput.OutputFunc(ctx, %s%s.RetCode, %s%s.RetMsg, %s%s.RetResult, %s%s.HttpStatus)\r\n", appName, fileName, appName, fileName, appName, fileName, appName, fileName)
	renderOutputVOContent += "return nil\r\n"
	renderOutputVOContent += "}\r\n"
	fileContent += renderOutputVOContent
	return fileContent
}

/**
 * @description: CreateContent
 * @param {string} appDir
 * @param {string} projectName
 * @param {string} appName
 * @param {string} fileName
 * @param {string} method
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-07 16:42:13
 * @return {*}
 */
func (c *CreateAppRoute) CreateContent(appDir string, projectName string, appName string, fileName string, method string) string {
	fileContent := fmt.Sprintf("package %s", appName)
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @import
	 **/
	importContent := fmt.Sprintf("import (\r\n")
	importContent += "\"github.com/gin-gonic/gin\"\r\n"
	importContent += "\"net/http\"\r\n"
	importContent += fmt.Sprintf("\"%s/config\"\r\n", projectName)
	importContent += ")\r\n"
	importContent += "\r\n"
	fileContent += importContent

	/**
	 * @step
	 * @interface
	 **/
	interfaceContent := fmt.Sprintf("type %s%sInterface interface {}", strings.Title(appName), strings.Title(fileName))
	fileContent += interfaceContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @structContent
	 **/
	structContent := fmt.Sprintf("type %s%s struct{}", strings.Title(appName), strings.Title(fileName))
	fileContent += structContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @routeFunc
	 **/
	routeFuncContent := fmt.Sprintf("func %sRouteFunc(ctx *gin.Context) {\r\n", strings.Title(appName))
	routeFuncContent += fmt.Sprintf("%sInputVO := %sInputVO{} \r\n", appName, strings.Title(appName))

	/**
	 * @step
	 * @判断什么请求方式
	 **/
	if method == "GET" {
		routeFuncContent += fmt.Sprintf("if err := ctx.ShouldBindQuery(&%sInputVO); err != nil {\r\n", appName)
	}

	if method == "POST" {
		routeFuncContent += fmt.Sprintf("if err := ctx.ShouldBind(%sInputVO); err != nil {\r\n", appName)
	}
	routeFuncContent += fmt.Sprintf("RenderOutputVOSimple(ctx, config.COMMON_ERROR, err.Error())\r\n")
	routeFuncContent += "return \r\n"
	routeFuncContent += "}\r\n"
	routeFuncContent += "\r\n"

	routeFuncContent += fmt.Sprintf("err := %sInputVO.%sServiceFunc(ctx)\r\n", appName, strings.Title(appName))
	routeFuncContent += "if err != nil {\r\n"
	routeFuncContent += "return"
	routeFuncContent += "}\r\n"
	routeFuncContent += "\r\n"

	routeFuncContent += fmt.Sprintf("%sOutputVO := %sOutputVO{} \r\n", appName, strings.Title(appName))
	routeFuncContent += fmt.Sprintf("%sOutputVO.HttpStatus = http.StatusOK \r\n", appName)
	routeFuncContent += fmt.Sprintf("%sOutputVO.RetCode = config.NO_ERROR \r\n", appName)
	routeFuncContent += fmt.Sprintf("%sOutputVO.RetResult = true \r\n", appName)
	routeFuncContent += fmt.Sprintf("%sOutputVO.RenderOutputVO(ctx) \r\n", appName)
	routeFuncContent += "return \r\n"
	routeFuncContent += "}\r\n"
	fileContent += routeFuncContent
	return fileContent
}

/**
 * @description: CreateContent
 * @param {string} appDir
 * @param {string} projectName
 * @param {string} appName
 * @param {string} fileName
 * @param {string} method
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-07 17:16:40
 * @return {*}
 */
func (c *CreateAppService) CreateContent(appDir string, projectName string, appName string, fileName string, method string) string {
	fileContent := fmt.Sprintf("package %s", appName)
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @import
	 **/
	importContent := fmt.Sprintf("import (\r\n")
	importContent += "\"github.com/gin-gonic/gin\"\r\n"
	importContent += ")\r\n"
	importContent += "\r\n"
	fileContent += importContent

	/**
	 * @step
	 * @interface
	 **/
	interfaceContent := fmt.Sprintf("type %s%sInterface interface {}", strings.Title(appName), strings.Title(fileName))
	fileContent += interfaceContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	/**
	 * @step
	 * @structContent
	 **/
	structContent := fmt.Sprintf("type %s%s struct{}", strings.Title(appName), strings.Title(fileName))
	fileContent += structContent
	fileContent += "\r\n"
	fileContent += "\r\n"

	serviceContent := fmt.Sprintf("func (vo *%sInputVO) %sServiceFunc(ctx *gin.Context) error {\r\n", strings.Title(appName), strings.Title(appName))
	serviceContent += "return nil \r\n"
	serviceContent += "}\r\n"
	fileContent += serviceContent
	return fileContent
}

/**
 * @description: CreateFile
 * @param {string} appDir
 * @param {string} appName
 * @param {string} fileName
 * @param {string} fileContent
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-07 15:11:15
 * @return {*}
 */
func (c *CreatedApp) CreateFile(appDir string, appName string, fileName string, fileContent string) error {
	/**
	 * @step
	 * @定义文件名称
	 **/
	inputFile := fmt.Sprintf("%s/%s%s.go", appDir, appName, fileName)

	/**
	 * @step
	 * @判断文件是否存在
	 **/
	_, err := os.Stat(inputFile)
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
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-07 14:47:04
 * @return {*}
 */
func (c *CreatedApp) CheckAppParams(appParams []string) error {

	/**
	 * @step
	 * @判断参数
	 **/
	if len(appParams) == 0 {
		return errors.New("daoName缺失!")
	}

	if len(appParams) == 1 {
		return errors.New("modelName缺失!")
	}
	return nil
}
