package 基金
import (
	"io/ioutil"
	"net/http"
	"net/url"
	"fmt"
	"encoding/json"
)

//----------------------------------
// 重仓股基金调用示例代码 － 聚合数据
// 在线接口文档：http://www.juhe.cn/docs/27
//----------------------------------

const APPKEY = "b3adaebb0b3f6a5a72a598844baba06c" //您申请的APPKEY
/*
func main(){

	//1.重仓股基金
	Request1()

}*/
//1.重仓股基金
func Request1(){
	//请求地址
	juheURL :="http://web.juhe.cn:8080/fund/zcgjj/"

	//初始化参数
	param:=url.Values{}

	//配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
	param.Set("key",APPKEY) //APPKEY值


	//发送请求
	data,err:=Get(juheURL,param)
	if err!=nil{
		fmt.Errorf("请求失败,错误信息:\r\n%v",err)
	}else{
		var netReturn map[string]interface{}
		json.Unmarshal(data,&netReturn)
		if netReturn["error_code"].(float64)==0{
			fmt.Printf("接口返回result字段是:\r\n%v",netReturn["result"])
		}
	}
}



// get 网络请求
func Get(apiURL string,params url.Values)(rs[]byte ,err error){
	var Url *url.URL
	Url,err=url.Parse(apiURL)
	if err!=nil{
		fmt.Printf("解析url错误:\r\n%v",err)
		return nil,err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery=params.Encode()
	resp,err:=http.Get(Url.String())
	if err!=nil{
		fmt.Println("err:",err)
		return nil,err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// post 网络请求 ,params 是url.Values类型
func Post(apiURL string, params url.Values)(rs[]byte,err error){
	resp,err:=http.PostForm(apiURL, params)
	if err!=nil{
		return nil ,err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
