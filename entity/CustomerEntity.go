package entity

type CustomerEntity struct {
	code string
	name string
}

func (customer *CustomerEntity) UpdateInforPointer(code string, name string) {
	customer.code = code
	customer.name = name
}

func (customer CustomerEntity) UpdateInforValue(code string, name string) {
	customer.code = code
	customer.name = name
}
