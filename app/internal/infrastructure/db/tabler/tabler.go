package tabler

import "reflect"

type Tabler interface {
	TableName() string
}

type StructInfo struct {
	Fields   []string
	Pointers []interface{}
}

func GetStructInfo(u interface{}, args ...func(*[]reflect.StructField)) StructInfo {
	val := reflect.ValueOf(u).Elem()
	var structFields []reflect.StructField

	for i := 0; i < val.NumField(); i++ {
		structFields = append(structFields, val.Type().Field(i))
	}

	for i := range args {
		if args[i] == nil {
			continue
		}
		args[i](&structFields)
	}

	var res StructInfo

	for _, field := range structFields {
		valueField := val.FieldByName(field.Name)
		res.Pointers = append(res.Pointers, valueField.Addr().Interface())
		res.Fields = append(res.Fields, field.Tag.Get("db"))
	}

	return res
}

func FilterByTags(tags map[string]func(value string) bool) func(fields *[]reflect.StructField) {
	return func(fields *[]reflect.StructField) {
		var res []reflect.StructField
		for _, field := range *fields {
			for key, fn := range tags {
				tagVal := field.Tag.Get(key)
				if fn(tagVal) {
					res = append(res, field)
				}
			}

		}
		*fields = res
	}
}
