package utils

import "strconv"

func GetDefaultIntVal(envVar string, d int) int {
	i, err := strconv.Atoi(envVar)
	if err != nil {
		return d
	}

	return i
}

func GetDefaultStringVal(envVar string, d string) string {
	if len(envVar) == 0 {
		return d
	}
	return envVar
}
