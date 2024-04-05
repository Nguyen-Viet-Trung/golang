package main
import(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"time"
)
type Person struct{
	Name string  `json:"name"`// các thuộc tính phải viết hoa thì các file ngoài package mới có thể truy cập?
	Age int `json:"age`
}
func main(){
	http.HandleFunc("/startPage",  func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w,"Hello World!")
	})
	//1 middleware
	http.Handle("/middleware", middleware(http.HandlerFunc(Executing))) //=> Middleware return another function to handle
	//2+ middleware
	endhandler := ContentType(setServerTimeCookie(http.HandlerFunc(PersonExecute)))
	http.Handle("/person",endhandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	//middleware support authentication
}
func middleware(handler http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		fmt.Println("Executing Middleware before response")
		handler.ServeHTTP(w,r) // Pass control back to the handler
		fmt.Println("Executing Middleware after response")
	})
}
func Executing(w http.ResponseWriter, r *http.Request){
	println("Executing through Middleware")
	fmt.Fprintln(w,"First Middleware from book")
}
//middleware to check content type as JSON
func ContentType(handler http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-type") != "application/json"{
			w.WriteHeader(http.StatusUnsupportedMediaType)
			fmt.Fprintln(w,"Không hỗ trợ nếu k phải JSON")

			return
		}
		handler.ServeHTTP(w,r)
	})
}
//middleware to add server timestamp for response
func setServerTimeCookie(handler http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w,r)
		cookie := http.Cookie{
			Name: "Server-Time(UTC)",
			Value: strconv.FormatInt(time.Now().Unix(),10), //chuyển số giờ về thập phân
		}
		http.SetCookie(w, &cookie)
	})
}
func PersonExecute(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		var temp Person
		decode := json.NewDecoder(r.Body)
		err := decode.Decode(&temp)
		if(err != nil){
			panic(err)
		}
		defer r.Body.Close()
		fmt.Fprintf(w,"Name: %s, Age: %d", temp.Name,temp.Age)
		w.WriteHeader(http.StatusOK)
	
	}else{
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}