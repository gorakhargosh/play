#!/usr/bin/env python
# -*- coding: utf-8 -*-

import ackermann
import timeit

def ack():
  for m in range(0, 10):
    for n in range(0, 10):
      print('ackermann(%d, %d) = %d' %
              (m, n, timeit.timeit('ackermann.ackermann(%d, %d)' % (m, n), number=1)))


if __name__ == '__main__':
  ack()
