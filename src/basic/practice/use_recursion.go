/**
 * 使用递归示例
 */

package practice

import (
	"fmt"
	"strconv"
)

var menu map[int]string

func init() {
	menu = make(map[int]string)
}

func subMenu(parentid int, menu map[int]string) {
	if parentid == 0 {
		parentid++
		subMenu(parentid, menu)
	} else if parentid == 4 {
		menu[parentid] = "menu" + strconv.Itoa(parentid*parentid)
	} else {
		menu[parentid] = "menu" + strconv.Itoa(parentid*parentid)
		parentid++
		subMenu(parentid, menu)
	}
}

func GetSubMenu() {
	subMenu(0, menu)

	fmt.Printf("%+v \n", menu)
}
