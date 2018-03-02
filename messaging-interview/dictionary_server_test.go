//curl -X POST -H "Content-Type: application/octet-stream" -H "X-Session-Token: XC09" --data-binary '@concatenated.txt' http://127.0.0.1:5050/split

package main

import (
	"fmt"
	"testing"
    "io/ioutil"
    "net/http"
    "bytes"
)

func TestSpliter(t *testing.T){
    
    
    
    b, err := ioutil.ReadFile("./assets/words.txt") 
    if err != nil {
        fmt.Print(err)
    }
    
    req, err := http.NewRequest("POST", "http://127.0.0.1:5050/dictionary",bytes.NewBuffer(b))
    req.Header.Set("X-Session-Token", "XC09")
    req.Header.Set("Content-Type", "application/octet-stream")

    client := &http.Client{}
     client.Do(req)
    
      
    
    b1, err1 := ioutil.ReadFile("./testing/concatenated_test.txt") 
    if err1 != nil {
        fmt.Print(err1)
    }
    
    req1, err := http.NewRequest("POST", "http://127.0.0.1:5050/splitExact",bytes.NewBuffer(b1))
    req1.Header.Set("X-Session-Token", "XC09")
    req1.Header.Set("Content-Type", "application/octet-stream")

    client1 := &http.Client{}
    resp1, err := client1.Do(req1)
    if err != nil {
        panic(err)
    }
    defer resp1.Body.Close()

    fmt.Println("response Status:", resp1.Status)
    fmt.Println("response Headers:", resp1.Header)
    body1, _ := ioutil.ReadAll(resp1.Body)
    fmt.Println("response Body:", string(body1))
    
    if string(body1)!="superpermpulverizegleefulnessovernippingAstarte"{
        t.Fail()
        t.Log("Was expecting superperm but received",string(body1))
    }
    
    
    
    
    
    
    
}

