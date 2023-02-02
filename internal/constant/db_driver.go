package constant

import "strings"

type Driver int

const (
	MYSQL Driver = iota
	POSTGRES
)


	var driverList = []string{
		"MYSQL",
		"POSTGRES",
	}

func (d Driver) String() string {
	return driverList[d]
}

func ConvertDriver(in string) Driver {
	for idx,v := range driverList{
		if strings.EqualFold(v,in){
			return Driver(idx)
		}
	}
	return -1
}
