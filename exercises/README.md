# Pointer exercises

1. Use an array instead of a slice to implement a function that updates its
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
   next element, rather than an index like in a slice or array. The last element of
   the collection always points to `nil. And an empty collection is represented as a
   `nil` pointer. E.g.

   ![singly linked list](resources/Singly-linked-list.svg.png)
   
   Note: `**Node` in the code below is not a typo. It really means a pointer to a pointer to a `Node`. Try to imagine that and even draw on a paper.

   Implement functions that:

   - Prepend new values (aka pushing front).

     ```go
     func Prepend(head **Node, x int)
     ```

   - Append new values (aka pushing back).

     ```go
     func Append(head **Node, x int)
     ```

   - Insert new elements keeping lexicographical order (keep elements sorted). Assuming current sequence is [1 5 7 9], if we insert 8,
     the sequence will become [1 5 7 8 9]. The behavior of `Insert` is not determined if the original sequence was not sorted. In this
     case you should attempt to insert assuming the sequence was sorted, and you will end up having some of the elements being in the
     lexicographical order, and some not. It's OK.

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

   Example: <!-- Implementation at https://goplay.tools/snippet/bQFXwxdhTP8 -->
   
   ```go
   var head *Node
   Append(&head, 7)
   Prepend(&head, 3)
   Insert(&head, 4)
   Insert(&head, 2)
   Insert(&head, 9)
   Remove(&head, 4)
   for cur := head; cur != nil; cur = cur.Next {
      fmt.Print(cur.Value, " ")
   } // Output: 2 3 7 9   
   fmt.Println()
   ```
