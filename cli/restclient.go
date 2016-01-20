package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"encoding/json"
)

type Todo struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}
type Todos []Todo

func main() {
	if len(os.Args) != 2 {
		fmt.Println(" Difficult to call a rest URI without the URI")
		os.Exit(1)
	}

	var restURIString = os.Args[1]

	client := &http.Client{
		CheckRedirect: nil,
	}

	var resp *http.Response
	var req *http.Request
	req, _ = http.NewRequest("GET", restURIString, nil)

	resp, _ = client.Do(req)
	//fmt.Print(resp)
	jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var todos Todos
	err = json.Unmarshal([]byte(jsonDataFromHttp), &todos) // here!

	if err != nil {
		panic(err)
	}

	// test struct data
	fmt.Println(todos)

}
