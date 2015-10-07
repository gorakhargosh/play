#include <assert.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdlib.h>

#include "macros.h"
#include "partition.h"

// NewPartition creates a new partition of size n.
Partition *NewPartition(Ordinal n) {
  assert(n > 0);
  Partition *p = calloc(1, sizeof(Partition));
  pie_return_val_if_null(p, NULL);

  if (p) {
    // memset(p, 0, sizeof(Partition));
    p->capacity = n;

    p->id = calloc(n, sizeof(Ordinal));
    if (!p->id) {
      pie_free(p);
      return NULL;
    }

    p->weight = calloc(n, sizeof(Weight));
    if (!p->weight) {
      pie_free(p->id);
      pie_free(p);
      return NULL;
    }

    p->seen = calloc(n, sizeof(bool));
    if (!p->seen) {
      pie_free(p->weight);
      pie_free(p->id);
      pie_free(p);
      return NULL;
    }

    for (Ordinal i = 0; i < n; i++) {
      p->id[i] = i;
      p->weight[i] = 1;
      // p->seen[i] = false; // calloc does this automatically.
    }
  }
  return p;
}

Ordinal partition_findSet1(Partition *p, Ordinal x) {
  // Single-pass point-at-grandparent path compression.
  while (x != p->id[x]) {
    p->id[x] = p->id[p->id[x]];
    x = p->id[x];
  }
  return x;
}

Ordinal partition_findSet2(Partition *p, Ordinal x) {
  // Two-pass point-at-root path compression.
  while (x != p->id[x]) {
    x = p->id[x];
  }
  // x is now root.
  for (Ordinal i = x; i != p->id[i];) {
    i = p->id[i];
    p->id[i] = x;
  }
  return x;
}

Ordinal partition_findSetRecursive(Partition *p, Ordinal x) {
  // Two-pass recursive point-all-nodes-at-root path compression. Unwinding the
  // recursion causes all the nodes in the path to point at root. This
  // implementation can be found in CLRS Chapter 21 and is particularly clean.
  if (x != p->id[x]) {
    p->id[x] = partition_findSetRecursive(p, p->id[x]);
  }
  return p->id[x];
}

Ordinal Partition_Capacity(Partition *p) { return p->capacity; }

Weight Partition_Weight(Partition *p, Ordinal x) {
  return p->weight[Partition_FindSet(p, x)];
}

Weight Partition_MinWeight(Partition *p, bool countIndividuals) {
  Weight min_weight = 0;
  Weight weight = 0;
  for (Ordinal i = 0; i < p->capacity; i++) {
    if (p->id[i] == i && (countIndividuals || p->seen[i])) {
      // We have a root element.
      weight = p->weight[i];
      if (min_weight == 0 || weight < min_weight) {
        min_weight = weight;
      }
    }
  }
  return min_weight;
}

Weight Partition_MaxWeight(Partition *p, bool countIndividuals) {
  Weight max_weight = 0;
  Weight weight = 0;
  for (Ordinal i = 0; i < p->capacity; i++) {
    if (p->id[i] == i && (countIndividuals || p->seen[i])) {
      // We have a root element.
      weight = p->weight[i];
      if (max_weight == 0 || weight > max_weight) {
        max_weight = weight;
      }
    }
  }
  return max_weight;
}

void Partition_Union(Partition *p, Ordinal x, Ordinal y) {
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

bool Partition_Connected(Partition *p, Ordinal x, Ordinal y) {
  return Partition_FindSet(p, x) == Partition_FindSet(p, y);
}

void Partition_Destroy(Partition *p) {
  if (p) {
    pie_free(p->id);
    pie_free(p->weight);
    pie_free(p->seen);
    pie_free(p);
  }
}
