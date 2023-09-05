// Package slice is a package of slice interfaces to handle common list-like operations.
//
// Slice contains a single Slice struct that exposes methods to perform traversal and mutation operations
// for a collection of Go interfaces. The Slice struct can be extended to handle
// the acceptance and selection of interface specific types. To extend the Slice an interface
// can be defined that calls the exposed Slice methods.
//
// Package slice comes with all Go primative types as interfaces out of the box.
//
// Each slice interface comes with a constructor function that takes zero to n arguments of the
// slice interface type.
//
// The slice interfaces to not expose the underlying interface slice to prevent a dirty reference.
// This pattern should be adopted when wrapping the Slice struct.
package slice
