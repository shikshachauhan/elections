package models

type Country struct {
	parties map[string]*Party
	contituencies map[string]*Constituency
}

func (c Country) GetContituency(name string) *Constituency{
	return c.contituencies[name]
}