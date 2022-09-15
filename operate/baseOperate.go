/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-07 14:30:23
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-15 17:17:57
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
 * @author: Jerry.Yang
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
