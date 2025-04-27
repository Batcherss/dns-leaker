package utils
// fucked our brain
func Compare(reference []string, target []string) (bool, []string, []string) {
	refMap := make(map[string]bool)
	for _, ip := range reference {
		refMap[ip] = true
	}

	targetMap := make(map[string]bool)
	for _, ip := range target {
		targetMap[ip] = true
	}

	missing := []string{}
	extra := []string{}

	for ip := range refMap {
		if !targetMap[ip] {
			missing = append(missing, ip)
		}
	}
	for ip := range targetMap {
		if !refMap[ip] {
			extra = append(extra, ip)
		}
	}

	isSame := len(missing) == 0 && len(extra) == 0
	return isSame, missing, extra
}
