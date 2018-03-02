
Steps to Execution:

Make sure that the entire zip folder is kept in the $GOPATH src folder , this avoids ambiguity with other dependencies and is a standard way of maintaining projects. We utilize MUX for handling routing hence make sure the unzipped folder has "github.com/gorilla/mux" in the same level as the project.

At the time of execution of the code in my machine , this is my directory structure "/Users/jaharsh/go/src/messaging-interview" and "/Users/jaharsh/go/src/github.com/gorilla/mux" ; here "Users/jaharsh/go/" is my $GOPATH

Once you have established setting up , you start the server by "go run dictionary_server.go"  or "./dictionary_server" , keep the terminal and open a new instance to test the api

`POST /dictionary`:
curl -X POST -H "Content-Type: application/octet-stream" -H "X-Session-Token: XC092WER34SE2" --data-binary '@./assets/words.txt' http://127.0.0.1:5050/dictionary

`POST /split`:
curl -X POST -H "Content-Type: application/octet-stream" -H "X-Session-Token: XC092WER34SE2" --data-binary '@./assets/concatenated.txt' http://127.0.0.1:5050/split

`GET /healthcheck`:
curl  http://127.0.0.1:5050/healthcheck

Additional Endpoint:

`POST /splitExact`:
curl -X POST -H "Content-Type: application/octet-stream" -H "X-Session-Token: XC092WER34SE2" --data-binary '@./assets/concatenated.txt' http://127.0.0.1:5050/splitExact


You can verify the output on the console of the Server(not Client) as stated in the question. The Client displays some returned http data which is used in writing unit test cases.


Description:

During the initialization we construct a Trie from the given strings in the words.txt .The construction of this Trie would cost us O(M*N) where M is the length of the longest word and N is the number of words, this is because we might need to traverse in completely different direction and keep constructing nodes for words having only partial similarity, but everything pays of during the traversal part.  

When the POST /splitExact is hit, we traverse through the concatenated string by keeping track of word in the trie and seeing if we have reached any complete word , if we do then we keep track of this position with a flag and see if it can be pushed , if it can't then we print the word achieved so far and continue traversing the string while starting from the root of the trie. Here we traverse through the string while keeping track of the last full word with a flag and backtrack to it if we do not see a word hence its O(L*(N-L)) where L is the length of longest partial word and N is the length of the concatenated string

When the POST /split is hit, we traverse through the string w.r.t to the Trie similar to the previous approach except that we do not keep track of a flag. Every time we encounter a word not in the Trie we continue the traversal building from the next letter of our concatenated string, the worst case time complexity of this approach is O(L*N) where L is the length of longest partial word and N is the length of the concatenated string

Testing:

We write some unit test cases using the GO's testing tool , here we check the splitExact endpoint by concatenating our output from the endpoint into a string and comparing it with the concatenated string input.