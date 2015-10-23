#include <stdio.h>
#include <stdarg.h>
#include <stdlib.h>
#include <stddef.h>

#include "testing.h"
#include "macros.h"

goog_testing_t *goog_testing_new(void) {
  goog_testing_t *t = calloc(1, sizeof(goog_testing_t));
  goog_return_if_null(t, NULL);
  t->count = 0;
  return t;
}

int goog_testing_errorf(goog_testing_t *t, char *sfmt, ...) {
  va_list args;
  va_start(args, sfmt);
  t->count++;
  int retc = vfprintf(stderr, sfmt, args);
  va_end(args);
  return retc;
}
