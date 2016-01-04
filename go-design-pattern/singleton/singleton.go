package singleton

type singeleton struct {
}

var instance *singleton

func GetInstance() *singleton {
	if insntance == nil {
		instance = &singleton{}
	}
	return instance
}
