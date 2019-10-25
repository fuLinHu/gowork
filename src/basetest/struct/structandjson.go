package main

type Student struct {
	name string
	age  int
	adre
	sore []int
}

type adre struct {
	city    string
	country string
}

func main() {
	s := []int{}
	var ints []int
	for i := 0; i < 10; i++ {
		ints = append(s, i)
	}
	student := new(Student)
	student.name = "付林虎"
	student.age = 20
	student.sore = ints
	a := adre{
		city:    "北京",
		country: "中国",
	}
	student.adre = a
	//bytes, e := json.
	/*fmt.Println(e)
	if(e==nil){
		fmt.Println(bytes)
	}*/

}
