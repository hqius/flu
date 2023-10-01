package process

import "fmt"

var strategyMap map[string]Strategy

// Init 统一在此处注册，不要分开注册，代码需要收敛
func init() {
	strategyMap[ACCESS] = &access{}
	strategyMap[INIT] = &init_{}
	strategyMap[AUDIT] = &audit{}
	strategyMap[QUERY] = &query{}
}

// GetStrategy 外部只能获取策略，不得写入
func GetStrategy(key string) Strategy {
	if strategy, ok := strategyMap[key]; ok {
		return strategy
	}
	panic(fmt.Sprintf("strategy %s not found", key))
}
