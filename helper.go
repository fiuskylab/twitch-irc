package twitchirc

func inArrayStr(arr []string, target string) (int, bool) {
	for pos, i := range arr {
		if i == target {
			return pos, true
		}
	}
	return -1, false
}
