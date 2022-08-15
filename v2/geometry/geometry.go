package geometry

import "errors"

func CubeVolume(n int) (int, error) {
	if n == 0 {
		return 0, errors.New("zero lenght edge is not allowed")
	}
	return n * n * n, nil
}
