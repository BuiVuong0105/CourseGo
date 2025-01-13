package main

import (
	"course/domain"
	"course/entity"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

func methodRecever() {

	customerEntityPointer := &entity.CustomerEntity{}
	customerEntityPointer.UpdateInforPointer("pt1", "pt2")
	fmt.Println("customerEntityPointer by method: ", customerEntityPointer)

	entity.UpdateInforPointer(customerEntityPointer, "ptV1", "ptV2")
	fmt.Println("customerEntityPointer by Function: ", customerEntityPointer)
	customerEntityPointer.UpdateInforValue("ptV1", "ptV2")
	fmt.Println("customerEntityPointer by method: ", customerEntityPointer)

	customerEntityValue := entity.CustomerEntity{}
	customerEntityValue.UpdateInforValue("vl1", "vl2")
	fmt.Println("customerEntityValue by method: ", customerEntityValue)
	(&customerEntityValue).UpdateInforValue("vl1", "vl2") // TỰ LẤY (*(&customerEntityValue)) ( PASS VALUE)
	fmt.Println("customerEntityValue by method: ", customerEntityValue)

	entity.UpdateInforValue(customerEntityValue, "vl1", "vl2")
	fmt.Println("customerEntityValue by function: ", customerEntityValue)

}

func checkNill() {
	var myInterface domain.MyInterface

	fmt.Println("MyInterFace: ", myInterface)

	var customerEntity domain.MyStructValue // default laf {}
	fmt.Println("Init Customer: ", customerEntity)
	myInterface = customerEntity
	fmt.Println("myInterface = customerEntity: ", myInterface)

	var customerEntityPointer *domain.MyStructValue // default laf nil
	fmt.Println("Init Customer Pointer: ", customerEntityPointer)
	myInterface = customerEntityPointer
	fmt.Println("myInterface = customerEntityPointer: ", myInterface)
	if myInterface == nil {
		fmt.Println("myInterface = customerEntityPointer: ", myInterface)
	}
}

func assertionType() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	myInterface, ok := i.(domain.MyStructValue) // panic
	fmt.Println(myInterface, ok)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func readerEx() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

func jsonDemo() {
	now := time.Now().UTC()
	var order *entity.Order = &entity.Order{
		Id:        1,
		Code:      "A",
		Status:    "ACTIVE",
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	jsonData, err := json.Marshal(order)
	if err != nil {
		log.Printf("ERROR: %s", err)
		return
	}
	jsonString := string(jsonData)
	log.Printf("jsonString: %v", jsonString)

	var order2 *entity.Order = &entity.Order{}
	json.Unmarshal([]byte(jsonData), order2)

	log.Printf("order2: %v", *order2)
}

func main() {
	jsonDemo()
	// domain.RunConcurrency()
	// domain.RundCmd()
	// readerEx()
	// methodRecever()
	// domain.RunInterface()
	// checkNill()
	// assertionType()
	// do(true)

	// if err := run(); err != nil {
	// 	fmt.Println(err)
	// }
}
