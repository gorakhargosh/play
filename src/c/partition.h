#include <stdbool.h>

#ifndef _X_PARTITION_H_
#define _X_PARTITION_H_

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
  Ordinal *id;

  // Each tree set in the partition carries a weight that indicates the number
  // of elements in that set.
  Weight *weight;

  // A seen array, for the lack of a hash-table, represents whether each
  // element of the set has been "seen" in a previous union operation.
  bool *seen;

  // The capacity of the partition.
  Ordinal capacity;
} Partition;

// Internal API.

// partition_findSet1 finds the representative of the disjoint set while
// compressing the path lazily.
Ordinal partition_findSet1(Partition *p, Ordinal x);

// partition_findSet2 finds the representative of the disjoint set while
// compressing the path eagerly.
Ordinal partition_findSet2(Partition *p, Ordinal x);

// partition_FindSetRecursive finds the representative of the disjoint set while
// compressing the path eagerly and recursively.
Ordinal partition_findSetRecursive(Partition *p, Ordinal x);

// partition_printWeights displays the weights of the entire partition.
void partition_printWeights(Partition *p);

// partition_printParents displays the weights of the entire partition.
void partition_printParents(Partition *p);

// Public API.

// Partition_FindSet determines the representative element of a disjoint set.
#define Partition_FindSet partition_findSetRecursive

Partition *NewPartition(Ordinal n);

// Partition_Weight determines the weight of the set to which x belongs.
Weight Partition_Weight(Partition *p, Ordinal x);

// Partition_MinWeight determines the minimum weight in the partition.
Weight Partition_MinWeight(Partition *p);

// Partition_MaxWeight determines the maximum weight in the partition.
Weight Partition_MaxWeight(Partition *p);

// Partition_Union performs a union of the sets represented by x and y.
void Partition_Union(Partition *p, Ordinal x, Ordinal y);

// Partition_Connected determines whether the elements x and y belong to the
// same disjoint set.
bool Partition_Connected(Partition *p, Ordinal x, Ordinal y);

// Partition_Destroy destroys and deallocates the partition.
void Partition_Destroy(Partition *p);

// Partition_Capacity determines the capacity of the partition.
Ordinal Partition_Capacity(Partition *p);

#endif /* _X_PARTITION_H_ */
