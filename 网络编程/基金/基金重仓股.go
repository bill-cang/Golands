package 基金

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)


var ShaMap map[string]Share

func main() {
	address := "http://web.juhe.cn:8080/fund/zcgjj/"
	appkey := "b3adaebb0b3f6a5a72a598844baba06c"

	urlStr := address + "?key=" + appkey
	ShaMap = GetHttpData(urlStr)
	fmt.Println(ShaMap)


}

func GetHttpData(url string) (ShaMap map[string]Share) {
	//打开请求网络
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("网络请求失败，err=", err)
		return
	}
	//挂起关闭网络数据流
	defer resp.Body.Close()
	//从reposBody中解析获取数据字节
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("解析网络数据失败，err=", err)
		return
	}
	ShaMap, err = unMarShare(bytes)
	if err != nil {
		fmt.Println("json反序列化失败，err=", err)
		return
	}
	return
}

func unMarShare(bytes []byte)(ShaMap map[string]Share,err error)  {

	tempMap := make(map[string]interface{})
	//反序列化json数据
	err = json.Unmarshal(bytes, &tempMap)
	if err != nil {
		fmt.Println("json反序列化失败，err=", err)
		return
	}

	temp := tempMap["result"].([]interface{})[0].(map[string]interface{})
	//fmt.Printf("%v\n", temp)

	share := Share{}
	ShaMap = make(map[string]Share)
	le := len(temp) + 1
	for i := 1; i < le; i++ {
		key := strconv.Itoa(i)
		tmp2 := temp[key].(map[string]interface{})
		code := tmp2["code"].(string)
		share.code = code
		share.name = tmp2["name"].(string)
		share.fundnum = tmp2["fundnum"].(string)
		share.total = tmp2["total"].(string)
		share.change = tmp2["change"].(string)
		share.totalcap = tmp2["totalcap"].(string)
		share.accrate = tmp2["accrate"].(string)
		share.changesta = tmp2["changesta"].(string)
		share.time = tmp2["time"].(string)

		ShaMap[code] = share
	}
	//fmt.Printf("len:%d", len(ShaMap))
	return
}
