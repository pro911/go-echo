package main

import (
	"fmt"
	"net/rpc"
	"os"
	"strconv"
)

func main() {

	f := "add"
	nums := make([]int, 0)
	for k,arg := range os.Args {
		if k == 0 {
			continue
		}
		if k == 1 {
			f = arg
			continue
		}
		i,err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("参数错误", k)
		}
		nums = append(nums, i)
	}

	serviceMethod := "Math.Add"
	switch f {
	case "add":
		serviceMethod = "Math.Add"
	case "sub":
		serviceMethod = "Math.Sub"
	}

	client, err:= rpc.Dial("tcp", ":18082")
	if err != nil {
		panic(err.Error())
	}

	var reply int
	err = client.Call(serviceMethod, nums, &reply)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(reply)
}