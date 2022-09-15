/*
 * @Author: yangjie04@qutoutiao.net
 * @Date: 2022-09-07 14:30:23
 * @LastEditors: yangjie04@qutoutiao.net
 * @LastEditTime: 2022-09-15 10:30:21
 * @Description: baseOperate
 */
package operate

import (
	"errors"
	"fmt"
)

type BaseOperateInterface interface{}

type BaseOperate struct{}

/**
 * @description: ExecOperate
 * @param {string} actOperate
 * @param {string} operateParam
 * @author: yangjie04@qutoutiao.net
 * @date: 2022-09-07 14:32:09
 * @return {*}
 */
func (b *BaseOperate) ExecOperate(actOperate string, operateParam string) error {

	switch actOperate {
	case "createApp":
		actionObj := CreatedApp{}
		actionObj.Action(operateParam)
	case "createDao":
		actionObj := CreateDao{}
		actionObj.Action(operateParam)
	default:
		fmt.Printf("没有符合的操作!")
		fmt.Print("\r\n")
		return errors.New("没有符合的操作!")
	}

	return nil
}
