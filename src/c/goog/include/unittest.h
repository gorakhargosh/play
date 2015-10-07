#ifndef _GOOG_UNITTEST_H_
#define _GOOG_UNITTEST_H_

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

#endif /* _GOOG_UNITTEST_H_ */
