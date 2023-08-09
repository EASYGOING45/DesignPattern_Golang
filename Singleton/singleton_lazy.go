package singleton

import "sync"

var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

//GetLazyInstance 懒汉式单例
func GetLazyInstance() *Singleton {
	if lazySingleton != nil {
		return lazySingleton
	}

	//once 内的方法只会执行一次，所以不需要再次判断
	once.Do(func() {
		lazySingleton = &singleton{a: "test"}
	})

	return lazySingleton
}
