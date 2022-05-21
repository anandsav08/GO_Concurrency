package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"os"
)

// Create a JSON struct for extracting required fields corresponding to the actual JSON response from http.Get() call
type User struct{
	Name string `json:"name"`					// name is actual field in JSON response
	PublicRepos int `json:"public_repos"`		// public_repos is actual field in JSON response from Github API
}

func userInfo(login string) (*User,error){
	url := fmt.Sprintf("https://api.github.com/users/%s",login)
	resp,err := http.Get(url)
	if err!=nil{
		return nil,err
	}	
	defer resp.Body.Close()						// Make sure resp is closed after the function is done
	user := &User{}
	dec := json.NewDecoder(resp.Body)			// Get a JSON decoder on resp.Body
	if err:=dec.Decode(user);err!=nil{			// Decode the response and extract required info into the User struct (user)
		return nil,err
	}
	return user,nil
}

func main(){
	user_login := "tebeka"
	user,err := userInfo(user_login)
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("%+v\n",user)
}