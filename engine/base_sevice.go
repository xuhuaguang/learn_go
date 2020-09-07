package engine

type BaseService struct {
	// 要操作的model结构体[必须为指针类型的结构体*slice]
	Model       interface{} //model必须是指针
	CachePrefix string      //缓存的前缀
	// 不同的业务，有不同的库查询逻辑，所以抽象此方法，让子结构体来实现。默认方法 queryList
	QueryList func(wheres interface{}, columns interface{}, orderBy interface{}, page, rows int, total *int) (list interface{}, err error)
}
