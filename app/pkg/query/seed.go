package query

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func CreateSeed(v ...any) {
	for _, x := range v {
		t := reflect.TypeOf(x)

		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}

		if t.Kind() != reflect.Struct {
			fmt.Println("Input need struct")
			return
		}

		var structure []map[string]string

		fields := make(map[string]string)

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)

			jsonTag := field.Tag.Get("json")
			structureTag := field.Tag.Get("structure")

			fields[jsonTag] = structureTag
		}

		structure = append(structure, fields)

		data, err := json.MarshalIndent(structure, "", "    ")
		if err != nil {
			fmt.Println("Error create json : ", err)
			return
		}

		err = os.WriteFile(fmt.Sprintf("./app/migration/seed/%s.json", t.Name()), data, 0644)
		if err != nil {
			fmt.Println("Error create file : ", err)
			return
		}

		fmt.Println(t.Name(), " success created")
	}

	fmt.Println("done.")
}
