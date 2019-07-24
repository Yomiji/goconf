package goconf

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)



func FromEnvironment(obj interface{}) error {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return errors.New("parameter must be pointer")
	}
	//get underlying struct
	objVal := reflect.ValueOf(obj).Elem()
	objType := objVal.Type()
	if objType.Kind() != reflect.Struct {
		return errors.New("cannot set a non-struct object")
	}
	for i := 0; i < objType.NumField(); i++ {
		fieldVal := objVal.Field(i)
		// get the type of the field, so that we can grab the name of the field
		fieldName := objType.Field(0).Name

		// skip unsettables
		if !fieldVal.CanSet() {
			continue
		}
		tag := objType.Field(i).Tag
		if tag.Get("env") != "" {
			fieldName = tag.Get("env")
		}
		// depending on the type, convert the environment variable
		switch fieldVal.Type().Kind() {
	  case reflect.String:
	  	fieldVal.Set(reflect.ValueOf(getEnv(fieldName)))
		case reflect.Int:
			fieldVal.Set(reflect.ValueOf(getIntEnv(fieldName)))
		case reflect.Float32:
			fieldVal.Set(reflect.ValueOf(getFloat32Env(fieldName)))
		case reflect.Float64:
			fieldVal.Set(reflect.ValueOf(getFloat64Env(fieldName)))
	  }
	}
	return nil
}

func  getFloat64Env(envVar string) interface{} {
	checkEnv(envVar)
	f,err := strconv.ParseFloat(os.Getenv(envVar), 64)
	if err != nil {
		panic(err)
	}
	return f
}

func getFloat32Env(envVar string) interface{} {
	checkEnv(envVar)
	f,err := strconv.ParseFloat(os.Getenv(envVar), 32)
	if err != nil {
		panic(err)
	}
	return float32(f)
}

func getIntEnv(envVar string) interface{} {
	checkEnv(envVar)
	i,err := strconv.Atoi(os.Getenv(envVar))
	if err != nil {
		panic(err)
	}
	return i
}

func getEnv(envVar string) string {
	checkEnv(envVar)
	return os.Getenv(envVar)
}

func checkEnv(env string) {
	if _,ok := os.LookupEnv(env); !ok {
		err := errors.New(fmt.Sprintf("%s environment variable not present", env))
		panic(err)
	}
}