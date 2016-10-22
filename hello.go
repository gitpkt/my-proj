package main


import(
	"fmt"
	"utility"
)

type detailer interface{
	sayHi()
}

type human struct{
	name string
	age int
}

func (h *human)sayHi(){
	fmt.Printf("my name is : %s of age : %d\n",h.name,h.age)
}

type student struct{
	human   //annonynous variable type
	grade int
}

func (s *student)sayHi(){
	fmt.Printf("my name is : %s of age : %d\n",s.name,s.age)
}

type employee struct{
	human
	company string
}

func (e *employee)sayHi(){
	fmt.Printf("my name is : %s of age : %d\n",e.name,e.age)
}

type info struct{
	details []detailer
}

func (i *info)sayHi(){
	for _,h := range i.details{
		h.sayHi()
	}	
}

func main(){

	utility.PrintData("I am in a package")

	m:=make(map[string]interface{})

	m["one"] = 1
	m["two"] = 2
	m["three"] = "THREE"

	fmt.Println(m)
	
	persons := info{details:[]detailer{&student{human{"x",6},1},&employee{human{"y",37},"huawei"},&human{"z",30}}}
	persons.sayHi()

	var persons_new info
	persons_new.details = []detailer{&student{human{"a",15},1},&employee{human{"P",47},"govt"},&human{"R",40}}
	persons_new.sayHi()

	var persons_another info
	persons_another.details = make([]detailer,0,3)
	persons_another.details = append(persons_another.details,&student{human{"a",15},1})
	persons_another.details = append(persons_another.details,&employee{human{"P",47},"govt"})
	persons_another.details = append(persons_another.details,&human{"R",40})
	persons_another.sayHi()
}

