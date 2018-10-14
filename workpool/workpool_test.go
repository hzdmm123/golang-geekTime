package workpool

import (
	"github.com/goinggo/workpool"
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
	"testing"
)

type MyWork struct {
	Name      string "The Name of a person"
	BirthYear int    "The Yea the person was born"
	WP        *workpool.WorkPool
}

func (workPool *MyWork) DoWork(workRoutine int) {
	fmt.Printf("%s : %d\n", workPool.Name, workPool.BirthYear)
	fmt.Printf("*******> workRoutine: %d  QueuedWork: %d  ActiveRoutines: %d\n", workRoutine, workPool.WP.QueuedWork(), workPool.WP.ActiveRoutines())
	time.Sleep(1000 * time.Millisecond)
	//panic("test")
}

func TestNew(t *testing.T) {
	main()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	workPool := workpool.New(runtime.NumCPU()*3, 10)
	shutdown := false // Just for testing, I Know
	go func() {
		for i := 0; i < 10; i++ {
			work := &MyWork{
				Name:      "A" + strconv.Itoa(i),
				BirthYear: i,
				WP:        workPool,
			}
			err := workPool.PostWork("name_routine", work)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err)
				time.Sleep(100 * time.Millisecond)
			}
			if shutdown == true {
				return
			}
		}
	}()
	fmt.Println("Hit any key to exit")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	shutdown = true
	fmt.Println("Shutting Down\n")
	workPool.Shutdown("name_routine")
}
