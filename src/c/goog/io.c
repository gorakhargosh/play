#include <errno.h>
#include <stdio.h>

#include "io.h"

// goog_fgetline reads a line from fp into line upto maxlen excluding the
// newline.
// Based on ftp://ftp.eskimo.com/home/scs/cclass/week2/fgetline.c
int goog_fgetline(FILE *fp, char s[], int maxlen) {
  int nch = 0;
  int c;
  maxlen = maxlen - 1;  // leave room for '\0'.
  while ((c = getc(fp)) != EOF) {
    if (c == '\n') {
      break;
    }

    if (nch < maxlen) {
      s[nch++] = c;
    }
  }
  if (c == EOF && nch == 0) {
    return EOF;
  }
  s[nch] = '\0';
  return nch;
}
