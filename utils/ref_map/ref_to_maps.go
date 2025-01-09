package ref_map

import "reflect"

func RefToMap(data any, tag string) map[string]any {
	maps := map[string]any{}
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		getTag, ok := field.Tag.Lookup(tag)
		if !ok {
			continue
		}
		val := v.Field(i)
		if val.IsZero() {
			continue
		}
		if field.Type.Kind() == reflect.Struct {
			newMaps := RefToMap(val.Interface(), tag)
			maps[getTag] = newMaps
			continue
		}
		if field.Type.Kind() == reflect.Ptr {
			if field.Type.Elem().Kind() == reflect.Struct {
				newMaps := RefToMap(val.Elem().Interface(), tag)
				maps[getTag] = newMaps
				continue
			}
			maps[getTag] = val.Elem().Interface()
			continue
		}
		maps[getTag] = val.Interface()
	}
	return maps
}

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
