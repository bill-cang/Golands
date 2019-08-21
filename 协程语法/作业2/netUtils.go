package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)
//mmp,还以为自己写错了，每次第五次调用都失败！排查一个小时，json===： map[showapi_res_error:您的调用过于频繁!1 showapi_res_code:-1009 showapi_res_body:map[]]
func RequestHttpGet(url string) (bytes []byte) {
	time.Sleep(1*time.Second)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("网络请求失败！err：", err)
		return
	}
	//挂起关闭
	defer response.Body.Close()
	//解析网络数据形成字节
	bytes, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("数据解析失败，err：", err)
		return
	}
	return
}

func GetKeyIdioms(key string) (idioms []string, err error) {
	url := idiomUrl + "&keyword=" + key + "&page=1&rows=2"
	//发起get请求
	bytes := RequestHttpGet(url)
	//将网络数据反序列化保存到变量
	tempMap := make(map[string]interface{})
	json.Unmarshal(bytes, &tempMap)

	idiomSlice := make([]interface{}, 10)
	//idiomSlice := []interface{}{}
	//fmt.Println("json===：",tempMap)

	idiomSlice = tempMap["showapi_res_body"].(map[string]interface{})["data"].([]interface{})

	idioms = make([]string,2)
	for _, vMap := range idiomSlice {
		idiom := vMap.(map[string]interface{})["title"].(string)
		ch002 <- idiom
		idioms = append(idioms, idiom)
	}

	fmt.Println("查询相关成语===：", idioms)
	return
}

func GetIdiomBody(str string) Idiom {
	url:=idiomBodyUrl+"&keyword="+str
	bytes := RequestHttpGet(url)

	tempMap := make(map[string]interface{})
	json.Unmarshal(bytes,&tempMap)
	tempMap = tempMap["showapi_res_body"].(map[string]interface{})["data"].(map[string]interface{})

	idiom:=Idiom{}
	idiom.Title=tempMap["title"].(string)
	idiom.Cont=tempMap["content"].(string)//这里因为属性的首字母小写。序列化输出到文件，总是缺少这个属性！
	idiom.Spell=tempMap["spell"].(string)

	fmt.Println("返回成语实体===",idiom)
	downMap[str]=idiom

	//当ch001缓存读取完毕，ch002缓存读取完毕，给出查询完毕信号。
	c01l := len(ch001)
	if c01l==0 {
		c02l := len(ch002)
		if c02l==0 {
			ch003<-"Query over"
		}
	}
	return idiom
}
