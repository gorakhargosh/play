#include <stdlib.h>

#ifndef _PIE_MACROS_H_
#define _PIE_MACROS_H_

#define pie_minimum(x, y) (((x) < (y)) ? (x) : (y))
#define pie_maximum(x, y) (((x) < (y)) ? (y) : (x))

#define pie_free(p) \
  do {              \
    if (p) {        \
      free(p);      \
    }               \
  } while (0)

#define pie_return_val_if_null(ptr, val) \
  do {                                   \
    if (!p) {                            \
      return (val);                      \
    }                                    \
  } while (0)

#define pie_return_if_null(ptr) \
  do {                          \
    if (!p) {                   \
      return;                   \
    }                           \
  } while (0)

#endif /* _PIE_MACROS_H_ */
