package lead_to_offer_ii

func addBinary(a, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	sa, sb := []byte(a), []byte(b)
	reverse(sa)
	reverse(sb)
	sizeA, sizeB := len(sa), len(sb)
	res := make([]byte, sizeA+1)
	carry := 0
	for i := 0; i < sizeA; i++ {
		if sa[i] == '1' {
			carry++
		}

		if i < sizeB && sb[i] == '1' {
			carry++
		}

		if carry == 1 || carry == 3 {
			res[i] = '1'
		} else {
			res[i] = '0'
		}

		carry >>= 1
	}

	if carry == 1 {
		res[sizeA] = '1'
	}

	reverse(res)
	if carry == 1 {
		return string(res)
	} else {
		return string(res[1:])
	}
}

func reverse(bytes []byte) {
	n := len(bytes)
	for i := 0; i < n/2; i++ {
		bytes[i], bytes[n-1-i] = bytes[n-1-i], bytes[i]
	}
}
