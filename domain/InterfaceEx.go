package domain

import "fmt"

type MyInterface interface {
	SetName(name string)
	GetName() string
}

type MyStructValue struct {
	Name string
}

/// implement Interface có thể nhận con trỏ hoặc giá trị
func (msv MyStructValue) SetName(name string) {
	msv.Name = name
	fmt.Println("Method called on value: ", msv.Name)
}

func (msv MyStructValue) GetName() string {
	return msv.Name
}

func (msv MyStructValue) MethodSpecial() {
	fmt.Println("Method called on value")
}

type MyStructPointer struct {
	Name string
}

/// implement Interface chỉ nhận giá trị địa chỉ
func (msp *MyStructPointer) SetName(name string) {
	msp.Name = name
	fmt.Println("Method called on Pointer: ", msp.Name)
}

func (msp *MyStructPointer) GetName() string {
	return msp.Name
}

func (msv *MyStructPointer) MethodSpecial(name string) {
	msv.Name = name
	fmt.Println("Method called on value")
}

func RunInterface() {

	var myStructValueImpl MyInterface = &MyStructValue{}
	myStructValueImpl.SetName("AAAAA") // tự động chuyển thành giá trị và clone object mới rồi setName trên object mới đó
	println("myStructValueImpl", myStructValueImpl.GetName())

	myStructValueImpl = MyStructValue{}
	myStructValueImpl.SetName("BBBB")
	println("myStructValueImpl", myStructValueImpl.GetName())

	var myStructValueImplPointer MyInterface = &MyStructPointer{}
	myStructValueImplPointer.SetName("AAAAA")
	println("myStructValueImplPointer: ", myStructValueImplPointer.GetName())

	var myStructPointer *MyStructPointer = &MyStructPointer{}
	myStructPointer.MethodSpecial("ToanBV")
	fmt.Println("My Struct Pointer: ", *myStructPointer)
}
