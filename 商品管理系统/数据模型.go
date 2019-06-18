package main

import (
	"flag"
	"fmt"
	"time"
)

/*
·结合命令行参数、复合类型、包管理、排序等知识，写一个商品管理系统
·商品属性包括：整型商品ID（gid）、字符串型名称(name)、整型类别(cid)、浮点型价格(price)
·自己造一些假数据，实现如下功能：
·shop.exe -cmd single -gid xxx 查看商品ID为gid的单品详细信息
·shop.exe -cmd category -cid xxx -order 0 查看指定类别下的所有商品，cid=类别ID，order为排序规则
·排序规则参数：0=按gid升序，1=按gid降序，2=按price升序，3=按price降序
*/
//商品的分类以常量的方式声明,并iota自增长
const (
	DzProduct = (iota + 1) * 100
	Book
	Flower
)

//创建商品的主体
type Goods struct {
	id    string
	cid   int
	name  string
	price float32
}

/*创建一个保存商品的切片用来保存商品的集合,如果我们创建一个商品再把它保存到切片【对值的拷贝】，那么后期
对该商品做修改时，切片里面保存的还是创建时的信息，并且商品原有值，再拷贝，浪费内存，这里我们把切片里面
只保存商品的指针，在后期的查找商品时，只需要根据指针去找商品，做到商品对切片的影响联动*/
//按分类存放所有
var allGoods []*Goods
//按类别查询商品:map[cid]=该cid下的所有的商品的切片
var cidmap map[int][]*Goods
//按商品gid存放对应商品
var gidGoodMap map[string]*Goods

/*
包的初始化函数
会在入口函数之前被执行
写在包内任何一个go中都是可以的
*/
func init() {
	//=============================假数据
	//初始化商品切片
	allGoods = make([]*Goods, 0)
	//初始化商品类别map，并创建空切片【cidmap[DzProduct]这是个切片】,并在初始化函数里面初始化，后期使用
	cidmap = make(map[int][]*Goods)
	cidmap[DzProduct] = []*Goods{}
	cidmap[Book] = []*Goods{}
	cidmap[Flower] = []*Goods{}
	//初始化单品map
	gidGoodMap = map[string]*Goods{}

	/*创建商品假数据保存到切片，方法一：通过工厂模式，创建商品再保存到切片，后期需要修改生产模式，
	只用修改工厂生成逻辑代码一处即可，生成效率高，维护方便*/
	g1 := createGoods("1001","三星Note9", DzProduct, 6799)
	g2 := createGoods("1002","Iphone X", DzProduct, 7688)
	g3 := createGoods("1003","华为Mate20Pro", DzProduct, 5000)
	allGoods = append(allGoods, g1, g2, g3)
	//方法1.1：直接append
	allGoods = append(allGoods, createGoods("1004","玫瑰", Flower, 99))
	allGoods = append(allGoods, createGoods("1005","百合", Flower, 66))
	allGoods = append(allGoods, createGoods("1006","康乃馨", Flower, 70))
	allGoods = append(allGoods, createGoods("1007","linux入门到精通", Book, 99))
	allGoods = append(allGoods, createGoods("1008","java第三版", Book, 66))
	allGoods = append(allGoods, createGoods("1009","go语言实战", Book, 70))


	/*
		//创建商品假数据保存到切片，方法二：直接创建商品，然后取地址保存到切片(该方法需要手动填写id参数，效率低)
		allGoods = append(allGoods, &Goods{GetNewid()+strconv.Itoa(Book)+strconv.Itoa(len(cidmap[Book])+1),Book,"linux入门到精通",79.9})
		allGoods = append(allGoods, &Goods{GetNewid()+strconv.Itoa(Book)+strconv.Itoa(len(cidmap[Book])),Book,"java第三版",67.9})
		allGoods = append(allGoods, &Goods{GetNewid()+strconv.Itoa(Book)+strconv.Itoa(len(cidmap[Book])),Book,"go语言实战",56})
		*/
	//============== 打印假数据，提供输入命令参数
	for index,val := range gidGoodMap{
		fmt.Println(index,":",*val)
	}

}

func main() {
	//=============================命令
	cmdInfoArray := [3]interface{}{"cmd", "你想干嘛", "d:单品查看，s:类型查看"}
	gidInfoArray := [3]interface{}{"gid", "0", "根据商品id查看商品信息"}
	cidInfoArray := [3]interface{}{"cid", 0, "根据类型id查看该类型的商品信息"}
	orderInfoArray := [3]interface{}{"order", 0, "排序规则参数：0=按gid升序，1=按gid降序，2=按price升序，3=按price降序"}

	theResult := commanderLinByArges(cmdInfoArray, gidInfoArray, cidInfoArray, orderInfoArray)
	fmt.Println(theResult)

	switch theResult["cmd"] {
	case "d":
			gid := theResult["gid"].(string)
			goodsPointer := gidGoodMap[gid]
			if goodsPointer != nil {
				fmt.Println(*goodsPointer)
			}else {
				fmt.Println("gid有误，商品不存在！")
			}
	case "s":
		if cid, ok := theResult["cid"].(int); ok {
			goodsPointer := cidmap[cid]
			if goodsPointer ==nil {
				fmt.Println("未找到该类型商品！")
				break
			}
			for _,val:=range goodsPointer{
				fmt.Println(*val)
			}
		}else{
			fmt.Println("cid参数类型有误！")
		}
	default:
		fmt.Println("cmd必输入参数不合法！")
	}

}

//创建工厂模式生成产品
func createGoods(gid string,name string, cid int, price float32) *Goods {
	goodsPointer := new(Goods)
	goodsPointer.cid = cid
	goodsPointer.name = name
	goodsPointer.price = price

	//商品id=时间戳+商品分类+商品的个数自增长
	//gid:=GetNewid() + strconv.Itoa(cid) + strconv.Itoa(len(cidmap[cid]))
	goodsPointer.id = gid
	//按照商品的类型存放到分类map
	cidmap[cid] = append(cidmap[cid], goodsPointer)
	//按照单品的id存放对应单品map
	gidGoodMap[gid]=goodsPointer
	return goodsPointer
}

//命令行生成
func commanderLinByArges(arges ...[3]interface{}) map[string]interface{} {
	resMap := make(map[string]interface{})

	var strPointer *string

	//map 存放值之前一定要先初始化，var strFlagMap map[string]*string 只是定义
	strFlagMap := make(map[string]*string)
	intFlagMap := map[string]*int{}

	for _, conArray := range arges {
		name, para, usage := conArray[0].(string), conArray[1], conArray[2].(string)
		switch para.(type) {
		case string:
			strPointer = flag.String(name, para.(string), usage)
			strFlagMap[name] = strPointer
		case int:
			rintPointer := flag.Int(name, para.(int), usage)
			intFlagMap[name] = rintPointer
		default:
			fmt.Println(para, "类型未开发")
		}
	}
	//必须先解析，才能将参数存放到指针
	flag.Parse()

	if len(strFlagMap) > 0 {
		for key, val := range strFlagMap {
			resMap[key] = *val
		}
	}
	if len(intFlagMap) > 0 {
		for key, val := range intFlagMap {
			resMap[key] = *val
		}
	}
	return resMap
}

//时间戳id生成方法
func GetNewid() string {
	return time.Now().Format("20060102150405")
}
