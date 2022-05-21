package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

type MathRequest struct{
	Op string `json:"op"`
	Left float64 `json:"left"`
	Right float64 `json:"right"`
}

type MathResponse struct{
	Result float64 `json:"result"`
	Error	string 	`json:"error"`
}

func welcomeHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Welcome Gophers!")
}

func MathHandler(w http.ResponseWriter,r *http.Request){
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	req := &MathRequest{}
	if err:=dec.Decode(req); err != nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	resp := &MathResponse{}
	switch req.Op{
	case "+":
		resp.Result = req.Left + req.Right
	case "-":
		resp.Result = req.Left - req.Right
	case "*":
		resp.Result = req.Left * req.Right
	case "/":
		if req.Right == 0.0{
			resp.Error = fmt.Sprintf("Division by Zero Error")
		} else{
			resp.Result = req.Left / req.Right
		}
	default:
		resp.Error = fmt.Sprintf("Unknown Operation: %s",req.Op)
	}
	w.Header().Set("Content-Type","application/json")
	if resp.Error != ""{
		w.WriteHeader(http.StatusBadRequest)
	}
	enc := json.NewEncoder(w)
	if err:=enc.Encode(resp); err != nil{
		log.Printf("Can't encode %v - %s\n",resp,err)
	}
}

func main(){
	http.HandleFunc("/welcome",welcomeHandler)
	http.HandleFunc("/math",MathHandler)
	if err:=http.ListenAndServe(":8080",nil); err != nil{
		log.Fatal(err)
	}
}