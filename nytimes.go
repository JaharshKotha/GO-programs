package main

import (
	"fmt"
	"net/http"
    "log"
	"io/ioutil"
    "strings"
    "trie"
)

var m map[string]string =make(map[string]string)
var al []string 
var t = trie.New() 

func build_dictionary(ip string){
    
    sp:= strings.Split(ip,"\n")
    for _,c := range sp{
        x := strings.TrimSpace(c)
        t.Add(x,1)
	} 
     
   
}


func display_dictionary(ip string){
    
     
    var tem string
    
    ip=ip+"#"
    var flg int
    for itr:=0;itr<len(ip)-1;itr=itr+1{
        
      ch:=ip[itr]
        tem=tem+string(ch)
        
        if t.HasChildren(tem,rune(ip[itr+1])){
            if _, ok := t.Find(tem);ok{
            fmt.Println(tem)
            flg=itr
            }
        }else{
            if _, ok := t.Find(tem);ok{
                 fmt.Println(tem)
            flg=itr
            }
          tem=""
         itr = flg  
        }
        
    }
    
}



func display_dictionary2(ip string){
    
     
    var tem string
    for i:=0;i<len(ip);i=i+1{
        for j:=i;j<len(ip);j=j+1{
            tem=ip[i:j+1]
            if _, ok := t.Find(tem);ok{
                fmt.Println(tem)
            }
        }
    }
    
}








func uploadHandler(w http.ResponseWriter, r *http.Request) {
	
    var s string
    bs,er := ioutil.ReadAll(r.Body)
    if er != nil{
    log.Println(er)
    http.Error(w,"Error ",http.StatusInternalServerError)
    return
    }
    s = string(bs)
    build_dictionary(s)
   
	
	
}


func uploadHandler2(w http.ResponseWriter, r *http.Request) {
	
    var s string
    bs,er := ioutil.ReadAll(r.Body)
    if er != nil{
    log.Println(er)
    http.Error(w,"Error ",http.StatusInternalServerError)
    return
    }
    s = string(bs)
    display_dictionary(s)
   
	
	
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
    http.HandleFunc("/display", uploadHandler2)
	http.ListenAndServe(":5050", nil)
}

