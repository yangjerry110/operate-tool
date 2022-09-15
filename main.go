/*
 * @Author: yangjie04@qutoutiao.net
 * @Date: 2022-09-07 14:20:30
 * @LastEditors: yangjie04@qutoutiao.net
 * @LastEditTime: 2022-09-07 16:07:24
 * @Description: main
 */
package main

import (
	"fmt"
	"os"

	"go.test.app/operate"
)

func main() {

	/**
	 * @step
	 * @判断args
	 **/
	if len(os.Args) < 3 {
		fmt.Printf("命令行参数错误!,end")
		fmt.Print("\r\n")
		return
	}

	/**
	 * @step
	 * @获取传入的参数
	 **/
	actOperate := os.Args[1]
	if actOperate == "" {
		fmt.Printf("不需要操作,end")
		fmt.Print("\r\n")
		return
	}

	/**
	 * @step
	 * @获取操作参数
	 **/
	operateParam := os.Args[2]
	if operateParam == "" {
		fmt.Printf("不需要操作,end")
		fmt.Print("\r\n")
		return
	}

	/**
	 * @step
	 * @调用对应的action
	 **/
	operateObj := operate.BaseOperate{}
	operateObj.ExecOperate(actOperate, operateParam)

	fmt.Printf("执行成功!end")
	fmt.Print("\r\n")
	return
}
