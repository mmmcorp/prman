package main

type users []user

type user struct {
	Login string `json:"login"`
}
