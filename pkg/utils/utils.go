/*
@Time : 2021/2/8 11:05
@Author : wangtao
@File : utils
*/
package utils

import (
	"io/ioutil"
	"os"
)

func LoadFile(path string) (data string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	byteTmp, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(byteTmp[:]), nil
}
