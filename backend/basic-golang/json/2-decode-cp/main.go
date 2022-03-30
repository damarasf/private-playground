package main

import (
	"encoding/json"
)

// Dari contoh yang telah diberikan, cobalah untuk melakukan decode dari json menjadi objek dari struct.
// Mengambil kasus pada encode, lengkapi function DecodeToLeaderboard agar json dapat di decode menjadi objek Leaderboard

// type UserRank struct {
// 	Name  string
// 	Email string
// 	Rank  int
// }

type UserRank struct {
	// TODO: answer here
	Name  string `json:"name"`
	Email string
	Rank  int `json:"rank"`
}

type Leaderboard struct {
	Users []*UserRank
}

func DecodeToLeaderboard(jsonData []byte) (Leaderboard, error) {
	// TODO: answer here
	byteJSONData := []byte(`{
		"Users":[
		   {
			  "Name":"Roger",
			  "Email":"
			  "Rank":1
		   },
		   {
			  "Name":"Tony",
			  "Email":"
			  "Rank":2
		   },
		   {
			  "Name":"Bruce",
			  "Email":"
			  "Rank":3
		   },
		   {
			  "Name":"Natasha",
			  "Email":"
			  "Rank":4
		   },
		   {
			  "Name":"Clint",
			  "Email":"
			  "Rank":5
		   }
		]
		}`)
	leaderboard := Leaderboard{}
	err := json.Unmarshal(byteJSONData, &leaderboard)
	if err != nil {
		return leaderboard, err
	}
	return leaderboard, nil
}
