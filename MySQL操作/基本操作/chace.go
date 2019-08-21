package main

import "sort"

var CacheMap = make(map[string]cache) //缓存map

//缓存的结构体
type cache struct {
	con    *[]kfperson //人员的切片
	uptime int64       //缓存最后一次更新时间
	cou    int         //缓存的使用次数
}

//缓存清理：时间最早剔除
func clrCacheMap(cp *map[string]cache) {
	mmp := *cp
	var sli = make([]int64, len(mmp))
	var tmp =make(map[int64]string , len(mmp))
	var i = 0
	for key, val := range mmp {
		tmkey:=val.uptime
		sli[i]=tmkey
		tmp[tmkey]=key
		i++
	}
	sort.Slice(sli, func(i, j int) bool {
		return false
	})
	delete(CacheMap, tmp[sli[0]])
}
