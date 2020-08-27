package main

import (
	"errors"
	"fmt"
)

type Dictionary map[string]string

var notFoundErr error = errors.New("this key is not defined")
var keyExitedErr error = errors.New("this key has been defined")

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", notFoundErr
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) (string, error) {
	fmt.Printf("function %p\n", &d)
	_, ok := d[key]
	if !ok {
		fmt.Println("ket is not defined  value:::", value)
		d[key] = value
		return value, nil
	}
	return "", keyExitedErr
}

func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
}

func main() {
	d := &Dictionary{}
	fmt.Printf("main %p\n", &d)
	value, ok := d.Add("test", "it is just a test")
	if ok == nil {
		fmt.Println(value)
	} else {
		fmt.Println(ok)
	}
}
