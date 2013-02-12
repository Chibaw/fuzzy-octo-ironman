#ifndef _SESSION_CALLBACKS_H_
# define _SESSION_CALLBACKS_H_

#include <libspotify/api.h>

void logged_in_callback(sp_session *session, sp_error err);
sp_session_callbacks *init_callbacks(sp_session_callbacks *scb);

search_complete_cb *get_search_cb();

extern const uint8_t g_appkey[];
extern const size_t g_appkey_size;

#endif
