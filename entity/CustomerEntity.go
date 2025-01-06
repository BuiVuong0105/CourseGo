package entity

type CustomerEntity struct {
	code string
	name string
}

// nhận tham số là pointer or value (nhận value thì tự chuyển thành pointer thông qua &)
func (customer *CustomerEntity) UpdateInforPointer(code string, name string) {
	customer.code = code
	customer.name = name
}

// nhận tham số là pointer or value (nhận pointer thì tự chuyển thành value thông qua *)
func (customer CustomerEntity) UpdateInforValue(code string, name string) {
	customer.code = code
	customer.name = name
}

// Nếu là function chỉ chấp nhận pointer (khác với method của 1 struct nhận cả 2)
func UpdateInforPointer(customer *CustomerEntity, code string, name string) {
	customer.code = code
	customer.name = name
}

// Nếu là function chỉ chấp nhận value (khác với method của 1 struct nhận cả 2)
func UpdateInforValue(customer CustomerEntity, code string, name string) {
	customer.code = code
	customer.name = name
}
