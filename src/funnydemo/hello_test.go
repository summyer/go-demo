package main

import (
	"encoding/json"
	"fmt"
	"funny/entity"
	util2 "funny/util"
	"github.com/graphql-go/graphql"
	jsoniter "github.com/json-iterator/go"
	"log"
	"os"
)

func main() {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
        {
            hello
        }
    `
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}
func test() {
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
