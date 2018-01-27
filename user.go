package main

type users []user

type user struct {
	Login string `json:"login"`
}

func (u users) contains(username string) bool {
	for _, user := range u {
		if user.Login == username {
			return true
		}
	}
	return false
}
