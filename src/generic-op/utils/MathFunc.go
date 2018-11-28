package utils

func AbsInt(val int) (ret int) {
	ret = val
	if val < 0 {
		ret = -val
	}
	return ret
}

func AbsInt64(val int64) (ret int64) {
	ret = val
	if val < 0 {
		ret = -val
	}
	return ret
}
