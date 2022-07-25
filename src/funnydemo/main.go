package main

import (
	"fmt"
	"funnydemo/entity"
	util2 "funnydemo/util"
	"github.com/json-iterator/go"
	"os"
)

func main() {
	util2.GetIp()
	fmt.Printf("%s:first", "dd")
	fmt.Println()
	// ================= 序列化 =====================
	group := entity.ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := jsoniter.Marshal(group)
	if err != nil {
		fmt.Println("error: ", err)
	}
	os.Stdout.Write(b)
	fmt.Println()
	fmt.Println(string(b[:]))
	fmt.Println()

	var jsonBlob = []byte(`[
        {"Name": "Platypus", "Order": "Monotremata"},
        {"Name": "Quoll",    "Order": "Dasyuromorphia"}
    ]`)
	var animals []entity.Animal
	if err := jsoniter.Unmarshal(jsonBlob, &animals); err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Printf("the unmarshal is  %+v", animals)
}
