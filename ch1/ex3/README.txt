Exercise 1.3: Experiment to measure the difference in running time
between our potentially inefficient versions and the one that uses 
strings.Join. (Section 1.6 illustrates part of the time package, and Section 
11.4 shows how to write benchmark tests for systematic performance evaluation.)

For now, I've implemented both 'our potentially inefficient version' and the
given example that uses `strings.Join`.  I used the following commands to test
efficiency:

time ./join `yes "Hello World " | head -n 1000`
time ./inefficient `yes "Hello World " | head -n 1000`

inefficient.go
--------------

real	0m0.182s
user	0m0.145s
sys	    0m0.135s

join.go
-------

real	0m0.012s
user	0m0.005s
sys	    0m0.007s