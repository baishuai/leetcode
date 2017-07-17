
#ifndef LEETCODE_332_HPP
#define LEETCODE_332_HPP

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
Given a list of airline tickets represented by pairs of departure and arrival airports [from, to],
 reconstruct the itinerary in order. All of the tickets belong to a man who departs from JFK.
 Thus, the itinerary must begin with JFK.

Note:
If there are multiple valid itineraries, you should return the itinerary that has
 the smallest lexical order when read as a single string.
 For example, the itinerary ["JFK", "LGA"] has a smaller lexical order than ["JFK", "LGB"].

All airports are represented by three capital letters (IATA code).
You may assume all tickets form at least one valid itinerary.
Example 1:
tickets = [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
Return ["JFK", "MUC", "LHR", "SFO", "SJC"].
Example 2:
tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
Return ["JFK","ATL","JFK","SFO","ATL","SFO"].
Another possible reconstruction is ["JFK","SFO","ATL","JFK","ATL","SFO"]. But it is larger in lexical order.
 */

class Solution {
public:
    vector<string> findItinerary(vector<pair<string, string>> tickets) {
        unordered_map<string, multiset<string>> pairs;
        for (auto &p: tickets) {
            pairs[p.first].insert(p.second);
        }

        vector<string> route;
        stack<string> st;
        st.emplace("JFK");

        while (!st.empty()) {
            while (pairs[st.top()].size() > 0) {
                string next = *pairs[st.top()].begin();
                pairs[st.top()].erase(pairs[st.top()].begin());
                st.emplace(next);
            }
            route.emplace_back(st.top());
            st.pop();
        }
        return vector<string>(route.rbegin(), route.rend());
    }
};

#endif //LEETCODE_332_HPP
