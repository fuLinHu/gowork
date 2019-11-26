package main

import "log"

func main() {
	//log.Panicln("这是一条会触发panic的日志。")
	//panic("日志是否打印了")
	/*log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	log.Fatalln("这是一条会触发fatal的日志。")*/
	//panic("日志是否打印了")

	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.Fatalln("这是一条会触发fatal的日志。")

}
