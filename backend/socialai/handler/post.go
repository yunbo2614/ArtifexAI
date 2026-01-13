package handler

import (
   "github.com/yunbo2614/SocialAI/backend/socialai/model"
   "github.com/yunbo2614/SocialAI/backend/socialai/service"
   "encoding/json"
   "fmt"
   "net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {

   fmt.Println("Received one upload request")
   // 1. process the request
   // json formatted string -> go struct
   // Parse from body of request to get a json object.

   decoder := json.NewDecoder(r.Body)
   var p model.Post
   if err := decoder.Decode(&p); err != nil {
       panic(err)
   }

   // 2. call service to handle request

   // 3. construct response
   fmt.Fprintf(w, "Post received: %s\n", p.Message)

}

func searchHandler(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Received one request for search")
   w.Header().Set("Content-Type", "application/json")

   // 1. process the request
   // URL param -> string variable
   user := r.URL.Query().Get("user")
   keywords := r.URL.Query().Get("keywords")

   // 2. call service to handle request
   var posts []model.Post
   var err error
   if user != "" {
       posts, err = service.SearchPostsByUser(user)
   } else {
       posts, err = service.SearchPostsByKeywords(keywords)
   }
   if err != nil {
       http.Error(w, "Failed to read post from backend", http.StatusInternalServerError)
       fmt.Printf("Failed to read post from backend %v.\n", err)
       return
   }

   // 3. construct response
   js, err := json.Marshal(posts)
   if err != nil {
       http.Error(w, "Failed to parse posts into JSON format", http.StatusInternalServerError)
       fmt.Printf("Failed to parse posts into JSON format %v.\n", err)
       return
   }
   w.Write(js)
}