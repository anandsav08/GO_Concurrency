package main
import (
	"fmt"
	"net/http"
)

func main(){
	sites := []string{
		"http://www.google.com",
		"http://www.youtube.com",
		"http://www.facebook.com",
		"http://www.linkedin.com",
	}
	ch := make(chan string)
	for _,url := range sites{
		go func(url string){
			resp,err := http.Get(url)
			if err!=nil{
				fmt.Printf("ERROR: %s (%s)\n",url,err)
				return
			}
			defer resp.Body.Close()
			ctype := resp.Header.Get("Content-Type")
			ch <- ctype
		}(url)
	}

	for _,url := range sites{
		ctype := <-ch
		fmt.Printf("(%s) TYPE: %s\n",url,ctype)
	}
}