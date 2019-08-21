package main

type rdsConfig struct {
	Host string
	Port int
	User string
	Pwd string
	Enable bool
}

var rdsCacheList = []rdsConfig{
	{
		"127.0.0.1",
		6379,
		"root",
		"123456",
		true,
	},
}

var RdsCache = rdsCacheList[0]
