#!/usr/bin/env python

def ackermann(m, n):
  """ackermann computes ackermann computes ackermann."""
  if m == 0:
    return n+1
  elif n == 0:
    return ackermann(m-1, 1)
  else:
    return ackermann(m-1, ackermann(m, n-1))