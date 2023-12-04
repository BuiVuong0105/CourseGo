package domain

import (
	"bytes"
	"fmt"
)

type Writer interface {
	write([]byte) (int, error)
}

type ConsoleWriter struct{}

// Implement Interface Writer via pointer
func (cw *ConsoleWriter) write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func RunWriter() {
	var w Writer = &ConsoleWriter{}
	w.write([]byte("Hello World"))
}

// /////////////////////////////////////////////////////////////////////////////////////////////
type DataTypeIncrement interface {
	Increment() int
}

// Kiểu IntCounter được xây dựng từ type int
type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
func RunIncrement() {
	var myInt IntCounter = IntCounter(0)
	var inc DataTypeIncrement = &myInt
	for i := 0; i < 5; i++ {
		fmt.Println(inc.Increment())
	}
}

// COMPOSE INTERFACE
// /////////////////////////////////////////////////////////////////////////////////////////////
type Closer interface {
	close() error
}

type SystemIOComposer interface {
	Writer
	Closer
}

type BufferredWriter struct {
	buffer *bytes.Buffer
}

func (bw *BufferredWriter) write(data []byte) (int, error) {

	n, err := bw.buffer.Write(data) // append data to buffer

	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)

	for bw.buffer.Len() > 8 {

		_, err := bw.buffer.Read(v) // Đọc từ buffer vào mảng byte v với số lượng là len(v), xóa dữ liệu đọc khỏi buffer

		if err != nil {
			return 0, err
		}

		fmt.Println(string(v))
	}

	return n, nil
}

func (bw *BufferredWriter) close() error {

	for bw.buffer.Len() > 0 {

		data := bw.buffer.Next(8) // trả về 1 slice chứa n byte đọc được từ buffer và next n byte tiếp từ buffer, nếu dữ liệu trong buffer nhỏ hơn n, thì đọc toàn bộ buffer

		fmt.Println(string(data))
	}

	return nil
}

func createNewBufferredWriter() *BufferredWriter {
	return &BufferredWriter{buffer: bytes.NewBuffer([]byte{})}
}

func RunCompose() {
	// interface rỗng đóng vai trò như object là cha của các interface và struct
	// var object interface{} = WokerDomain{}
	// var object interface{} = BufferredWriter{buffer: bytes.NewBuffer([]byte{})}
	// var object interface{} = &BufferredWriter{buffer: bytes.NewBuffer([]byte{})}

	// var sysIO SystemIOComposer = createNewBufferredWriter()
	// sysIO.write([]byte("Chuc mung nam moi 2025, Happly New Year ... !"))
	// sysIO.close()

	var obj interface{} = createNewBufferredWriter()
	obj.(*BufferredWriter).write([]byte("Chuc mung nam moi 2025, Happly New Year ... !"))
	obj.(*BufferredWriter).close()

	switch obj.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case float32:
		fmt.Println("float32")
	case *BufferredWriter:
		fmt.Println("createNewBufferredWriter")
	default:
		fmt.Println("Not Found")
	}
}

// Inherite multple
// //////////////////////////////////////////////////////////////////////
type BaseWork interface {
	work()
}

type Woker struct{}

type Teacher struct{}

type WokerCompose struct {
	Woker
	Teacher
}

func (w *Woker) work() {
	fmt.Println("Worker Doing Job....")
}

func (w *Teacher) work() {
	fmt.Println("Teacher Doing Job....")
}

func (w *WokerCompose) work() {
	w.Teacher.work()
	w.Woker.work()
}

func RunWorkCompose() {
	var w BaseWork = &WokerCompose{}
	w.work()
}
