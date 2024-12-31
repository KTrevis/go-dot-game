package classes

import (
	mage "server/classes/Mage"
	base_class "server/classes/base"
)

var classes = map[string]base_class.IBaseClass{
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

func GetClasses() map[string]base_class.IBaseClass {
	return classes
}

func GetClass(name string) base_class.IBaseClass {
	return classes[name]
}
