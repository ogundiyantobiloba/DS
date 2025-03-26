## Squares of a Sorted Array
> A function that returns a sorted array of squared integers and its corresponding test case.

Problem can be found [here](https://leetcode.com/problems/squares-of-a-sorted-array/description/)

#### THE FUNCTION
To return sorted squares from a sorted integer array, the two-pointer approach is used.

Two pointers (`left` and `right`) start at each end of the input array. At each step, the **absolute values** pointed to are compared. This comparison is essential since negative numbers, when squared, may become larger than some positive numbers.

Because the input array is already sorted, the largest squared values will always be at either end. Therefore, we insert values starting from the **end** of the result array, filling it backward. Doing so ensures that we correctly place the largest squared value first, followed by the smaller ones, without needing additional sorting.

Go lacks a built-in integer-based absolute-value function (`abs`). Thus, a custom `abs` function for integers was implemented explicitly to handle these absolute-value comparisons.

Pointers move inward until all elements are processed, resulting in a solution with **O(n)** time complexity.

#### THE TEST CASE
To validate the correctness of the function, the following test case is used:

- **Input**:  
  `nums = [-4, -1, 0, 3, 10]`

- **Expected Output**:  
  `[0, 1, 9, 16, 100]`

Go's built-in `reflect.DeepEqual` is used to compare the actual output to the expected output.
