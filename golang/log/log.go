package log

import (
	"fmt"
	"log"
)

func Print() {
	log.Print("this is a log")
	log.Printf("this is a log: %d", 100) // 格式化输出
	name := "wafer"
	age := 20
	log.Println(name, " ", age)
}

func Panic() {
	defer fmt.Println("defer...") //exit here
	log.Print("log...")
	log.Panic("panic log...")

	fmt.Println("ending...")
}

func Fatal() {
	defer fmt.Println("defer...")
	log.Print("log...")
	log.Fatal("fatal log...") // exit here

	fmt.Println("ending...")
}

func Flag() {

	f := log.Flags()
	fmt.Println(f)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Print("log...")
	f = log.Flags()
	fmt.Println(f)
}

func Prefix() {
	s := log.Prefix()
	fmt.Println("s: ", s)
	log.SetPrefix("[MyLog] ")
	s = log.Prefix()
	fmt.Println("s: ", s)
	log.Println("logging")
}
