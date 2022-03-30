package main

import "encoding/json"

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
	leaderboard := Leaderboard{}
	jsonData = []byte(`{
		"Users":[
		   {
			  "name":"Roger",
			  "Email":"
			  "rank":1
		   },
		   {
			  "name":"Tony",
			  "Email":"
			  "rank":2
		   },
		   {
			  "name":"Bruce",
			  "Email":"
			  "rank":3
		   },
		   {
			  "name":"Natasha",
			  "Email":"
			  "rank":4
		   },
		   {
			  "name":"Clint",
			  "Email":"
			  "rank":5
		   }
		]
		}`)
	err := json.Unmarshal(jsonData, &leaderboard)
	if err != nil {
		return leaderboard, err
	}
	return leaderboard, nil
}
