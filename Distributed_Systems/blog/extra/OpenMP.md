# Notes on OpenMP

OpenMP is a C/C++ library that adds support for multi-threading. 
This is a translation of notes that I made on my notebook for anyone who needs a cheat sheet.

Compiling a program with support for OpenMP - ```g++ -qopenmp hello_omp.cpp```

```#pragma omp parallel```
Runs the the succeeding section of the progam amongst multiple threads, in parallel.

```export OMP_NUM_THREADS=x```
Specifying the number of threads you wish to use. Default is the number of logical cores in your processor.

If you want different threads to work on the same instance of a variable, either initialize it before the omp parallel directive or you could explicitly specify it using - ```#pragma omp parallel shared(A)``` if the variable in consideration is 'A'. 

If you want each thread to have its own instance of the variable to work on, mark it as private - ```#pragma omp parallel private(A)```. By default, variables inside a parallel scope are always private to each thread.

If you wish to run different iterations of a for loop in parallel, use - ```#pragma omp parallel for```
Or 
you could 
```
#pragma omp parallel 

//code that you want to multi thread

#pragma omp for
for(int i=0; i<n; ++i)
{
	//logic
}
```

The end of this 'parallel' for loop has an implicit barrier where threads are halted until ALL threads finish the loop.

So we know we can parallelize loops. But there are different modes. of parralelization:

1. Static
OpenMP assigns the same number of iteration chunks to each thread. The scheduling overhead is low, but this is not the optimal way of running a for loop in parallel.
2. Dynamic
Threads are assigned small iteration chunks. Whenever a thread finishes a particular chunk, it is assigned a new chink of iterations. This involves a high scheduling overhead.
3. Guided
It is similar to the dynamic mode of scheduling. In the beginning, threads are assigned chunks with large number of iterations but as they finish, they are assigned smaller chunks. The minimum chunk size is specified beforehand.

Syntax for specifying the mode of scheduling:
```#pragma omp for schedule(<mode>)```

Declaring a critical section - ```#program omp critical```
Protects a {} block of code from being executed in parallel

Declaring an atmomic operation - ```#progrma omp atomic```
Protects the succeeding instruction from being executed in parallel

Always remember, good parallel programs using minimal number of mutex locks.

Parallel reduction 
```#pragma omp parallel for reduction(operation:variable)```
You are specifying exactly how your shared variable is being changed. OpenMP will produce assembly code that is highly parallel but whose execution is predictable. (*prevents data race*)

Bonus:
```#pragma omp flush```
Enforce memory consistency

```#pragma omp ordered```
Partial loop serialization