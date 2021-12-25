package helpers

import (
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

func GetEnv(key, defValue string) string {

	value, ok := os.LookupEnv(key)

	if ok {
		return value
	}
	return defValue
}
