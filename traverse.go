package json

import (
	"reflect"
)

type JsonObject map[string]interface{}

type JsonEntity struct {
	Path  string
	Value interface{}
}

func Traverse(j map[string]interface{}) chan JsonEntity {
	entities := make(chan JsonEntity)
	done := make(chan int)
	go func() {
		<-done
		close(entities)
	}()
	go traverse(j, entities, "", done)
	return entities
}

func traverse(value interface{}, entities chan JsonEntity, basePath string, done chan int) {
	t := reflect.TypeOf(value)

	if t.Kind() == reflect.Map {
		childDone := make(chan int)

		go func() {
			count := len(value.(map[string]interface{}))
			for range childDone {
				count = count - 1
				if count == 0 {
					done <- 1
					break
				}
			}
		}()

		// if t.Elem().Kind() == reflect.String {
		// 	m := make(map[string]interface{})
		// 	for k, v := range value.(map[string]string) {}
		// }

		for k, v := range value.(map[string]interface{}) {
			go traverse(v, entities, basePath+"/"+k, childDone)
		}
	} else {
		entities <- JsonEntity{basePath, value}
		done <- 1
	}
}
