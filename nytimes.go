//curl -X POST -H "Content-Type: application/octet-stream" -H "X-Session-Token: XC09" --data-binary '@words.txt' http://127.0.0.1:5050/dictionary

package main

import (
	"fmt"
	"net/http"
    "log"
	"io/ioutil"
    "strings"
    "user_messaging_challenge/trie"
    "github.com/gorilla/mux"
)



type Health_Param struct {
    param_name        string   
    param_status      string   
    
}

type authenticationMiddleware struct {
    tokenUsers map[string]string
}

func (amw *authenticationMiddleware) Populate() {
    amw.tokenUsers = make(map[string]string)
    amw.tokenUsers["00000000"] = "user0"
    amw.tokenUsers["XC09"] = "user0"
    amw.tokenUsers["aaaaaaaa"] = "userA"
    amw.tokenUsers["05f717e5"] = "randomUser"
    amw.tokenUsers["deadbeef"] = "user0"
}

func AuthRequired(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")
        if token=="XC09"{
            handler.ServeHTTP(w, r)
            return
        }
	
        fmt.Println("Not Authenticated")
	}
}

    
func (amw *authenticationMiddleware) Middleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("X-Session-Token")
        if user, found := amw.tokenUsers[token]; found {
            log.Printf("Authenticated user %s\n", user)
            next.ServeHTTP(w, r)
        } else {
            http.Error(w, "Forbidden", http.StatusForbidden)
        }
    }
}


var m map[string]string =make(map[string]string)
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
            flg=itr
            }
        }else{
            if _, ok := t.Find(tem);ok{
            flg=itr
            }
            tem=tem[0:(len(tem)-(itr-flg))]
            fmt.Println(tem)
            tem=""
            itr = flg  
        }
        
    }
    
}



func backtrack_display_dictionary(ip string){
    tem:=""
    ip=ip+"#"
    for i:=0;i<len(ip);i=i+1{
        tem=tem+string(ip[i])
        for j:=i+1;j<len(ip);j=j+1{
            if t.HasChildren(tem,rune(ip[j])){
            tem=ip[i:j+1]
            if _, ok := t.Find(tem);ok{
                fmt.Println(tem)
            }
            }else{
                tem=""
                break
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


func displayHandler(w http.ResponseWriter, r *http.Request) {
	
    var s string
    bs,er := ioutil.ReadAll(r.Body)
    if er != nil{
    log.Println(er)
    http.Error(w,"Error ",http.StatusInternalServerError)
    return
    }
    s = string(bs)
    backtrack_display_dictionary(s)
}


func healthcheck(w http.ResponseWriter, r *http.Request) {
	
//    //simple_check:=Health_Param{"status","OK"}
//    jData, err := json.MarshalIndent(simple_check,"","  ")
//if err != nil {
//    panic(err)
//    return
//}
//    fmt.Println(string(jData))
//w.Header().Set("Content-Type", "application/json")
    var jsonBlob = []byte(`{"status": "OK"}`)
//    w.Header().Set("Content-Type", "application/json")
//    json.NewEncoder(w).Encode(jsonBlob)
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonBlob)
	
}
 

func main() {
    // Init router
	r := mux.NewRouter()

    amw := authenticationMiddleware{ }
    amw.Populate()
    
    // Route handles & endpoints
    r.HandleFunc("/dictionary", amw.Middleware(uploadHandler)).Methods("POST")
    r.HandleFunc("/split", AuthRequired(displayHandler)).Methods("POST")
	r.HandleFunc("/healthcheck", healthcheck).Methods("GET")
	
    
    
    
   
	// Start server
	log.Fatal(http.ListenAndServe(":5050", r))
    
	
}

