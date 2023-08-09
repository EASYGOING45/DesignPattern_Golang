package singleton

//Singleton 饿汉式单例
type Singleton struct {
	a string // 测试字段
}

var singleton *Singleton // 创建一个单例对象

func init() {
	singleton = &Singleton{a: "test"}
}

//GetInstance 获取实例
func GetInstance() *Singleton {
	return singleton
}
