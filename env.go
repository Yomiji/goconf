package goconf

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type EnvConfiguration struct {}

func(conf EnvConfiguration) Load(obj interface{}) error {
	if reflect.TypeOf(obj).Kind() != reflect.Struct {
		return errors.New("cannot set a non-struct object")
	}
	if !reflect.ValueOf(obj).CanSet() {
		return errors.New("cannot set the given object")
	}
	env := Env{}
	for i := 0; i < reflect.TypeOf(obj).NumField(); i++ {
		fieldVal := reflect.ValueOf(obj).Field(i)
		// skip unsettables
		if !fieldVal.CanSet() {
			continue
		}
		// get the type of the field, so that we can grab the name of the field
		fieldName := reflect.TypeOf(obj).Field(0).Name
		// depending on the type, convert the environment variable
		switch fieldVal.Type().Kind() {
	  case reflect.String:
	  	fieldVal.Set(reflect.ValueOf(env.getEnv(fieldName)))
		case reflect.Int:
			fieldVal.Set(reflect.ValueOf(env.getIntEnv(fieldName)))
		case reflect.Float32:
			fieldVal.Set(reflect.ValueOf(env.getFloat32Env(fieldName)))
		case reflect.Float64:
			fieldVal.Set(reflect.ValueOf(env.getFloat64Env(fieldName)))
	  }
	}
	return nil
}
type Env struct{}

func(env Env) getFloat64Env(envVar string) interface{} {
	checkEnv(envVar)
	f,err := strconv.ParseFloat(os.Getenv(envVar), 64)
	if err != nil {
		panic(err)
	}
	return f
}

func(env Env) getFloat32Env(envVar string) interface{} {
	checkEnv(envVar)
	f,err := strconv.ParseFloat(os.Getenv(envVar), 32)
	if err != nil {
		panic(err)
	}
	return float32(f)
}

func(env Env) getIntEnv(envVar string) interface{} {
	checkEnv(envVar)
	i,err := strconv.Atoi(os.Getenv(envVar))
	if err != nil {
		panic(err)
	}
	return i
}

func(env Env) getEnv(envVar string) string {
	checkEnv(envVar)
	return os.Getenv(envVar)
}

func checkEnv(env string) {
	if _,ok := os.LookupEnv(env); !ok {
		err := errors.New(fmt.Sprintf("%s environment variable not present", env))
		panic(err)
	}
}