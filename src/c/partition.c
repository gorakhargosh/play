#include <assert.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>

#define DEBUG 0

#define minimum(x, y) (((x) < (y)) ? (x) : (y))
#define maximum(x, y) (((x) < (y)) ? (y) : (x))

#define BUF_SIZ 512

/**
 * The weight of each tree in the forest needs to have at most 2^64 unique
 * values.
 */
typedef unsigned long long weight_t;

/**
 * The ordinal of each node.
 */
typedef unsigned long long ordinal_t;

/**
 * A forest of disjoint set trees represents the partition. A connection between
 * two nodes marks one as the parent of another. See Sedgewick and CLRS for more
 * information about the forest representation.
 *
 * Three variations of path compression are available:
 *
 * 1. Lazy point-node-at-grand-parent.
 * 2. Eager point-all-nodes-at-root.
 * 3. Eager & recursive point-all-nodes-at-root.
 */
typedef struct {
  // Each element is represented by its ordinal and the set representative of
  // each element is the root of the set tree.
  ordinal_t *id;

  // Each tree set in the partition carries a weight that indicates the number
  // of elements in that set.
  weight_t *weight;

  // A seen array, for the lack of a hash-table, represents whether each
  // element of the set has been "seen" in a previous union operation.
  bool *seen;

  // The capacity of the partition.
  ordinal_t capacity;
} partition_t;

#define Partition_FindSet partition_FindSetRecursive

// NewPartition creates a new partition of size n.
partition_t *NewPartition(ordinal_t n) {
  assert(n > 0);
  partition_t *p = malloc(sizeof(partition_t));
  if (NULL != p) {
    memset(p, 0, sizeof(partition_t));
    p->capacity = n;

    p->id = malloc(n * sizeof(ordinal_t));
    if (NULL == p->id) {
      free(p);
      return NULL;
    }

    p->weight = malloc(n * sizeof(weight_t));
    if (NULL == p->weight) {
      free(p->id);
      free(p);
      return NULL;
    }

    p->seen = malloc(n * sizeof(bool));
    if (NULL == p->seen) {
      free(p->weight);
      free(p->id);
      free(p);
      return NULL;
    }

    for (ordinal_t i = 0; i < n; i++) {
      p->id[i] = i;
      p->weight[i] = 1;
      p->seen[i] = false;
    }
  }
  return p;
}

// partition_FindSet1 finds the representative of the disjoint set while
// compressing the path lazily.
ordinal_t partition_FindSet1(partition_t *p, ordinal_t x) {
  // Single-pass point-at-grandparent path compression.
  while (x != p->id[x]) {
    p->id[x] = p->id[p->id[x]];
    x = p->id[x];
  }
  return x;
}

// partition_FindSet2 finds the representative of the disjoint set while
// compressing the path eagerly.
ordinal_t partition_FindSet2(partition_t *p, ordinal_t x) {
  // Two-pass point-at-root path compression.
  while (x != p->id[x]) {
    x = p->id[x];
  }
  // x is now root.
  for (ordinal_t i = x; i != p->id[i];) {
    i = p->id[i];
    p->id[i] = x;
  }
  return x;
}

// partition_FindSetRecursive finds the representative of the disjoint set while
// compressing the path eagerly and recursively.
ordinal_t partition_FindSetRecursive(partition_t *p, ordinal_t x) {
  // Two-pass recursive point-all-nodes-at-root path compression. Unwinding the
  // recursion causes all the nodes in the path to point at root. This
  // implementation can be found in CLRS Chapter 21 and is particularly clean.
  if (x != p->id[x]) {
    p->id[x] = partition_FindSetRecursive(p, p->id[x]);
  }
  return p->id[x];
}

// Partition_Weight determines the weight of the set to which x belongs.
weight_t Partition_Weight(partition_t *p, ordinal_t x) {
  return p->weight[Partition_FindSet(p, x)];
}

// Partition_MinWeight determines the minimum weight of the partition.
weight_t Partition_MinWeight(partition_t *p) {
  weight_t min_weight = 0;
  weight_t weight = 0;
  for (ordinal_t i = 0; i < p->capacity; i++) {
    if (p->seen[i] && p->id[i] == i) {
      // We have a root element.
      weight = p->weight[i];
      if (min_weight == 0 || weight < min_weight) {
        min_weight = weight;
      }
    }
  }
  return min_weight;
}

// Partition_MaxWeight determines the maximum weight of the partition.
weight_t Partition_MaxWeight(partition_t *p) {
  weight_t max_weight = 0;
  weight_t weight = 0;
  for (ordinal_t i = 0; i < p->capacity; i++) {
    if (p->seen[i] && p->id[i] == i) {
      // We have a root element.
      weight = p->weight[i];
      if (max_weight == 0 || weight > max_weight) {
        max_weight = weight;
      }
    }
  }
  return max_weight;
}

// Partition_Union performs a union of the sets represented by x and y.
void Partition_Union(partition_t *p, ordinal_t x, ordinal_t y) {
  int a = Partition_FindSet(p, x);
  int b = Partition_FindSet(p, y);

  p->seen[a] = true;
  p->seen[b] = true;
  if (p->weight[a] < p->weight[b]) {
    p->id[a] = b;
    p->weight[b] += p->weight[a];
  } else {
    p->id[b] = a;
    p->weight[a] += p->weight[b];
  }
}

// Partition_Connected determines whether the elements x and y belong to the
// same disjoint set.
bool Partition_Connected(partition_t *p, ordinal_t x, ordinal_t y) {
  return Partition_FindSet(p, x) == Partition_FindSet(p, y);
}

// Partition_PrintWeights displays the weights of the entire partition.
void Partition_PrintWeights(partition_t *p) {
  printf("[");
  for (ordinal_t i = 0; i < p->capacity; i++) {
    if (p->seen[i]) {
      printf("%llu, ", Partition_Weight(p, i));
    } else {
      printf("- ");
    }
  }
  printf("]\n");
}

// Partition_PrintParents displays the weights of the entire partition.
void Partition_PrintParents(partition_t *p) {
  printf("[");
  for (ordinal_t i = 0; i < p->capacity; i++) {
    printf("%llu, ", p->id[i]);
  }
  printf("]\n");
}

// Partition_Destroy destroys and deallocates the partition.
void Partition_Destroy(partition_t *p) {
  free(p->id);
  free(p->weight);
  free(p->seen);
  free(p);
}

void Partition_Test() {
  partition_t *p = NewPartition(10);
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
  fflush(stdout);
  Partition_Union(p, 5, 0);
  Partition_Union(p, 7, 2);
  Partition_Union(p, 6, 1);
  Partition_Union(p, 1, 0);
  printf("(0, 7) == true: %d\n", Partition_Connected(p, 0, 7));
  fflush(stdout);
  Partition_Destroy(p);
}

// fgetline reads a line from fp into line upto maxlen excluding the newline.
// ftp://ftp.eskimo.com/home/scs/cclass/week2/fgetline.c
int fgetline(FILE *fp, char s[], int maxlen) {
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

int main(int argc, char *argv[]) {
  // Leave some space for the trailing '\0'.
  char buf[BUF_SIZ + 1];
  FILE *fp = stdin;
  if (argc > 1) {
    fp = fopen(argv[1], "rb");
  }

  ordinal_t n = 0;
  fgetline(fp, buf, BUF_SIZ);
  sscanf(buf, "%llu", &n);

  partition_t *p = NewPartition(2 * n);
  for (ordinal_t i = 0, j = 0; n > 0; n--) {
    fgetline(fp, buf, BUF_SIZ);
    sscanf(buf, "%llu %llu", &i, &j);
    // The problem domain uses 1-based indexing. We translate each index within
    // the partition to index - 1.
    assert(i > 0 && j > 0);
    if (!Partition_Connected(p, i - 1, j - 1)) {
      Partition_Union(p, i - 1, j - 1);
    }
  }
  printf("%llu %llu\n", Partition_MinWeight(p), Partition_MaxWeight(p));
  Partition_Destroy(p);
  return 0;
}
