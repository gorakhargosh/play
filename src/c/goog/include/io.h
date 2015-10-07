#ifndef _GOOG_IO_H_
#define _GOOG_IO_H_

#include <stdio.h>

/**
 * Reads a line from fp into s upto maxlen, EOF, or EOL.
 *
 * @param fp     Pointer to the file from which to read.
 * @param s      The array of characters into which to read.
 * @param maxlen The maximum length of input to read.
 */
int goog_fgetline(FILE *fp, char s[], int maxlen);

#endif /* _GOOG_IO_H_ */
