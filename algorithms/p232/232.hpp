
#ifndef LEETCODE_232_HPP
#define LEETCODE_232_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <set>
#include <numeric>
#include <stack>
#include <string>

using namespace std;

/*
Implement the following operations of a queue using stacks.

push(x) -- Push element x to the back of queue.
pop() -- Removes the element from in front of queue.
peek() -- Get the front element.
empty() -- Return whether the queue is empty.
Notes:
You must use only standard operations of a stack
 -- which means only push to top, peek/pop from top, size, and is empty operations are valid.
Depending on your language, stack may not be supported natively.
 You may simulate a stack by using a list or deque (double-ended queue),
 as long as you use only standard operations of a stack.
You may assume that all operations are valid
 (for example, no pop or peek operations will be called on an empty queue).

 */

class MyQueue {
public:
    /** Initialize your data structure here. */
    MyQueue() {}

    /** Push element x to the back of queue. */
    void push(int x) {
        tail.push(x);
    }

    /** Removes the element from in front of queue and returns that element. */
    int pop() {
        if (head.empty())
            transformStack();
        int v = head.top();
        head.pop();
        return v;
    }

    /** Get the front element. */
    int peek() {
        if (head.empty())
            transformStack();
        return head.top();
    }

    /** Returns whether the queue is empty. */
    bool empty() {
        return head.empty() && tail.empty();
    }

private:
    stack<int> tail;
    stack<int> head;

    void transformStack() {
        if (!head.empty())
            return;
        while (!tail.empty()) {
            head.push(move(tail.top()));
            tail.pop();
        }
    }
};


#endif //LEETCODE_232_HPP
