package ref_map

import "reflect"

// MapToStruct 把map里面的字段映射到结构体对象上,dst必须是个指针
func MapToStruct(data map[string]any, dst any) {
	t := reflect.TypeOf(dst).Elem()
	v := reflect.ValueOf(dst).Elem()
	for i := range t.NumField() {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}
		mapField, ok := data[tag]
		if !ok {
			continue
		}
		val := v.Field(i)
		if field.Type.Kind() == reflect.Ptr {
			switch field.Type.Elem().Kind() {
			case reflect.String:
				mapFiledValue := reflect.TypeOf(mapField)
				if mapFiledValue.Kind() == reflect.String {
					strVal := mapField.(string)
					val.Set(reflect.ValueOf(&strVal))
				}
			default:
				panic("unhandled default case")
			}
		}
	}
}
