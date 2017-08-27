package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type User struct {
	id uint32
	email string
	first_name string
	last_name string
	gender string
	birth_date float64
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readUsers() (users []User) {
	var parsedChunk map[string]interface{}
	dat, err := ioutil.ReadFile("hlcupdocs/data/TRAIN/data/users_1.json")
	check(err)
	err = json.Unmarshal(dat, &parsedChunk)
	check(err)
	userMaps := parsedChunk["users"].([]interface{})
	for _, u := range userMaps {
		var user User
		user.id = uint32(u.(map[string]interface{})["id"].(float64))
		user.email = u.(map[string]interface{})["email"].(string)
		user.first_name = u.(map[string]interface{})["first_name"].(string)
		user.last_name = u.(map[string]interface{})["last_name"].(string)
		user.gender = u.(map[string]interface{})["gender"].(string)
		user.birth_date = u.(map[string]interface{})["birth_date"].(float64)
		fmt.Println(user)
	}
	return
}

func main () {
	readUsers()
}