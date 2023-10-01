package factory

import "fmt"

var factoryMap map[string]IGun

func init() {
	factoryMap["ak47"] = newAk47()
	factoryMap["musket"] = newAk47()
}

func GetGun(key string) IGun {
	if gun, ok := factoryMap[key]; ok {
		return gun
	}
	panic(fmt.Sprintf("Wrong gun type %s passed", key))
}
