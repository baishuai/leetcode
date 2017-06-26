//
// Created by bai on 17-6-26.
//

#ifndef LEETCODE_313_HPP
#define LEETCODE_313_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <string>

using namespace std;

/**
 * Write a program to find the nth super ugly number.

Super ugly numbers are positive numbers whose all prime factors are in the given prime list primes of size k.
 For example, [1, 2, 4, 7, 8, 13, 14, 16, 19, 26, 28, 32] is the sequence of the
 first 12 super ugly numbers given primes = [2, 7, 13, 19] of size 4.

Note:
(1) 1 is a super ugly number for any given primes.
(2) The given numbers in primes are in ascending order.
(3) 0 < k ≤ 100, 0 < n ≤ 10^6, 0 < primes[i] < 1000.
(4) The nth super ugly number is guaranteed to fit in a 32-bit signed integer.
 */

class ugly {

public:
    int value;
    int qid;

    ugly(int value, int qid) : value(value), qid(qid) {}

    bool operator<(const ugly &rhs) const {
        return value > rhs.value;
    }
};

class Solution {
public:
    int nthSuperUglyNumber(int n, vector<int> &primes) {
        if (n == 1) {
            return 1;
        }
        n--;
        priority_queue<ugly, vector<ugly>, less<ugly>> uglypq;
        vector<queue<int>> waitq = vector<queue<int>>(primes.size());

        for (int i = 0; i < primes.size(); i++) {
            uglypq.emplace(ugly(primes[i], i));
        }
        int nthUgly = 1;
        while (n-- > 0) {
            nthUgly = uglypq.top().value;
            int qid = uglypq.top().qid;
            uglypq.pop();

            for (int i = qid; i < primes.size(); ++i) {
                waitq[i].emplace(nthUgly * primes[i]);
            }
            uglypq.emplace(ugly(waitq[qid].front(), qid));
            waitq[qid].pop();
        }
        return nthUgly;
    }

};


#endif //LEETCODE_313_HPP
