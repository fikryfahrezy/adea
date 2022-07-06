package findmostleast

func FindMostLeast(s string) [2]string {
	if s == "" {
		return [2]string{"", ""}
	}

	type cache struct {
		key        rune
		index      int
		appearance int
	}

	cacheMap := make(map[rune]cache)

	for i, c := range s {
		cm, ok := cacheMap[c]
		if !ok {
			cacheMap[c] = cache{
				key:        c,
				index:      i,
				appearance: 1,
			}
			continue
		}

		cm.appearance += 1
		cacheMap[c] = cm
	}

	least := cacheMap[rune(s[0])].key
	most := cacheMap[rune(s[0])].key
	for _, vm := range cacheMap {

		if string(vm.key) == " " {
			continue
		}

		if cacheMap[most].appearance < vm.appearance {
			most = vm.key
		}

		if cacheMap[least].appearance > vm.appearance {
			least = vm.key
		}

		if cacheMap[most].appearance == vm.appearance && cacheMap[most].index > vm.index {
			most = vm.key
		}

		if cacheMap[least].appearance == vm.appearance && cacheMap[least].index > vm.index {
			least = vm.key
		}
	}

	return [2]string{string(most), string(least)}
}
