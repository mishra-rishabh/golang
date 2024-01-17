# Memory Management In Go

* Memory Management is the process of managing computer memory, moving processes between main memory and disk to boost the overall performance of the system.
* As programs run they write objects to memory. At some point these objects should be removed when they’re not needed anymore. This process is called memory management.


## Manual Memory Management

* In a language like C the programmer will call a function such as `malloc` or `calloc` to write an object to memory
* These functions return a pointer to the location of that object in `heap` memory.
* When this object is not needed anymore, the programmer calls the free function to use this chunk of memory again.
* This method of memory management is called **explicit deallocation** and is quite powerful.
* It gives the programmer greater control over the memory in use, which allows for some types of easier optimisation, particularly in low memory environments.
* However, it leads to two types of programming errors.
    1. Calling free prematurely which creates a **dangling pointer**.
        - Dangling pointers are pointers that no longer point to valid objects in memory.
        - This is bad because the program expects a defined value to live at a pointer.
        - When this pointer is later accessed there’s no guarantee of what value exists at that location in memory.
        - There may be nothing there, or some other value entirely.
    2. Failing to free memory at all.
        - If the programmer forgets to free an object they may face a memory leak as memory fills up with more and more objects.
        - This can lead to the program slowing down or crashing if it runs out of memory.
        - Unpredictable bugs can be introduced into a program when memory has to be explicitly managed.


## Automatic Memory Management

This is why languages like Go offer **automatic dynamic memory management**, or more simply, **garbage collection**.

Languages with garbage collection offer benefits like:

* increased security
* better portability across operating systems
* less code to write
* runtime verification of code
* bounds checking of arrays

Garbage collection has a performance overhead, but it isn’t as much as is often assumed. The tradeoff is that a programmer can focus on the business logic of their program and ensure it is fit for purpose, instead of worrying about managing memory.


## Memory Model, The Go Way

Go supports **automatic memory management**, such as **automatic memory allocation** and **automatic garbage collection**, which avoids many hidden bugs.


In Go, each thread has its own stack. When we start a thread, we allocate a block of memory to be used as that thread’s stack. The problem comes when this block of memory is not enough. To overcome this we can increase the default size of stack but increasing the size of stack means that the size will increase for every thread. This will be highly inefficient if we want to use lots of threads.

Another option is to decide stack size for each thread individually. Again this will be inefficient in a setup where we have lots of threads since we need to figure out how much memory we should allocate for each stack.

What Go creators came up with is instead of giving each goroutine a fixed amount of stack memory, the Go runtime attempts to give goroutines the stack space they need on demand. This way we don’t need to think about the stack size when we are creating a thread.

A goroutine starts with a 2 kb stack size which can grow and shrink as needed. Go checks whether the amount of stack required for the function it’s about to execute is available, if not a call is made to morestack which allocates a new frame and only then the function is executed. When that function exits, its return arguments are copied back to the original stack frame, and any unneeded stack space is released.

Go allocates memory in two places: a **global heap for dynamic allocations** and a **local stack for each goroutine**. One major difference Go has compared to many garbage collected languages is that many objects are allocated directly on the program stack. Go prefers allocation on the stack. Stack allocation is cheaper because it only requires two CPU instructions: one to push onto the stack for allocation, and another to release from the stack.

The Go compiler uses a process called **escape analysis** to find objects whose lifetime is known at compile-time and allocates them on the stack rather than in garbage-collected heap memory. The basic idea is to do the work of garbage collection at compile time. The compiler tracks the scope of variables across regions of code. It uses this data to determine which variables hold to a set of checks that prove their lifetime is entirely knowable at runtime. If the variable passes these checks, the value can be allocated on the stack. If not, it is said to escape, and must be heap allocated.

Whether memory is allocated on the stack or escapes to the heap is entirely dependent on how you use the memory, not on how you declare the variable.


## Garbage Collection

* Garbage collection is a form of automatic memory management. The garbage collector attempts to reclaim memory which was allocated by the program, but is no longer referenced.
* Go’s garbage collector is a **non-generational concurrent**, **tri-color mark** and **sweep garbage collector**.
* A generational garbage collector focuses on recently allocated objects because it assumes that short lived objects, like temporary variables, are reclaimed most often.
* Go compiler prefers allocating objects on the stack, short-lived objects are often allocated on a stack instead of heap; this means that there is no need for generational GC.
* Go’s garbage collection works in two phases, the **mark phase**, and the **sweep phase**. GC uses the **tri-color algorithm** to analyze the use of memory blocks. This algorithm first **marks** objects that are still being referenced as “alive” and in the next phase (**sweep**) frees the memory of objects that are not alive.

## `new()`

* Allocate memory but no INIT.
* You will get a memory address.
* Zeroed storage.


## `make()`

* Allocate memory and INIT.
* You will get a memory address.
* Non-zeroed storage.