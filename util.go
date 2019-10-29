package common

import (
	"strconv"
)

func StringSliceToIntSlice(arr []string) ([]int, error) {
	toArr := make([]int, 0)
	for _, v := range arr {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		toArr = append(toArr, n)
	}

	return toArr, nil
}

func StringSliceToInt64Slice(arr []string) ([]int64, error) {
	toArr := make([]int64, 0)
	for _, v := range arr {
		n, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		toArr = append(toArr, n)
	}

	return toArr, nil
}

func StringSliceToFloat64Slice(arr []string) ([]float64, error) {
	toArr := make([]float64, 0)
	for _, v := range arr {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		toArr = append(toArr, n)
	}

	return toArr, nil
}
