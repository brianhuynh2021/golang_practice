package utils

import (
)


func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}