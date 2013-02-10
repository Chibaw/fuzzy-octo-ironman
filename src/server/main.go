package main

// #cgo CFLAGS: -lspotify
// #include <libspotify/api.h>
import (
    "./spotify"
    "time"
)

func main() {
    session := spotify.Authentifaction(,)
    for {
        spotify.ConnectionState(session)
        time.Sleep(1 * time.Second)
    }
}
