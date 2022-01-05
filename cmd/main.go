package main

import (
	"flag"
	_ "github.com/lib/pq"
)

var resourcesPath = flag.String("f", ".", "set resources path viper will loading.")

func main() {
	//todo 1.注册好我的额client，注册到依赖管理里面
	//todo 2.在controller里定义一个api，传入用户名和repo名字，得到对应的commit列表
	//todo 3.在service里引用client，然后用service调用这个client取得commit记录
	flag.Parse()
	application, clean, err := CreateApp(*resourcesPath)
	defer clean()
	if err != nil {
		panic(err)
	}
	if err := application.Start(); err != nil {
		panic(err)
	}
	application.AwaitSignal()
}
