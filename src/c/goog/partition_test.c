#include <stddef.h>
#include <stdio.h>

#include <goog/partition.h>

void Partition_Test() {
  Partition *p = NewPartition(10);
  if (NULL != p) {
    for (int i = 0; i < 10; i++) {
      printf("%llu %llu\n", p->id[i], p->weight[i]);
    }
  }
  Partition_Union(p, 4, 3);
  Partition_Union(p, 3, 8);
  Partition_Union(p, 6, 5);
  Partition_Union(p, 9, 4);
  Partition_Union(p, 2, 1);
  printf("(0, 7) == false: %d\n", Partition_Connected(p, 0, 7));
  printf("(8, 9) == true: %d\n", Partition_Connected(p, 8, 9));
  Partition_Union(p, 5, 0);
  Partition_Union(p, 7, 2);
  Partition_Union(p, 6, 1);
  Partition_Union(p, 1, 0);
  printf("(0, 7) == true: %d\n", Partition_Connected(p, 0, 7));
  Partition_Destroy(p);
}
