package main
import("fmt"
		"log"
	"net/http"
)
func LogInHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{ //FailHandler
		fmt.Fprintf(w,"Đã xảy ra lỗi %v",err)
	}
	fmt.Fprintf(w, "hoàn thành POST\n")
	name := r.FormValue(("username")) // truy thẻ ở file html để nhận giá trị 
	password := r.FormValue("password") 

	fmt.Fprintf(w,"Kiểm tra nhận POST\n")
	fmt.Fprintf(w,"Username: %v\n",name) //trả về giá trị vừa nhận được
	fmt.Fprintf(w,"Password: %v",password)
}
func main(){
	http.HandleFunc("/startPage",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"This is starter page!")
	})
	fileServer := http.FileServer(http.Dir("LocalHost/static"))
	http.Handle("/",fileServer) // đường dẫn mặc định
	
	http.HandleFunc("/login", LogInHandler)
	
	if err := http.ListenAndServe(":8080", nil); err != nil{ // giá trị đầu là port, giá trị sau là để cấu hình máy chủ 
		log.Fatal(err) // sử dụng method listenAndServe từ thư viện http để bắt đầu web server		
	}

}
