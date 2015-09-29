#!/usr/bin/env python

import sys
import time

def factorial(n):
  """factorial calculates the factorial of n."""
  if n == 0:
    return 1
  else:
    return n * factorial(n - 1)


def fibonacci(n):
  """Fibonacci calculates the nth fibonacci number."""
  if n < 2:
    return n
  return fibonacci(n - 2) + fibonacci(n - 1)


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


def A(m, n, s="%s"):
  print s % ("A(%d,%d)" % (m, n))
  if m == 0:
    return n + 1
  if n == 0:
    return A(m - 1, 1, s)
  n2 = A(m, n - 1, s % ("A(%d,%%s)" % (m - 1)))
  return A(m - 1, n2, s)


if __name__ == '__main__':
  sys.setrecursionlimit(1500)
  ack()
