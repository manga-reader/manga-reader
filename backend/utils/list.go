package utils

func ReverseStringSlice(list []string) []string {
	_list := make([]string, len(list))
	copy(_list, list)
	for i, j := 0, len(_list)-1; i < j; i, j = i+1, j-1 {
		_list[i], _list[j] = _list[j], _list[i]
	}
	return _list
}
