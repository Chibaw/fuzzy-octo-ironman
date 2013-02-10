package spotify

/*
#cgo LDFLAGS: -lspotify
#include <stdio.h>
#include <libspotify/api.h>
#include "api_key.h"
void logged_in_callback(sp_session *session, sp_error err) {

}
*/
import "C"

import (
    "unsafe"
)

func InitSessionCallbacks() *C.sp_session_callbacks {

    scb := new(C.sp_session_callbacks)
    return scb
}

func InitSessionConfig() *C.sp_session_config {

    sc := new(C.sp_session_config)
    sc.api_version = C.SPOTIFY_API_VERSION
    sc.cache_location = C.CString("/tmp/spotify_cache/")
    sc.settings_location = C.CString("/tmp/spotify_config/")
    sc.application_key = ((unsafe.Pointer)(&C.g_appkey))
    sc.application_key_size = C.g_appkey_size
    sc.user_agent = C.CString("Spotify webserver")
    sc.callbacks = InitSessionCallbacks()
    sc.userdata = nil
    sc.compress_playlists = 0
    sc.dont_save_metadata_for_playlists = 0
    sc.initially_unload_playlists = 0
    sc.device_id = C.CString("72694876234987628791242")
    sc.proxy = nil
    sc.proxy_username = nil
    sc.proxy_password = nil
    sc.ca_certs_filename = C.CString(".")
    sc.tracefile = nil
    return sc
}

func Authentifaction(login string, password string) *C.sp_session {

    sc := InitSessionConfig()
    session := new(C.sp_session)
    err := C.sp_session_create(sc, &session)
    C.puts(C.sp_error_message(err))
    err = C.sp_session_login(session, C.CString(login), C.CString(password), 0, nil)
    C.puts(C.sp_error_message(err))
    return session
}

func ConnectionState(session *C.sp_session) {
    connection_state := C.sp_session_connectionstate(session)
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
}
