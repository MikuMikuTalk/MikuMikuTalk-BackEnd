package list_util

// DeduplicationList 去重
func DeduplicationList[T string | int | uint32](req []T) (response []T) {
	i32Map := make(map[T]bool)
	for _, i32 := range req {
		if !i32Map[i32] {
			i32Map[i32] = true
		}
	}
	for key, _ := range i32Map {
		response = append(response, key)
	}
	return
}
