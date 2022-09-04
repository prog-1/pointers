# Pointer excercises

1. Use an array instead of a slice to implement a function, that updates its
   elements. Answer the following questions:
   - Try comparing the behavior to other programming languages that you may know
     or heard about.
   - Are arrays references or values?
2. Try using a map instead of a slice. Are maps references or values?
3. Create a function that appends elements to a slice. The function must update
   the slice provided rather than returning a value (which would be more like a
   pure-function). Please note that pure-functions are cleaner and better, but
   in this particular case we are playing with pointers, trying to feel and
   understand them.
4. Create a `double word` variable (e.g. `uint32`), apply `[4]byte` array on
   top of it and tell what byte has a lower address - the least or the most
   significant (e.g. `0x11` is the most significant and `0x44` is the least
   significant in `0x11223344` double word).
5. You are given the following structure:

   ```go
   type Node struct {
     Value int
     Next *Node
   }
   ```

   This structure represents a linear collection of elements, where the order of
   elements in the physical memory is not defined. We use a reference to the
   next element, rather than an index like in a slice or array. E.g.

   ![singly linked list](resources/Singly-linked-list.svg.png)

   Implement functions that:

   - Prepend new values (aka pushing front).

     ```go
     func Prepend(head **Node, x int)
     ```

   - Append new values (aka pushing back).

     ```go
     func Append(head **Node, x int)
     ```

   - Insert new elements keeping lexicographical order (keep elements sorted).

     ```go
     func Insert(head **Node, x int)
     ```

   - Checks whether a value exists in the collection.

     ```go
     func Contains(head *Node, x int) bool
     ```

   - Removes given elements

     ```go
     func Remove(head **Node, x int) (removed bool)
     ```
