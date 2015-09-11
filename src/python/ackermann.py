#!/usr/bin/env python

import time

def ackermann(m, n):
  """ackermann computes ackermann computes ackermann."""
  if m == 0:
    return n+1
  elif n == 0:
    return ackermann(m-1, 1)
  else:
    return ackermann(m-1, ackermann(m, n-1))


def ack():
  for m in range(0, 10):
    for n in range(0, 10):
      start = time.time()
      ret = ackermann(m, n)
      end = time.time()
      print('ackermann(%d,%d) = %d (%s)' % (m, n, ret, (end - start) * 1000.0))


if __name__ == '__main__':
  ack()
