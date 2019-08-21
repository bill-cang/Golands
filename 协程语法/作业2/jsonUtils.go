package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func writeJsonToFile(downMap map[string]Idiom) (err error) {
	file, _ := os.OpenFile(localPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0754)
	defer file.Close()
	encoder := json.NewEncoder(file)
	fmt.Println(downMap)
	err = encoder.Encode(downMap)
	if err != nil {
		fmt.Println("持久化到本地失败！err：",err)
		return
	}

	fmt.Println("持久化到本地成功！")
	return nil

}
