package com.google.basics;

public class Math {
  public static long factorial(long number) {
    if (number == 1) {
      return 1;
    } else if (number == 0) {
      return 1;
    } else {
      return number * factorial(number - 1);
    }
  }
}
