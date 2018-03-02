//curl -X POST -H "Content-Type: application/octet-stream" -H "X-Session-Token: XC09" --data-binary '@concatenated.txt' http://127.0.0.1:5050/split

package main

import (
	"fmt"
	"net/http"
    "log"
	"io/ioutil"
    "strings"
    "messaging-interview/utils/trie"
    "github.com/gorilla/mux"
    "unicode"
)


// Creating a struct so that we can have a map of the struct literals representing the different checks such as DB connectivity , Cache systems (Redis) ,etc..

type Health_Param struct {
    param_name        string   `json:"name"`
    param_status      string   `json:"kind"`   
}

//Useful to create a map of different roles when each of them would have varying priviledges ,we can restrict the access based on the Key given 
type authenticationMiddleware struct {
    tokenUsers map[string]string
}

func (amw *authenticationMiddleware) Populate() {
    amw.tokenUsers = make(map[string]string)
    amw.tokenUsers["XC092WER34SE2"] = "regularUser"
    amw.tokenUsers["ACG75PIR98SH1"] = "superUser"
}

//Check if the input key is valid

func (amw *authenticationMiddleware) Middleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("X-Session-Token")
        if user, found := amw.tokenUsers[token]; found {
            log.Printf("Authenticated %s\n", user)
            next.ServeHTTP(w, r)
        } else {
            http.Error(w, "Forbidden", http.StatusForbidden)
        }
    }
}


var m map[string]string =make(map[string]string)
var t = trie.New() 

//Build the Trie

func build_dictionary(ip string){
    
    sp:= strings.Split(ip,"\n")
    for _,c := range sp{
        x := strings.TrimSpace(c)
        t.Add(x,1)
	} 
     
   
}


//Display the Exact split of the concatenated string

func display_dictionary(ip string) string{
    var tem,generate string    
    ip=ip+"#"
    var flg int
    for itr:=0;itr<len(ip)-1;itr=itr+1{       
        ch:=ip[itr]
        if !unicode.IsLetter(rune(ch)) {    //Check for alphabetic input - Input Handling! 
            continue
        }
        tem=tem+string(ch)
        if t.HasChildren(tem,rune(ip[itr+1])){  //Build the string if you are still part of the Trie 
            if _, ok := t.Find(tem);ok{
            flg=itr                            //Keep track of the last valid word position
            }
        }else{
            if _, ok := t.Find(tem);ok{  
            flg=itr
            }
            tem=tem[0:(len(tem)-(itr-flg))]  //Print the word formed so far
            generate=generate+tem
            fmt.Println(tem)
            tem=""
            itr = flg  
        }
        
    }
    
    return generate
    
}



func backtrack_display_dictionary(ip string){
    tem:=""
    ip=ip+"#"
    for i:=0;i<len(ip);i=i+1{                   // Traverse throught the concatenated string
        tem=tem+string(ip[i])
        for j:=i+1;j<len(ip);j=j+1{
            if t.HasChildren(tem,rune(ip[j])){  //Check if word is not anymore a part of the Trie else Bactrack ; This avoids bruteforcing
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

// POST dictionary handler

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

//POST exact split handler

func displayHandler(w http.ResponseWriter, r *http.Request) {
	
    var s string
    bs,er := ioutil.ReadAll(r.Body)
    if er != nil{
    log.Println(er)
    http.Error(w,"Error ",http.StatusInternalServerError)
    return
    }
    s = string(bs)
    res:=display_dictionary(s)
    mySlice := []byte(res)
    w.Write(mySlice)
}

//POST split handler 
func displayAllHandler(w http.ResponseWriter, r *http.Request) {
	
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
	
//    simple_check:=Health_Param{param_name:"status",param_status:"OK"}
//    fmt.Println(simple_check)
//    jData, err := json.Marshal(simple_check)
//    if err != nil {
//        panic(err)
//        return
//    }
//    fmt.Println(string(jData))
//  w.Header().Set("Content-Type", "application/json")
    var jsonBlob = []byte(`{"status": "OK"}`)
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
    r.HandleFunc("/splitExact", amw.Middleware(displayHandler)).Methods("POST")
    r.HandleFunc("/split", amw.Middleware(displayAllHandler)).Methods("POST")
	r.HandleFunc("/healthcheck", healthcheck).Methods("GET")

    // Start server
	log.Fatal(http.ListenAndServe(":5050", r))
    
	
}

