# GoStructs
GoStructs — A library that provides basic data structures and their rare variations made with generics.   
**Installation**: ```go get github.com/mosgor/GoStructs```

There are such data structures: 
1) [Heap](#heap)
    - [MaxHeap](#maxheap)
    - [MinHeap](#minheap)
2) [Queue](#queue)
## Heap
Tree-based sorted data structure.
Every operation takes about $O(log{n})$.  
Every slice can be transformed to the Heap with $O(n)$ complexity.  
It's based on slices, so such functions as len or cap are also **working** with it.   
   
Every Heap type supports all int types, all uint types (including uintptr), floats, and strings.
### MaxHeap
Import path: ```github.com/mosgor/GoStructs/pkg/maxHeap```  
Heap sorted from the biggest number to the smallest.  
There are such methods implemented in MaxHeap:
- **Insert.** Inserts an element into a Heap. Takes this element as an argument. Returns error value if the receiver is nil.
- **Extract.** Extracts the biggest element of the Heap and returns it and the true bool value.    
When we try to extract from an empty Heap or when the receiver is nil,
we get the default value of a type contained in the Heap and the false bool value.
- **Max.** Returns the biggest element of the Heap and the true bool value.   
When this method is used on an empty Heap or when the receiver is nil,
  we get the default value of a type contained in the Heap and the false bool value.

There is only one function:
- **Heapify.** It takes slice as an argument and returns new Heap consisting of slice elements.
### MinHeap
Import path: ```github.com/mosgor/GoStructs/pkg/minHeap```  
Heap sorted from the smallest number to the biggest.  
There are such methods implemented in MinHeap:
- **Insert.** Inserts an element into a Heap. Takes this element as an argument. Returns error value if the receiver is nil.
- **Extract.**
  Extracts the smallest element of the Heap and returns it and the true bool value.    
  When we try to extract from an empty Heap or when the receiver is nil,
  we get the default value of a type contained in the Heap and the false bool value.
- **Max.** Returns the smallest element of the Heap and the true bool value.   
  When this method is used on an empty Heap or when the receiver is nil,
  we get the default value of a type contained in the Heap and the false bool value.

There is only one function:
- **Heapify.** It takes slice as an argument and returns new Heap consisting of slice elements.

## Queue
Import path: ```github.com/mosgor/GoStructs/pkg/queue```   
Queue is a linked list type data structure following the FIFO (First In - First Out) principle.
It allows adding an element to the end, and extracting an element from the beginning.
Every method takes $O(1)$.  
The only way to get a random element from a queue is to iterate through it with a cycle.
So it takes $O(n)$.
That's why if you need to get random elements often, it is better to use non-list structures.   
Every version of a queue can take any type.

Methods of basic queue type:
- **Push**. Add an element at the end of the queue. Takes an element value as an argument and return nothing.
- **Pop**.
  Extracts the first element of the queue and returns it. 
Takes nothing.
  Also returns boolean value indicating whether the operation was successful.
  Returns default value of a type and false if the queue is empty.
- **Front**.
  Returns the first element of the queue. 
  Takes nothing.
  Also returns boolean value indicating whether the operation was successful.
  Returns default value of a type and false if the queue is empty.
- **Clear**. Makes the queue empty. Takes nothing and returns nothing.
- **Size**. Returns the size of the queue.
