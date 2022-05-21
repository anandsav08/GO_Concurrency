package main

import(
	"fmt"
	"strings"
	"bufio"
	"crypto/md5"
	"os"
	"io"
)

func parseSignatureFile(path string) (map[string]string,error){
	file,err := os.Open(path)
	if err != nil{
		return nil,err
	}
	defer file.Close()
	file_sig := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := strings.Fields(scanner.Text())
		if len(line) != 2{
			return nil,fmt.Errorf("%s bad line \n",path)
		}
		file_sig[line[1]] = line[0]
	}
	return file_sig,nil
}

func fileMd5(path string) (string,error){
	file,err := os.Open("nasa-logs/"+path)
	if err!= nil{
		return "",err
	}
	defer file.Close()
	hash := md5.New()
	_,err = io.Copy(hash,file)
	if err!=nil{
		return "",err
	}
	return fmt.Sprintf("%x",hash.Sum(nil)),nil
}

type result struct{
	path string
	match bool
	err error
}

func md5Worker(path string,sig string,out chan *result){
	r := &result{path: path}
	s,err := fileMd5(path)
	if err!=nil{
		r.err = err
		out <- r
		return
	}
	r.match = (sig == s)
	r.err = nil
	out <- r
}

func main(){
	file_map,err := parseSignatureFile("nasa-logs/md5sum.txt")
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	ch := make(chan *result)
	for key,val := range file_map{
		go md5Worker(key,val,ch)
	}

	for range file_map{
		res := <- ch
		if res.err != nil{
			fmt.Printf("%s - Error %s\n",res.path,res.err)
		} else if res.match == false{
			fmt.Printf("%s signature Mismatch \n",res.path)
		} else{
			fmt.Printf("%s signature matched \n",res.path)
		}
	}
}