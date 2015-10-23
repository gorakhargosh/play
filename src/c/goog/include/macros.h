#ifndef _GOOG_MACROS_H_
#define _GOOG_MACROS_H_

#include <stdlib.h>

/**
 * Determines the minimum of two values.
 *
 * @param x The first value.
 * @param y The second value.
 */
#define goog_min(x, y) (((x) < (y)) ? (x) : (y))

/**
 * Determines the maximum of two values.
 *
 * @param x The first value.
 * @param y The second value.
 */
#define goog_max(x, y) (((x) < (y)) ? (y) : (x))

/**
 * Checks a pointer before calling free on it.
 *
 * @param p The pointer to the memory to free.
 */
#define goog_free(p) \
  do {               \
    if (p) {         \
      free(p);       \
    }                \
  } while (0)

/**
 * Returns value if the condition is true.
 *
 * @param cond The condition to test.
 * @param val  The value to return.
 */
#define goog_return_if(cond, val) \
  do {                            \
    if ((cond)) {                 \
      return (val);               \
    }                             \
  } while (0)

/**
 * Returns value if the pointer is NULL.
 *
 * @param ptr The pointer to check.
 * @param val The value to return.
 */
#define goog_return_if_null(ptr, val) \
  do {                                \
    if (!(ptr)) {                     \
      return (val);                   \
    }                                 \
  } while (0)

/**
 * Bail by returning from the function if the pointer is NULL.
 *
 * @param cond The condition to test.
 */
#define goog_bail_if_null(ptr) \
  do {                         \
    if (!(ptr)) {              \
      return;                  \
    }                          \
  } while (0)

/**
 * Bail by returning from the function if the condition is satisfied.
 *
 * @param cond The condition to test.
 */
#define goog_bail_if(cond) \
  do {                     \
    if (cond) {            \
      return;              \
    }                      \
  } while (0)

#endif /* _GOOG_MACROS_H_ */
