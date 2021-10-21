package main
//json 与struct转化
import (
	"encoding/json"
	"fmt"
	"reflect"
)
type Student struct {
	id    int
	name  string
	score float64
}
func test(){
	s := []Student{
		Student{
			1,
			"yy",
			82,
		},
		Student{
			2,
			"yang",
			67,
		},
		Student{
			3,
			"go",
			91,
		},
	}
	fmt.Println(s)
}
// 结构体首字母要大写打印出来后给的别名
type Book struct {
	Title string `json:"title"`
	Price   int    `json:"price"`
}
//struct 转json
func StructToJsonDemo() {
	books :=[]Book{
		Book{
			"alpha",
			200,
		},
		Book{
			"spider",
			200,
		},
	}
	books[0].Price=500
	book1 :=[2]Book{}
	book1[0].Price =100
	book1[0].Title="alpha"
	book1[1].Price =300

	book2 := Book{
		"alpha",
		100,
	}
	fmt.Println(book2.Title,book2.Price)
	book2.Title = "spider"
	fmt.Println(book2.Title,book2.Price)
	book3 := &Book{
		"alpha1",
		1000,
	}
	fmt.Println(book3.Title,book3.Price)
	book3.Title = "spider"
	fmt.Println(book3.Title,book3.Price)

	fmt.Println(reflect.TypeOf(book2))
	fmt.Println(reflect.TypeOf(book3))


	jsonBytes, err := json.Marshal(book1) //结构体转json序列化
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBytes))
}
//json转化为结构体
/*
func JsonToStructDemo(){
	jsonStr := `
        {
                "name_title": "alpha"
                "age_size":12
        }
        `
	var people People
	json.Unmarshal([]byte(jsonStr), &people)
	fmt.Println(people.Age,people.Name)
}


// json转map

func JsonToMapDemo(){
	jsonStr := `
        {
                "name": "alpha",
                "age": 18
        }
        `
	var mapResult map[string]interface{}//传地址才会修改该map
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}
	fmt.Println(mapResult)
}
//map转json

func MapToJsonDemo(){
	mapInstances := []map[string]interface{}{}
	instance_1 := map[string]interface{}{"name": "John", "age": 10}
	instance_2 := map[string]interface{}{"name": "Alex", "age": 12}

	mapInstances = append(mapInstances, instance_1, instance_2)

	jsonStr, err := json.Marshal(mapInstances)

	if err != nil {
		fmt.Println("MapToJsonDemo err: ", err)
	}
	fmt.Println(string(jsonStr))
}
*/
func main()  {
	//StructToJsonDemo()
	StructToJsonDemo()
	//JsonToMapDemo()
	//MapToJsonDemo()

}
