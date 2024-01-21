package adapter

import (
	"fmt"
	"projects/LDmitryLD/repository/app/internal/infrastructure/db/tabler"
	"reflect"
	"strings"
)

type Condition struct {
	Equal       map[string]interface{}
	NotEqual    map[string]interface{}
	Order       []*Order
	LimitOffset *LimitOffset
	ForUpdate   bool
	Upsert      bool
}

type Order struct {
	Field string
	Asc   bool //поменял с limit
}

type LimitOffset struct {
	Offset int64
	Limit  int64
}

func filterByTag(tag string, tvalue string) func(fields *[]reflect.StructField) {
	return tabler.FilterByTags(map[string]func(value string) bool{
		tag: func(value string) bool {
			return strings.Contains(value, tvalue)
		},
	})
}

func prepareMap(info tabler.StructInfo) map[string]interface{} {
	m := make(map[string]interface{})
	for i, pointer := range info.Pointers {
		switch t := pointer.(type) {
		case *int:
			m[info.Fields[i]] = fmt.Sprintf("%v", *t)
		case *string:
			m[info.Fields[i]] = *t
		}
	}

	return m
}
