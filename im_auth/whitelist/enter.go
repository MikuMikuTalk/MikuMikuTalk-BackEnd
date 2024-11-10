package whitelist

func IsInList(path string, whitelist []string) bool {
	for _, v := range whitelist {
		if path == v {
			return true
		}
	}
	return false
}
