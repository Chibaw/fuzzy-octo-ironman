package spotify

/*
#cgo LDFLAGS: -lspotify
#include <stdio.h>
#include <libspotify/api.h>
#include "session_callbacks.h"
*/
import "C"

import (
    "unsafe"
    "fmt"
)

func initSessionCallbacks() *C.sp_session_callbacks {
    scb := new(C.sp_session_callbacks)
    return C.init_callbacks(scb)
}

func initSessionConfig() *C.sp_session_config {

    sc := new(C.sp_session_config)
    sc.api_version = C.SPOTIFY_API_VERSION
    sc.cache_location = C.CString("tmp")
    sc.settings_location = C.CString("tmp")
    sc.application_key = ((unsafe.Pointer)(&C.g_appkey))
    sc.application_key_size = C.g_appkey_size
    sc.user_agent = C.CString("Spotify webserver")
    sc.callbacks = initSessionCallbacks()
    sc.userdata = nil
    sc.compress_playlists = 0
    sc.dont_save_metadata_for_playlists = 0
    sc.initially_unload_playlists = 0
    //sc.device_id = C.CString("72694876234987628791242")
    sc.proxy = nil
    sc.proxy_username = nil
    sc.proxy_password = nil
    sc.tracefile = nil
    return sc
}

func Authentifaction(login string, password string) *C.sp_session {

    sc := initSessionConfig()
    session := new(C.sp_session)
    err := C.sp_session_create(sc, &session)
    fmt.Printf("%d\n", err)
    C.puts(C.sp_error_message(err))
    C.sp_session_set_connection_rules(session, C.SP_CONNECTION_RULE_NETWORK)
    return session
}

func Login(session *C.sp_session, login string, password string) bool {

    C.sp_session_login(session, C.CString(login), C.CString(password), 0, nil)
    cs := ConnectionState(session)
    if cs == C.SP_CONNECTION_STATE_LOGGED_IN {
        return true
    }
    return false
}

func ConnectionState(session *C.sp_session) C.sp_connectionstate {

    connection_state := C.sp_session_connectionstate(session)
    fmt.Printf("ConnectionState: %d\n", connection_state)
    if connection_state == C.SP_CONNECTION_STATE_LOGGED_IN {
        C.puts(C.CString("Logged_in"))
    } else if connection_state == C.SP_CONNECTION_STATE_LOGGED_OUT {
        C.puts(C.CString("Logged_out"))
    } else if connection_state == C.SP_CONNECTION_STATE_DISCONNECTED {
        C.puts(C.CString("Disconnected"))
    } else if connection_state == C.SP_CONNECTION_STATE_UNDEFINED {
        C.puts(C.CString("Undefined"))
    } else if connection_state == C.SP_CONNECTION_STATE_OFFLINE {
        C.puts(C.CString("Offline"))
    }
    return connection_state
}

var WaitSearchCh = make(chan unsafe.Pointer)

//export GoWaitSearch
func GoWaitSearch(sc unsafe.Pointer) {
    WaitSearchCh <- sc
}

func Search(session *C.sp_session, what string) {
    C.sp_search_create(
        session,
        C.CString(what),
        0, 5,
        0, 5,
        0, 5,
        0, 0,
        C.SP_SEARCH_STANDARD,
        C.get_search_cb(), nil)
}

func ProcessEvents(session *C.sp_session) {
    var i C.int
    C.sp_session_process_events(session, &i)
}

func TracksSearch(res *C.sp_search) []string {
    nb := C.sp_search_num_tracks(res)
    tab := make([]string, nb)
    var i C.int = 0
    for ; i < nb; i++ {
        C.puts(C.sp_track_name(C.sp_search_track(res, i)))
        tab[i] = C.GoString(C.sp_track_name(C.sp_search_track(res, i)))
    }
    return tab
}

func DeleteSearch(res *C.sp_search) {
    C.sp_search_release(res)
}
