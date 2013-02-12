#include <stdio.h>
#include <libspotify/api.h>
#include "session_callbacks.h"

void logged_in_callback(sp_session *session, sp_error err) {

    puts("logged_in_callback called");
}

void logged_out_callback(sp_session *session) {
    puts("logged_out_callback called");
}

void metadata_updated_callback(sp_session *session) {
    puts("metadata_updated_callback called");
}

void message_to_user_callback(sp_session *session, const char *msg) {
    puts("message_to_user_callback called");
}

void notify_main_thread_callback(sp_session *session) {
    printf("notify_main_thread\n");
}

int music_delivery_callback(sp_session *session,
                    const sp_audioformat *format,
                    const void *frames,
                    int num_frames) {
    puts("music_delivery called");
}

void play_token_lost_callback(sp_session *session) {
    puts("play_token_lost_callback called");
}

void log_message_callback(sp_session *session, const char *msg) {
    printf("log_message_callback called: %s\n", msg);
}

void end_of_track_callback(sp_session *session) {
    puts("end_of_track_callback called");
}

void streaming_error_callback(sp_session *session, sp_error error) {
    puts("streaming_error_callback called");
}

void userinfo_updated_callback(sp_session *session) {
    puts("userinfo_updated_callback called");
}

void start_playback_callback(sp_session *session) {
    puts("start_playback_callback called");
}

void stop_playback_callback(sp_session *session) {
    puts("stop_playback_callback called");
}

void get_audio_buffer_stats_callback(sp_session *session, sp_audio_buffer_stats * stats) {
    puts("get_audio_buffer_stats_callback called");
}

void offline_status_updated_callback(sp_session *session) {
    puts("offline_status_updated_callback called");
}

void offline_error_callback(sp_session *session, sp_error error) {
    puts("offline_error_callback called");
}

void connectionstate_updated_callback(sp_session *session) {
    puts("connectionstate_updated_callback called");
}

void connection_error_callback(sp_session *session, sp_error error) {
    printf("connection_error called %s\n", sp_error_message(error));
}

sp_session_callbacks *init_callbacks(sp_session_callbacks *scb) {

    scb->logged_in = &logged_in_callback;
    scb->logged_out = &logged_out_callback;
    scb->metadata_updated = &metadata_updated_callback;
    scb->message_to_user = &message_to_user_callback;
    scb->notify_main_thread = &notify_main_thread_callback;
    scb->music_delivery = &music_delivery_callback;
    scb->log_message = &log_message_callback;
    scb->offline_error = &offline_error_callback;
    scb->offline_status_updated = &offline_status_updated_callback;
    scb->connectionstate_updated = &connectionstate_updated_callback;
    scb->connection_error = &connection_error_callback;
    return (scb);
}

extern void GoWaitSearch(sp_search* s);

void SP_CALLCONV search_cb(sp_search *res, void *userdata) {
    printf("LOOOOOL\n");
    GoWaitSearch(res);
}

search_complete_cb *get_search_cb() {
    return &search_cb;
}
