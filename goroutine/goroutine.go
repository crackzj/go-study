package goroutine

/**
*	协程与文件操作
*	channel
 */
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

var channel = make(chan int)

func Write() {
	_, err := os.Stat("./test.md")
	if err != nil {
		_, err = os.Create("./test.md")
		if err != nil {
			fmt.Println("create file failed", err)
			return
		}
	}
	f, err := os.OpenFile("./test.md", os.O_RDWR, 6)
	defer f.Close()
	if err != nil {
		fmt.Println("open file failed", err)
		return
	}
	count, _ := f.Seek(0, 2)
	fmt.Println("writing file...")
	f.WriteAt([]byte("hello,hukang \n"), count)
	time.Sleep(time.Second)
	fmt.Println("writing file...")
	count, _ = f.Seek(0, 2)
	f.WriteAt([]byte("你好，胡康 \n"), count)
	fmt.Println("write end")
	channel <- 2
}

func Read() {
	<-channel
	f, err := os.OpenFile("./test.md", os.O_RDONLY, 2)
	defer f.Close()
	if err != nil {
		fmt.Println("read file failed", err)
		return
	}
	reader := bufio.NewReader(f)
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			fmt.Println("read line failed", err)
			return
		}
		if err == io.EOF {
			fmt.Println("read end")
			break
		}
		fmt.Println("read context:", string(buf))
	}
	runtime.Goexit()
}
