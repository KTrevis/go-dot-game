package mage

type Mage struct {
	Name	string
}

func (this *Mage) GetName() string {
	return this.Name
}
