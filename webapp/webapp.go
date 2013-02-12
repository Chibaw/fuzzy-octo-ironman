package webapp

import (
    "net/http"
    "log"
    "fmt"
    "html"
    "io"
    "os"
)

func SendFile(w io.Writer, filename string) {

    fmt.Printf("filename: %s\n", filename)
    file, err := os.Open(filename)
    if err == nil {
       io.Copy(w, file)
        file.Close()
    }
}

func HandleViewFile(w http.ResponseWriter, r *http.Request) {

    fmt.Printf("url requested: %s\n", r.URL.Path)
    if r.URL.Path == "/" {
        SendFile(w, "views/login.html")
    } else {
        SendFile(w, "views" + html.EscapeString(r.URL.Path))
    }
}

func HandleAssetsFile(w http.ResponseWriter, r *http.Request) {

    fmt.Printf("url requested: %s\n", r.URL.Path)
    SendFile(w, "assets" + html.EscapeString(r.URL.Path))
}
func ServeViews() {

    http.HandleFunc("/", HandleViewFile)
    http.HandleFunc("/index.html", HandleViewFile)
}

func ServeAssets() {
    http.HandleFunc("/js/app.js", HandleAssetsFile)
    http.HandleFunc("/js/ember.js", HandleAssetsFile)
    http.HandleFunc("/js/jquery.js", HandleAssetsFile)
    http.HandleFunc("/js/handlebars.js", HandleAssetsFile)
}

func LoginMethods(ch chan string) {

    http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            fmt.Printf("url requested: %s\n", r.URL.Path)
            SendFile(w, "views" + html.EscapeString(r.URL.Path) + ".html")
        } else if r.Method == "POST" {
            fmt.Printf("Name: %s\nPass: lol\n", r.FormValue("login"))
            ch <- r.FormValue("login")
            ch <- r.FormValue("password")
            http.Redirect(w, r, "/search", http.StatusFound)
        }
    })
}

func SearchMethods(app chan string, send chan string) {
    http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            app <- "search"
            app <- r.FormValue("search")
            fmt.Fprintf(w, "%s", <-send)
        } else if r.Method == "GET" {
            fmt.Printf("url requested: %s\n", r.URL.Path)
            SendFile(w, "views" + html.EscapeString(r.URL.Path) + ".html")
        }
    })
}

func ServeData(app chan string, send chan string) {
    LoginMethods(app)
    SearchMethods(app, send)
}

func ServeAll(app chan string, send chan string) {

    ServeViews()
    ServeAssets()
    ServeData(app, send)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
