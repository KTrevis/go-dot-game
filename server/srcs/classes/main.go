package classes

import (
	mage "server/classes/Mage"
	"server/classes/base"
)

var classes = map[string]base.IBaseClass{
	"Mage": &mage.Mage{
		Name: "Mage",
	},
}

func GetClassesName() []string {
	var names []string
	for key := range classes {
		names = append(names, key)
	}

	return names
}

func GetClasses() map[string]base.IBaseClass {
	return classes
}

func GetClass(name string) base.IBaseClass {
	return classes[name]
}
