#include <assert.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdlib.h>

#include "macros.h"
#include "partition.h"

goog_partition_t *goog_partition_new(goog_ord_t n) {
  assert(n > 0);
  goog_partition_t *p = calloc(1, sizeof(goog_partition_t));
  goog_return_if_null(p, NULL);

  if (p) {
    p->capacity = n;

    p->id = calloc(n, sizeof(goog_ord_t));
    if (!p->id) {
      free(p);
      return NULL;
    }

    p->weight = calloc(n, sizeof(goog_weight_t));
    if (!p->weight) {
      free(p->id);
      free(p);
      return NULL;
    }

    p->seen = calloc(n, sizeof(bool));
    if (!p->seen) {
      free(p->weight);
      free(p->id);
      free(p);
      return NULL;
    }

    for (goog_ord_t i = 0; i < n; i++) {
      p->id[i] = i;
      p->weight[i] = 1;
      // p->seen[i] = false; // calloc does this automatically.
    }
  }
  return p;
}

goog_ord_t goog_partition_find_set1(goog_partition_t *p, goog_ord_t x) {
  // Single-pass point-at-grandparent path compression.
  while (x != p->id[x]) {
    p->id[x] = p->id[p->id[x]];
    x = p->id[x];
  }
  return x;
}

goog_ord_t goog_partition_find_set2(goog_partition_t *p, goog_ord_t x) {
  // Two-pass point-at-root path compression.
  while (x != p->id[x]) {
    x = p->id[x];
  }
  // x is now root.
  for (goog_ord_t i = x; i != p->id[i];) {
    i = p->id[i];
    p->id[i] = x;
  }
  return x;
}

goog_ord_t goog_partition_find_set_recursive(goog_partition_t *p,
                                             goog_ord_t x) {
  // Two-pass recursive point-all-nodes-at-root path compression. Unwinding the
  // recursion causes all the nodes in the path to point at root. This
  // implementation can be found in CLRS Chapter 21 and is particularly clean.
  if (x != p->id[x]) {
    p->id[x] = goog_partition_find_set_recursive(p, p->id[x]);
  }
  return p->id[x];
}

goog_ord_t goog_partition_capacity(goog_partition_t *p) { return p->capacity; }

goog_weight_t goog_partition_weight(goog_partition_t *p, goog_ord_t x) {
  return p->weight[goog_partition_find_set(p, x)];
}

goog_weight_t goog_partition_min_weight(goog_partition_t *p,
                                        bool count_individuals) {
  goog_weight_t min_weight = 0;
  goog_weight_t weight = 0;
  for (goog_ord_t i = 0; i < p->capacity; i++) {
    if (p->id[i] == i && (count_individuals || p->seen[i])) {
      // We have a root element.
      weight = p->weight[i];
      if (min_weight == 0 || weight < min_weight) {
        min_weight = weight;
      }
    }
  }
  return min_weight;
}

goog_weight_t goog_partition_max_weight(goog_partition_t *p,
                                        bool count_individuals) {
  goog_weight_t max_weight = 0;
  goog_weight_t weight = 0;
  for (goog_ord_t i = 0; i < p->capacity; i++) {
    if (p->id[i] == i && (count_individuals || p->seen[i])) {
      // We have a root element.
      weight = p->weight[i];
      if (max_weight == 0 || weight > max_weight) {
        max_weight = weight;
      }
    }
  }
  return max_weight;
}

void goog_partition_union(goog_partition_t *p, goog_ord_t x, goog_ord_t y) {
  int a = goog_partition_find_set(p, x);
  int b = goog_partition_find_set(p, y);

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

bool goog_partition_connected(goog_partition_t *p, goog_ord_t x, goog_ord_t y) {
  return goog_partition_find_set(p, x) == goog_partition_find_set(p, y);
}

void goog_partition_destroy(goog_partition_t *p) {
  if (p) {
    goog_free(p->id);
    goog_free(p->weight);
    goog_free(p->seen);
    free(p);
  }
}
