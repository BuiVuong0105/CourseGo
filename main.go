package main

import (
	"context"
	"course/domain"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

func goroutineExample() {

	start := time.Now()
	workerDomain := domain.WokerDomain{}
	workerDomain.SetName("vuongbv")

	// Sử dụng channel thì việc gửi và nhận phải ở trong các gorountine khác nhau, khi gửi vào thì nếu vượt quá buffer mà channel không có ai đăng ký nhận sẽ bị báo lỗi
	resultChannel := make(chan string, 10)
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(3)

	go workerDomain.DoWork(1*time.Second, "Fetch Account", waitGroup, resultChannel)
	go workerDomain.DoWork(2*time.Second, "Get Account By Id", waitGroup, resultChannel)
	go workerDomain.DoWork(3*time.Second, "Search Account", waitGroup, resultChannel)

	go func() {
		waitGroup.Wait()
		close(resultChannel)
		fmt.Println("----------------Close Channel Success----------------------")
	}()

	time.Sleep(time.Second * 5)

	workerDomain.ReadChannel(resultChannel)
	_, ok := <-resultChannel
	fmt.Println("RS: ", ok)

	fmt.Printf("Total Dowork: %v second \n", time.Since(start))
}

// func divide(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("Cannot divide by zero")
// 	}
// 	return (a / b), nil
// }

// func caculate() {
// 	result, err := divide(5, 0)
// 	if err != nil {
// 		fmt.Println("ERROR: ", err)
// 	} else {
// 		fmt.Println("RESULT: ", result)
// 	}
// }

// func anynomouseFunction() {
// 	for i := 0; i <= 10; i++ {
// 		func(i int) {
// 			fmt.Println("Anynomouse Function: ", i)
// 		}(i)
// 	}
// }

func calculateData(ctx context.Context, resultCh chan<- int) {
	time.Sleep(5 * time.Second)
	resultCh <- 42
}

func runTimeOut() {
	// Tạo một context với timeout là 3 giây
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Tạo một channel để nhận kết quả từ goroutine
	resultCh := make(chan int)

	// Chạy hàm tính toán trong một goroutine
	go calculateData(ctx, resultCh)

	// Sử dụng select để kiểm tra kết quả hoặc timeout
	select {
	case result := <-resultCh:
		fmt.Printf("Received result: %d\n", result)
	case <-ctx.Done():
		fmt.Println("Function finished due to timeout or canceled context")
	}

	// Tiếp tục chương trình
	fmt.Println("Main program continues")
}

func test() {
	fmt.Println("TEST")
}

// Hàm thực thi sẽ được gọi auto
func init() {
	fmt.Println("INIT SYSTEM")
}

const (
	httpReadTimeout  = 2 * time.Minute
	httpWriteTimeout = time.Hour
)

var (
	listenAddr = 3000
)

const (
	flagAddr   = "addr"
	flagDBPath = "db_path"
	flagEnv    = "env"
)

// StartCommand ...
func StartCommand() *cobra.Command {
	var listenAddr, dbPath, envName string
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start server ",
		RunE: func(cmd *cobra.Command, args []string) error {
			return StartServer(listenAddr, dbPath, envName)
		},
	}
	cmd.Flags().StringVarP(&listenAddr, flagAddr, "1", ":3000", "Address to open rest server")
	cmd.Flags().StringVarP(&dbPath, flagDBPath, "2", os.ExpandEnv("$HOME/.hashcode"), "db path")
	cmd.Flags().StringVarP(&envName, flagEnv, "3", "", "env default: prod")
	return cmd
}

// StartServer ...
func StartServer(listenAddr, dbPath, envName string) error {

	// handler, err := core.Run("35056997918317u85346cccx3i@123", store, envName)
	// if err != nil {
	// 	return err
	// }

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})

	h := c.Handler(nil)

	srv := &http.Server{
		Handler: h,
		Addr:    listenAddr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: httpWriteTimeout,
		ReadTimeout:  httpReadTimeout,
	}
	log.Printf("http listening on port http://localhost%s", listenAddr)
	return srv.ListenAndServe()
}

// func main() {

// rootCmd := &cobra.Command{}
// cobra.EnableCommandSorting = false
// rootCmd.AddCommand(StartCommand())
// err := rootCmd.Execute()
// if err != nil {
// 	// handle with #870
// 	panic(err)
// }

// fmt.Println(quote.Go())
// test := test
// test()
// fmt.Println()
// goroutineExample()
// caculate()
// anynomouseFunction()
// domain.RunWriter()
// domain.RunIncrement()
// domain.RunCompose()
// domain.RunWorkCompose()
// domain.RunInterface()
// domain.CaculateSQRT()
// hashCode := domain.NewHashCode()
// qrCodes := hashCode.EndCode(26691, 1000000)
// fmt.Println(qrCodes)
// runTimeOut()
// singleton.TestSingletonPattern()
// builder.TestBuilder()

// wg := sync.WaitGroup{}
// wg.Add(2)
// resultChannel := make(chan int, 2)

// go func() {
// 	fmt.Println("RS: ", <-resultChannel) // Đoc channel nếu đọc mà không có data thì báo lỗi
// 	wg.Done()
// }()
// go func() {
// 	resultChannel <- 42 // khi gửi vào thì nếu vượt quá buffer mà channel không có ai đăng ký nhận sẽ bị báo lỗi
// 	wg.Done()
// }()
// wg.Wait()

// }
