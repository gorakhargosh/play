#ifndef _GOOG_PARTITION_H_
#define _GOOG_PARTITION_H_

#include <stdbool.h>

#include "types.h"

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
  goog_ordinal_t *id;

  // Each tree set in the partition carries a weight that indicates the number
  // of elements in that set.
  goog_weight_t *weight;

  // A seen array, for the lack of a hash-table, represents whether each
  // element of the set has been "seen" in a previous union operation.
  bool *seen;  // TODO(yesudeep): Replace this with a hash-table or bit vector
               // when we have an implementation.

  // The capacity of the partition.
  goog_ordinal_t capacity;

} goog_partition_t;

// Internal API.

// goog_partition_find_set1 finds the representative of the disjoint set while
// compressing the path lazily.
goog_ordinal_t goog_partition_find_set1(goog_partition_t *p, goog_ordinal_t x);

// goog_partition_find_set2 finds the representative of the disjoint set while
// compressing the path eagerly.
goog_ordinal_t goog_partition_find_set2(goog_partition_t *p, goog_ordinal_t x);

// goog_partition_find_set_recursive finds the representative of the disjoint
// set while compressing the path eagerly and recursively.
goog_ordinal_t goog_partition_find_set_recursive(goog_partition_t *p,
                                                 goog_ordinal_t x);

// Public API.

// goog_partition_find_set determines the representative element of a disjoint
// set.
#define goog_partition_find_set goog_partition_find_set_recursive

goog_partition_t *goog_partition_new(goog_ordinal_t n);

// goog_partition_weight determines the weight of the set to which x belongs.
goog_weight_t goog_partition_weight(goog_partition_t *p, goog_ordinal_t x);

// goog_partition_min_weight determines the minimum weight in the partition.
goog_weight_t goog_partition_min_weight(goog_partition_t *p,
                                        bool countIndividuals);

// goog_partition_max_weight determines the maximum weight in the partition.
goog_weight_t goog_partition_max_weight(goog_partition_t *p,
                                        bool countIndividuals);

// goog_partition_union performs a union of the sets represented by x and y.
void goog_partition_union(goog_partition_t *p, goog_ordinal_t x,
                          goog_ordinal_t y);

// goog_partition_connected determines whether the elements x and y belong to
// the same disjoint set.
bool goog_partition_connected(goog_partition_t *p, goog_ordinal_t x,
                              goog_ordinal_t y);

// goog_partition_destroy destroys and deallocates the partition.
void goog_partition_destroy(goog_partition_t *p);

// goog_partition_capacity determines the capacity of the partition.
goog_ordinal_t goog_partition_capacity(goog_partition_t *p);

#endif /* _GOOG_PARTITION_H_ */
