package main

import (
	"math/rand"
)

var letterRunes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func containsKey(log map[string]string, key string) (string, bool){
	value, ok := log[key];
	return value, ok
}

func containsValue(log map[string]string, value string) (string, bool){
	for key, iValue := range log{
		if iValue == value{
			return key, true
		}
	}
	return "", false
}

func randomPath(log map[string]string) string{
	rStream := make([]byte, 8)
	for i := range rStream{
		rStream[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	rPath := string(rStream[:8])
	_, contain := containsValue(log, rPath)
	if contain{
		return randomPath(log)
	} else{
		return rPath;
	}
}

func shorten(log map[string]string, origin string) string{
	value, ok := containsKey(log, origin);
	if ok{
		return value
	} else{
		return randomPath(log)
	}
}
