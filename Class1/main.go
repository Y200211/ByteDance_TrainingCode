package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	maxNum := 100
	realNum := rand.Intn(maxNum)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入数字:")
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("reader.ReadString failed, err:", err)
			return
		}
		readNum, err := strconv.Atoi(strings.TrimSuffix(readString, "\r\n"))
		if err != nil {
			fmt.Println("strconv.Atoi failed, err:", err)
		}

		if readNum > realNum {
			fmt.Println("猜错了，猜的数大于真实数")
			continue
		}
		if readNum < realNum {
			fmt.Println("猜错了，猜的数小于真实数")
			continue
		}
		if readNum == realNum {
			fmt.Println("恭喜你，猜对了！")
			break
		}
	}
}
