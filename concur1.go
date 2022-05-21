package main

import (

	"fmt"
	"net/http"
)

func returnType(url string){
	resp,err := http.Get(url)
	if err!=nil{
		fmt.Printf("ERROR : %s\n",err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("Content-Type")
	fmt.Printf("%s -> %s\n",url,ctype)
}

func main(){
	urls := []string{
		"http://www.google.com",
		"http://www.linkedin.com",
		"http://www.youtube.com",
		"http://www.dota2.com",
	}
	for _,url := range urls{
		returnType(url)
	}
}