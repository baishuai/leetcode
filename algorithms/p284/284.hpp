
#ifndef LEETCODE_284_HPP
#define LEETCODE_284_HPP

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


class Iterator {
    struct Data;
    Data *data;
public:
    Iterator(const vector<int> &nums);

    Iterator(const Iterator &iter);

    virtual ~Iterator();

    // Returns the next element in the iteration.
    int next();

    // Returns true if the iteration has more elements.
    bool hasNext() const;
};


class PeekingIterator : public Iterator {
public:
    PeekingIterator(const vector<int> &nums) : Iterator(nums) {
        // Initialize any member here.
        // **DO NOT** save a copy of nums and manipulate it directly.
        // You should only use the Iterator interface methods.
    }

    // Returns the next element in the iteration without advancing the iterator.
    int peek() {
        if (pek.size() == 0)
            pek.push_back(Iterator::next());
        return pek.front();
    }

    // hasNext() and next() should behave the same as in the Iterator interface.
    // Override them if needed.
    int next() {
        int v;
        if (pek.size() == 1) {
            v = pek.front();
            pek.pop_back();
        } else {
            v = Iterator::next();
        }
        return v;
    }

    bool hasNext() const {
        return pek.size() > 0 || Iterator::hasNext();
    }

private:

    vector<int> pek;

};

#endif //LEETCODE_284_HPP
