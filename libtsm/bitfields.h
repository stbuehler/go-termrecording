#define GO_TSM_SCREEN_ATTR_BOLD      (1u << 0)
#define GO_TSM_SCREEN_ATTR_UNDERLINE (1u << 1)
#define GO_TSM_SCREEN_ATTR_INVERSE   (1u << 2)
#define GO_TSM_SCREEN_ATTR_PROTECT   (1u << 3)
#define GO_TSM_SCREEN_ATTR_BLINK     (1u << 4)

static void go_set_tsm_screen_attr_bitfields(struct tsm_screen_attr *attr, uint32_t bits) {
	attr->bold      = (0 != (bits & GO_TSM_SCREEN_ATTR_BOLD));
	attr->underline = (0 != (bits & GO_TSM_SCREEN_ATTR_UNDERLINE));
	attr->inverse   = (0 != (bits & GO_TSM_SCREEN_ATTR_INVERSE));
	attr->protect   = (0 != (bits & GO_TSM_SCREEN_ATTR_PROTECT));
	attr->blink     = (0 != (bits & GO_TSM_SCREEN_ATTR_BLINK));
}

static uint32_t go_get_tsm_screen_attr_bitfields(struct tsm_screen_attr *attr) {
	return 0
		| (attr->bold      ? GO_TSM_SCREEN_ATTR_BOLD : 0)
		| (attr->underline ? GO_TSM_SCREEN_ATTR_UNDERLINE : 0)
		| (attr->inverse   ? GO_TSM_SCREEN_ATTR_INVERSE : 0)
		| (attr->protect   ? GO_TSM_SCREEN_ATTR_PROTECT : 0)
		| (attr->blink     ? GO_TSM_SCREEN_ATTR_BLINK : 0)
		;
}
