func max(a int, b int, args ...int) int {
	mx := a
	if b > mx {
		mx = b
	}
	for _, val := range args {
		if val > mx {
			mx = val
		}
	}
	return mx
}

func min(a int, b int, args ...int) int {
	mn := a
	if b < mn {
		mn = b
	}
	for _, val := range args {
		if val < mn {
			mn = val
		}
	}
	return mn
}
