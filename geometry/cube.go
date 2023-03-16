package geometry

import "errors"

func Cubevolume(n int) (int, error) {
	if n != 0 {
		return n * n * n, nil
	} else {
		return 0, errors.New("Zero length edge is not allowed")
	}
}
