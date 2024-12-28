package mage

type Mage struct {
	Name	string
}

func (this *Mage) Setup(name string) {
	this.Name = name
}
