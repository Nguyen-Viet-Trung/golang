package main
import(
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)
func main(){
	db, err := sql.Open("mysql","root:Trung@614@tcp(127.0.0.1:3306)/devmaster")
	if(err != nil){
		panic(err.Error())
	}
	defer db.Close()
	insert, err := db.Query("insert into student values(7,'Quan','Nguyen','HN',3)")
	if(err != nil){
		panic(err.Error())
	}
	defer insert.Close()
	delete, err := db.Query("delete from student where first_name ='Tho'")
	if(err != nil){
		panic(err.Error())
	}
	defer delete.Close()
	update, err:= db.Query("update student set last_name ='Chu' where id = 1")
	if(err != nil){
		panic(err.Error())
	}
	defer update.Close()
	/*Notable: từ khóa defer trong Golang
	Các lệnh được gọi qua từ khóa defer sẽ được đưa vào một stack, tức là hoạt động theo cơ chế vào sau ra trước (last-in-first-out). Lệnh nào defer sau sẽ được thực thi trước, 
	giống như xếp 1 chồng đĩa thì chiếc đĩa sau cùng (ở trên cùng) sẽ được lấy ra trước.
	Khi gọi lệnh defer thì giá trị của biến trong câu lệnh sẽ là giá trị tại thời điểm gọi 
	chứ không phải giá trị tại thời điểm thực thi. */
}