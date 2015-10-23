#include <stddef.h>
#include <stdio.h>

#include "unittest.h"
#include "partition.h"

static char* test_partition_connectivity(t* goog_testing_t) {
  goog_partition_t* p = goog_partition_new(10);
  if (NULL != p) {
    for (int i = 0; i < 10; i++) {
      printf("%llu %llu\n", p->id[i], p->weight[i]);
    }
  }
  goog_partition_union(p, 4, 3);
  goog_partition_union(p, 3, 8);
  goog_partition_union(p, 6, 5);
  goog_partition_union(p, 9, 4);
  goog_partition_union(p, 2, 1);
  printf("(0, 7) == false: %d\n", goog_partition_connected(p, 0, 7));
  printf("(8, 9) == true: %d\n", goog_partition_connected(p, 8, 9));
  goog_partition_union(p, 5, 0);
  goog_partition_union(p, 7, 2);
  goog_partition_union(p, 6, 1);
  goog_partition_union(p, 1, 0);
  printf("(0, 7) == true: %d\n", goog_partition_connected(p, 0, 7));
  goog_partition_destroy(p);
}
