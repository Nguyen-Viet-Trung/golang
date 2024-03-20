package main
import "fmt"
func main(){
	
	fmt.Println("Hello world")
	var user = "Trung"
	const age = 21
	var a = 14
	fmt.Println("Welcome to "+ user +"'s golang file") //+ dùng đc cho string
	fmt.Println("Trung năm nay",age,"tuổi còn Khôi năm nay",a) // còn , dùng được cho all
	var year = 2024
	fmt.Printf("Năm nay là năm %v\n",year)
	var m = "Hương"
	fmt.Printf("Mẹ tôi tên là %v\n",m)
	fmt.Printf("Tôi năm nay %v\n",age) // const vẫn là %v

	var username string
	var id int
	id = 1
	username = "Nguyen Khoi"
	fmt.Printf("Người dùng có tên là %v và có id là %v\n",username,id)
	fmt.Printf("Tên người dùng có kiểu dữ liệu là %T\n",username) //%type
	//Array
		var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
		fmt.Print(cars)
		cars[2] = "New Ford"
		fmt.Printf(cars[2])
		t:=1 // có thể áp dụng với string, bool
  		t++
		fmt.Print(t)
	
	arr1 := [6]int{10, 11, 12, 13, 14,15}
  	myslice := arr1[2:4] //slice an array

  	fmt.Printf("myslice = %v\n", myslice)
  	fmt.Printf("length = %d\n", len(myslice))
  	fmt.Printf("capacity = %d\n", cap(myslice))

	myslice1 := make([]int, 5, 10) //slice_name := make([]type, length, capacity)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	
	myslice2 := []int{1,2,3}
  	myslice3 := []int{4,5,6}
  	myslice4 := append(myslice2, myslice3...) // cú pháp cộng các array vào sau array
	myslice4 = append(myslice4, 7, 8)

  	fmt.Printf("myslice3=%v\n", myslice4)
	var x = 50
	var y = 50
	
  	if x == y {
	  fmt.Print("1")
	} else if x > y {
	  fmt.Print("2")  
	} else{
	  fmt.Print("3")   
	}

	var day= 4
  	switch day {
	case(1):
    fmt.Printf("Saturday")
  
	case(2):
    fmt.Printf("Sunday")
  
	default:
    fmt.Printf("Weekday")   
  }
  for i:=0; i < 10; i++ {
    if i==5 {
		fmt.Println(i)
		break
    }else{
		continue
	}
  }
  myFunction("John")
  fmt.Printf("%v\n",test(1))
  
 //struct
  var pers1 Person
  var pers2 Person

  // Pers1 specification
  pers1.name = "Hege"
  pers1.age = 45
  pers1.job = "Teacher"
  pers1.salary = 6000

  // Pers2 specification
  pers2.name = "Cecilie"
  pers2.age = 24
  pers2.job = "Marketing"
  pers2.salary = 4500

  // Print Pers1 info by calling a function
  printPerson(pers1)

  // Print Pers2 info by calling a function
  printPerson(pers2)

  //Map
  //b := make(map[KeyType]ValueType) -> cấu trúc của 1 map
  var q = make(map[string]string) // not null chỉ là empty map
  //var n1 map[string]int // null 
  var n = make(map[string]int)
  q["brand"] = "Ford"
  q["model"] = "Mustang"
  q["year"] = "1964"
  q["year"] = "1970" // Updating an element
  q["color"] = "red" // Adding an element
  fmt.Printf(q["brand"]) // Access an element
  delete(q,"year") //delete(map_name,key)

  n["Oslo"] = 1
  n["Bergen"] = 2
  n["Trondheim"] = 3
  n["Stavanger"] = 4

  //val, ok :=map_name[key] -> check if key exist and its value
  val1, ok1 := n["Oslo"]
  fmt.Println(val1,ok1) // key value and exist or not
  _, ok4 := n["model"]// Only checking for existing key and not its value -> use _ instead of val
  fmt.Println(ok4) 
  var n2 = make(map[string]int)
  n2 = n
  n3 := n
  for k, v := range n2 {
    fmt.Printf("%v : %v, ", k, v) //print a map without order
  }
  var order []string
  order = append(order, "Oslo","Bergen","Trondheim","Stavanger")
  for _,element := range order{
	fmt.Printf("%v : %v",element,n3[element])
  }
}
func myFunction(fname string) {
	  fmt.Printf("%v Doe\n", fname)
	}
func test(x int) int {
  	return 5 + x
}
type Person struct {
	name string
	age int
	job string
	salary int
  }
  func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
  }
  