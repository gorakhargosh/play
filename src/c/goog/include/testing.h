#ifndef _GOOG_TESTING_H_
#define _GOOG_TESTING_H_

typedef struct { int count; } goog_testing_t;

goog_testing_t *goog_testing_new(void);
int goog_testing_errorf(goog_testing_t *t, char *sfmt, ...);

#if 0
/**
 * Minimal unit test runner.
 *
 * Based on http://www.jera.com/techinfo/jtns/jtn002.html
 */
#define goog_assert(message, test) \
  do {                             \
    if (!(test)) {                 \
      return (message);            \
    }                              \
  } while (0)

#define goog_test(test)       \
  do {                        \
    char *message = (test)(); \
    goog_tests_run++;         \
    if (message) {            \
      return message;         \
    }                         \
  } while (0)

// The number of tests run.
extern int goog_tests_run;
#endif /* 0 */

#endif /* _GOOG_TESTING_H_ */
