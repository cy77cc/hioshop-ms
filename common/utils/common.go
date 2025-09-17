package utils

import "time"

func GetFileExt(fileName string) string {
	fileName = fileName[len(fileName)-1:]
	return fileName
}

func GetTimestamp() int64 {
	return time.Now().Unix()
}
