const IntMax = (^uint32(0)) >> 1

func myAtoi(str string) int {
	str = strings.Trim(str, " ")
	var isNeg bool
	var num int
	for i := 0; i < len(str); i++ {
		if i == 0 {
			if str[i] == '+' {
				continue
			} else if str[i] == '-' {
				isNeg = true
				continue
			}
		}
		if str[i] >= '0' && str[i] <= '9' {
			num = num*10 + int(str[i]-'0')
		} else {
			break
		}
		if isNeg {
			if num > int(IntMax+1) {
				return ^int(IntMax)
			}
		} else {
			if num > int(IntMax) {
				return int(IntMax)
			}
		}
	}
	if isNeg {
		num = -num
	}
	return int(num)
}