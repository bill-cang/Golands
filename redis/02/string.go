/*string 【key|string/int/float】
可以对字符串进行操作，对整数型进行加减
【选项】
EX seconds – 设置键key的过期时间，单位时秒
PX milliseconds – 设置键key的过期时间，单位时毫秒
NX – 只有键key不存在的时候才会设置key的值
XX – 只有键key存在的时候才会设置key的值
*/
package main

import "fmt"

func StrFactory(r *RedisConn) {
	//添加一个key-value
	r.Do("SET","vocation", "码农")
	r.Do("SET","age", "28")
	r.Do("INCR", "age")
	reply2, err := r.Do("INCR", "vocation")
	fmt.Println("***",reply2,"\n",err)

	/*	r.Do("SET","weight", "65.3")
		reply, _ := r.Do("GET", "vocation")
		s := GetString(reply, "")
		fmt.Println("\n===String：",s)*/
/*
	//追加一个值到key上
	r.Do("APPEND","vocation",".SZ")
	reply, _ = r.Do("MGET", "vocation","age")
	//类型断言
	sp,_ := reply.([]interface{})
	s = GetString(sp[0], "")
	fmt.Printf("\n===String：%T \n %v\n",s,s)

	//INCR 和 INCRBY
	r.Do("INCR","age")
	reply2, _ := r.Do("MGET", "age", "weight")
	rep2:=reply2.([]interface{})
	age:=GetString(rep2[0],"")
	weight:=GetString(rep2[1],"")
	fmt.Println("\n===",age,"\n",weight)
*/

}
