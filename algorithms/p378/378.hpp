
#ifndef LEETCODE_378_HPP
#define LEETCODE_378_HPP

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
Given a n x n matrix where each of the rows and columns are sorted in ascending order,
 find the kth smallest element in the matrix.

Note that it is the kth smallest element in the sorted order, not the kth distinct element.

Example:

matrix = [
   [ 1,  5,  9],
   [10, 11, 13],
   [12, 13, 15]
],
k = 8,

return 13.
Note:
You may assume k is always valid, 1 < k < n2.
 */

struct Tuple {
    int x;
    int y;
    int val;

    Tuple(int x, int y, int val) : x(x), y(y), val(val) {};

    friend bool operator<(const Tuple &lhs, const Tuple &rhs) {
        return lhs.val < rhs.val;
    }

    friend bool operator>(const Tuple &lhs, const Tuple &rhs) {
        return rhs < lhs;
    }
};

class Solution {
public:
    int kthSmallest(vector<vector<int>> &matrix, int k) {
        priority_queue<Tuple, vector<Tuple>, greater<Tuple>> pq;
        for (int i = 0; i < matrix.size(); ++i) {
            pq.push(Tuple(i, 0, matrix[i][0]));
        }
        while (--k > 0) {
            auto t = pq.top();
            pq.pop();
            if (t.y + 1 < matrix.size())
                pq.push(Tuple(t.x, t.y + 1, matrix[t.x][t.y + 1]));
        }
        return pq.top().val;
    }
};


#endif //LEETCODE_378_HPP
