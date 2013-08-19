package main

import "fmt"
import "io/ioutil"
import "track/db"
// import "tomtom/data"
import "net/http"
import "encoding/json"
import "time"
//import "log"
// import "code.google.com/p/goauth2/oauth"
// import "github.com/gorilla/sessions"
//import "strconv"

// var store = sessions.NewCookieStore([]byte("tomtom-secret-key"))

/*
func getUserId (w http.ResponseWriter, r *http.Request) string {
    session, err := store.Get(r, "session")
    if err != nil {
        http.Redirect(w, r, "/", http.StatusFound)
    }
    

    id := session.Values["UserId"].(string)
    
    if len(id) == 0 {
        http.Redirect(w, r, "/", http.StatusFound)
    }
    
    return id
}
*/

/*
func authenticationHandler (w http.ResponseWriter, r *http.Request) {
    url := oauthCfg.AuthCodeURL("")
    http.Redirect (w, r, url, http.StatusFound)
}
*/

type User struct {
    Id string
    Name string
    Given_Name string
    Family_Name string
    Link string
    Gender string
    Locale string
}

/*
func oauthCallbackHandler (w http.ResponseWriter, r *http.Request) {
    profileInfoURL := "https://www.googleapis.com/oauth2/v1/userinfo?alt=json"
    code := r.FormValue ("code")
    t := oauth.Transport { Config: oauthCfg }
    t.Exchange (code)
    resp, err := t.Client().Get(profileInfoURL)
    if err != nil { 
        panic (err)
    }
    defer resp.Body.Close()

    user := User {}
    contents, err := ioutil.ReadAll(resp.Body)
    json.Unmarshal (contents, &user)
    
    session, _ := store.Get (r, "session")
    session.Values["UserId"] = user.Id
    session.Values["GivenName"] = user.Given_Name
    log.Printf("** %s Logged in **", user.Name)
    session.Save (r, w)
    
    http.Redirect(w, r, "/view/", http.StatusFound)
}
*/

func getItemsHandler (w http.ResponseWriter, r *http.Request) {
    searchString := r.FormValue ("name")
    items, _ := db.FindItems (searchString)

    w.Header ().Add ("Content-Type", "application/json")
    
    data, _ := json.Marshal (items);

    fmt.Fprintf (w, "%s", data)
}

func getVendorsHandler (w http.ResponseWriter, r *http.Request) {
    searchString := r.FormValue ("name")
    items, _ := db.FindVendors (searchString)

    w.Header ().Add ("Content-Type", "application/json")
    
    data, _ := json.Marshal (items);

    fmt.Fprintf (w, "%s", data)
}
func saveHandler (w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        panic(err)
    }
    type JsonData struct {
        Item string
        Quantity float32
        Price float32
        Vendor string
        Date time.Time
    }
    type JsonRetData struct {
        Msg string
    }
    
    var jsonData JsonData
    json.Unmarshal (body, &jsonData)
    
    defer r.Body.Close()

    err = db.Save (jsonData.Item, jsonData.Quantity, jsonData.Price, jsonData.Vendor, time.Now())
    
    
    var jsonRetData JsonRetData
    jsonRetData.Msg = "Saved"
    
    if err != nil {
        fmt.Println (err)
        jsonRetData.Msg = err.Error()
    }
    
    w.Header ().Add ("Content-Type", "application/json")
    data, err := json.Marshal (jsonRetData)

    fmt.Fprintf (w, "%s", data)
}

func initWebServer() {
    http.HandleFunc ("/save", saveHandler)
    http.HandleFunc ("/items", getItemsHandler)
    http.Handle ("/view/", http.StripPrefix("/view/", http.FileServer(http.Dir("/home/saaadhu/code/git/track/src/track/www"))))
    http.ListenAndServe(":8081", nil)
}

func main() {
    initWebServer()
}
