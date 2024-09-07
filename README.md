# GoStructs
GoStructs - A library that provides basic data structures and their rare variations.

There are such data structures: 
1) [Heap](#heap)
    - [MaxHeap](#maxheap)
    - [MinHeap](#minheap)
## Heap

---
Tree-based sorted data structure. Every operation takes about $O(log{n})$.  
Every slice can be transformed to the Heap with $O(n)$ complexity.  
It's based on slices, so such functions as len or cap are also **working** with it.

### MaxHeap
Heap sorted from the biggest number to the smallest.  
There are such methods implemented in MaxHeap:
- **Insert.** Inserts an element into a Heap. Takes this element as an argument. Returns error value.
- **Extract.** Extracts the biggest element of the Heap and returns it and the true bool value.    
When we try to extract from an empty Heap, we get the default value of a type contained in the Heap and the false bool value.
- **Max.** Returns the biggest element of the Heap and the true bool value.   
When this method is used on an empty Heap, we get the default value of a type contained in the Heap and the false bool value.

There is only one function:
- **Heapify.** It takes slice as an argument and returns new Heap consisting of slice elements.

### MinHeap
