package env

import (
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var (
	// Env is the environment variable
	Data = env{}
)

func init() {
	for i := 0; i < reflect.TypeOf(Data).NumField(); i++ {
		field := reflect.TypeOf(Data).Field(i)
		value := reflect.ValueOf(&Data).Elem().FieldByName(field.Name)
		upper := strings.ToUpper(field.Name)
		data := os.Getenv(upper)

		switch field.Type.Kind() {
		case reflect.String:
			value.SetString(data)
		case reflect.Int:
			casted, _ := strconv.Atoi(data)
			value.SetInt(int64(casted))
		case reflect.Bool:
			casted, _ := strconv.ParseBool(data)
			value.SetBool(casted)
		case reflect.Float64:
			casted, _ := strconv.ParseFloat(data, 64)
			value.SetFloat(casted)
		case reflect.Float32:
			casted, _ := strconv.ParseFloat(data, 32)
			value.SetFloat(casted)
		default:
			log.Println("Type not supported: ", field.Type.Kind(), " for field: ", field.Name, " with value: ", data)
		}

		log.Println("Setting ENV: ", field.Name, ": ", value)
	}
}
