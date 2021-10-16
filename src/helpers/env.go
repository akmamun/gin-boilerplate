package helpers

import (
	"fmt"
	"os"
)

// import "os"

// func getEnv(key, fallback string) string {
// 	value, exists := os.LookupEnv(key)
// 	if !exists {
// 		value = fallback
// 	}
// 	return value
// }

func GetEnv(key, fallback string) string {

	value, ok := os.LookupEnv(key)
	fmt.Println(value)
	fmt.Println("exuts", ok)

	if ok {
		return value
	}
	return fallback
}
