// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo

import (
	"gopl.io/ch9/memotest"
	"testing"
)

//最终指向一个进行网络访问的方法
var httpGetBody = memotest.HTTPGetBody

/*func TestSync(t *testing.T) {
	//以【网络访问函数】为【获取函数】创建备忘录对象
	memo := NewMemo(httpGetBody)

	//使用备忘录的【二级缓存版Get方法】同步获取多个url的页面数据
	memotest.DoGetUrlSync(t, memo)
}*/

// NOTE: not concurrency-safe!  TestSync fails.
func TestConcurrent(t *testing.T) {
	//以【网络访问函数】为【获取函数】创建备忘录对象
	memo := NewMemo(httpGetBody)

	//使用备忘录的【二级缓存版Get方法】异步获取多个url的页面数据
	memotest.DoGetUrlAync(t, memo)
}

/*
//!+output
$ go test -v gopl.io/ch9/memo1
=== RUN   TestSync
https://golang.org, 175.026418ms, 7537 bytes
https://godoc.org, 172.686825ms, 6878 bytes
https://play.golang.org, 115.762377ms, 5767 bytes
http://gopl.io, 749.887242ms, 2856 bytes

https://golang.org, 721ns, 7537 bytes
https://godoc.org, 152ns, 6878 bytes
https://play.golang.org, 205ns, 5767 bytes
http://gopl.io, 326ns, 2856 bytes
--- PASS: TestSync (1.21s)
PASS
ok  gopl.io/ch9/memo1	1.257s
//!-output
*/

/*
//!+race
$ go test -run=TestConcurrent -race -v gopl.io/ch9/memo1
=== RUN   TestConcurrent
...
WARNING: DATA RACE
Write by goroutine 36:
  runtime.mapassign1()
      ~/go/src/runtime/hashmap.go:411 +0x0
  gopl.io/ch9/memo1.(*Memo).Get()
      ~/gobook2/src/gopl.io/ch9/memo1/memo.go:32 +0x205
  ...

Previous write by goroutine 35:
  runtime.mapassign1()
      ~/go/src/runtime/hashmap.go:411 +0x0
  gopl.io/ch9/memo1.(*Memo).Get()
      ~/gobook2/src/gopl.io/ch9/memo1/memo.go:32 +0x205
...
Found 1 data race(s)
FAIL	gopl.io/ch9/memo1	2.393s
//!-race
*/
