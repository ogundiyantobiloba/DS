## REVERSE STRING
> A function that reverses a string and its corresponding test case.

#### THE FUNCTION
To reverse a string, the two-pointer technique is one of the most optimized solutions.  
Since strings are immutable in Go, the input and output are represented as a slice of bytes, which is mutable.

First, two pointers are defined: one is placed at the beginning of the slice and the other at the end.  
Reversing the string means reading the letters in reverse order.

A for loop controls how far the left and right pointers move before they meet. Inside the loop, the byte at the beginning pointer is swapped with the byte at the end pointer repeatedly until all elements have been exchanged.

After each swap, the pointers move one step closer to each other so that they point to the next unswapped values. When the loop finishes, the reversed slice is returned.

#### THE TEST CASE
To validate this solution, we define our `TestReverseString` function with an input:
`input := []byte{'h', 'e', 'l', 'l', 'o'}`  
and the expected output:
`want := []byte{'o', 'l', 'l', 'e', 'h'}`  
After calling the function with our input, the output should match the expected result. An algorithm should always produce the same output for the same input.

To confirm that our algorithm is correct, we use the `reflect.DeepEqual` function. This function recursively compares each element in the slices to ensure that they are equal.
