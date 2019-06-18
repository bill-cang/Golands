// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 272.

//!+

// Package memo provides a concurrency-unsafe
// memoization of a function of type Func.
package memo

//备忘录：如果缓存cache中有结果就拿取缓存，缓存中没有就调用f去获取结果（获取到了以后再丢一份到缓存）
// A Memo caches the results of calling a Func.
type Memo struct {
	//结果获取函数
	f     Func
	//结果缓存容器：string->url
	cache map[string]result
}

//Func是类型【func(key string) (interface{}, error)】的别名
//【func(key string) (interface{}, error)】是一个可以自定义的抽象方法，具体实现可以是：网络访问、数据库访问、文件访问...
// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

//结果数据模型
type result struct {
	//正常结果
	value interface{}
	//拿不到结果时的异常
	err   error
}

//备忘录的工厂方法
func NewMemo(f Func) *Memo {
	//创建一个以f为【获取函数】，并将缓存容器初始化了的备忘录对象的指针
	return &Memo{f: f, cache: make(map[string]result)}
}

//备忘录的数据获取方法
//逻辑出发点：有缓存拿缓存，没缓存调用【获取函数f】进行获取（获取到了以后再丢一份到缓存）
//备忘录是【数据获取器接口IDataGetter】的【二级缓存版实现】
// NOTE: not concurrency-safe!
func (memo *Memo) Get(key string) (interface{}, error) {
//判断缓存里有没有数据
res, ok := memo.cache[key]

//缓存没有数据，则调用f进行获取，拿到后再丢一份去缓存
if !ok {
//则调用f进行获取
res.value, res.err = memo.f(key)
//拿到后再丢一份去缓存
memo.cache[key] = res
}

//最终返回真正的结果
return res.value, res.err
}

//!-
