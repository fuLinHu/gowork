package main

func out(ch  chan <- int)  {
	for i:=0;i<100;i++{
		ch <- i
	}
}
func in(ch chan <- int){
	/*for i:=range ch{
		fmt.Println(i)
	}*/
}
func main() {
	ch := make(chan  int)
	/*go in(ch)*/
	go out(ch)

}
