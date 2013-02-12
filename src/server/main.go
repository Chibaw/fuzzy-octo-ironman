package main

// #cgo CFLAGS: -lspotify
// #include <libspotify/api.h>
import "C"

import (
    "./webapp"
    "./spotify"
    "fmt"
    "time"
    "strings"
)

var send = make(chan string)
func WaitEvent() {
    for {
        sc := <- spotify.WaitSearchCh
        tab := spotify.TracksSearch((*C.sp_search)(sc))
        send <- strings.Join(tab,"\n")
    }
}

func SearchLol(session *C.sp_session, app chan string, send chan string) {
    for {
        action := <-app
        if action == "search" {
            option := <-app
            spotify.Search(session, option)
        }
    }
}

func main() {
    data := make(chan string)
    go webapp.ServeAll(data, send)
    fmt.Printf("Listening at localhost:8080 - Please login yourself\n")

    login := <-data
    password := <-data

    session := spotify.Authentifaction(login, password)
    is_connected := spotify.Login(session, login, password)
    go WaitEvent()
    go SearchLol(session, data, send)
    for {
        if is_connected == true {
            fmt.Printf("waiting\n")
        } else {
            is_connected = spotify.Login(session, login, password)
        }
        spotify.ConnectionState(session)
        spotify.ProcessEvents(session)
        time.Sleep(2 * time.Second)
    }
}
