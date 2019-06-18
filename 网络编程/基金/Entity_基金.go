package 基金

type Share struct {
	code      string //股票代码
	name      string //名称
	fundnum   string //持有基金家数
	total     string //持股总量（万股）
	change    string //持股变化和上季度比
	totalcap  string //持股总市值(万元)
	accrate   string //流通市值占比（%）
	changesta string //遇上季变化
	time      string //时间
}