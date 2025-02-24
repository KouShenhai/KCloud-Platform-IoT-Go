package main

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	//var header = map[string]string{
	//	"Content-Type": "application/json",
	//}
	//var param = map[string]string{
	//	"id": "1",
	//}
	//body, err := core.SendRequest(http.MethodGet, "http://jsonplaceholder.typicode.com/posts", param, header)
	//if err != nil {
	//	panic(err)
	//}
	//var b []Post
	//_ = json.Unmarshal(body, &b)
	//fmt.Println(b[0].Id)

}
