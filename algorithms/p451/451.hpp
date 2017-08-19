
#ifndef LEETCODE_451_HPP
#define LEETCODE_451_HPP

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

//Given a string, sort it in decreasing order based on the frequency of characters.

class Solution {
public:
    string frequencySort(string s) {
        unordered_map<char, int> freq;
        for (char c: s)
            freq[c]++;

        auto comp = [](const pair<char, int> &l, const pair<char, int> &r) -> bool {
            return l.second < r.second;
        };

        priority_queue<pair<char, int>, vector<pair<char, int>>, decltype(comp)> pq(comp);
        for (auto &p : freq)
            pq.push(p);

        string res;
        while (!pq.empty()) {
            auto t = pq.top();
            pq.pop();
            for (int i = 0; i < t.second; ++i) {
                res.push_back(t.first);
            }
        }
        return res;

    }
};


#endif //LEETCODE_451_HPP
