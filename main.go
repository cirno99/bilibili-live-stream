package main

import (
	biliLiveStream "bilibili-live-stream/src"
	"os"
)

func main() {
	// fmt.Println("请输入API类型: ")
	// fmt.Println("1 V1API")
	// fmt.Println("2 V2API")
	// if apiType == 1 {
	// 	biliLiveStream.V1Initialization()
	// } else {
	biliLiveStream.V2Initialization(os.Args[1])
	// }
}
