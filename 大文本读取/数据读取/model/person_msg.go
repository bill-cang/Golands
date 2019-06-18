package models

type PersonMsg struct {
	Name      string `xorm:"comment('姓名') VARCHAR(32)"`
	CardId    string `xorm:"comment('身份证号') VARCHAR(32)"`
	Sex       int    `xorm:"comment('性别 0，女，1，男') INT(11)"`
	Address   string `xorm:"comment('家庭住址') VARCHAR(128)"`
	Phone     string `xorm:"comment('手机') VARCHAR(32)"`
	Telephone string `xorm:"comment('电话') VARCHAR(32)"`
	Email     string `xorm:"comment('邮箱') VARCHAR(128)"`
}
