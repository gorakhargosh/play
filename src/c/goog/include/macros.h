#include <stdlib.h>

#ifndef _GOOG_MACROS_H_
#define _GOOG_MACROS_H_

#define goog_minimum(x, y) (((x) < (y)) ? (x) : (y))
#define goog_maximum(x, y) (((x) < (y)) ? (y) : (x))

#define goog_free(p) \
  do {               \
    if (p) {         \
      free(p);       \
    }                \
  } while (0)

#define goog_return_val_if_null(ptr, val) \
  do {                                    \
    if (!p) {                             \
      return (val);                       \
    }                                     \
  } while (0)

#define goog_return_if_null(ptr) \
  do {                           \
    if (!p) {                    \
      return;                    \
    }                            \
  } while (0)

#endif /* _GOOG_MACROS_H_ */
