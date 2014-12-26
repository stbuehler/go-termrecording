#ifndef _GO_LIBTSM_CALLBACK_WRAPPER_H
#define _GO_LIBTSM_CALLBACK_WRAPPER_H

/* tsm_screen_draw_cb */
extern int go_wrap_tsm_screen_draw_cb(
	uint32_t id,
	uint32_t *ch,
	size_t len,
	unsigned int width,
	unsigned int posx,
	unsigned int posy,
	struct tsm_screen_attr *attr,
	tsm_age_t age,
	void *data);

static int wrap_tsm_screen_draw_cb(
	struct tsm_screen *con,
	uint32_t id,
	const uint32_t *ch,
	size_t len,
	unsigned int width,
	unsigned int posx,
	unsigned int posy,
	const struct tsm_screen_attr *attr,
	tsm_age_t age,
	void *data) {
	(void)(con); /* unused */
	/* go doesn't know about "const", cast... */
	return go_wrap_tsm_screen_draw_cb(id, (uint32_t*)ch, len, width, posx, posy, (struct tsm_screen_attr*) attr, age, data);
}

static tsm_age_t go_tsm_screen_draw(struct tsm_screen *con, void *data) {
	return tsm_screen_draw(con, &wrap_tsm_screen_draw_cb, data);
}

/* tsm_vte_write_cb */
extern void go_wrap_tsm_vte_write_cb(char* u8, size_t len, void* data);

static void wrap_tsm_vte_write_cb(struct tsm_vte *vte, const char *u8, size_t len, void *data) {
	(void)(vte); /* unused */
	/* go doesn't know about "const", cast... */
	go_wrap_tsm_vte_write_cb((char*) u8, len, data);
}

static int go_tsm_vte_new(struct tsm_vte **out, struct tsm_screen *con, void *data) {
	return tsm_vte_new(out, con,
		wrap_tsm_vte_write_cb, data,
		NULL, NULL);
}

#endif /* _GO_LIBTSM_CALLBACK_WRAPPER_H */
