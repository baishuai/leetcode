
#ifndef LEETCODE_599_HPP
#define LEETCODE_599_HPP

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

You need to help them find out their common interest with the least list index sum.
 If there is a choice tie between answers, output all of them with no order requirement.
 You could assume there always exists an answer.

Example 1:
Input:
["Shogun", "Tapioca Express", "Burger King", "KFC"]
["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]
Output: ["Shogun"]
Explanation: The only restaurant they both like is "Shogun".
Example 2:
Input:
["Shogun", "Tapioca Express", "Burger King", "KFC"]
["KFC", "Shogun", "Burger King"]
Output: ["Shogun"]
Explanation: The restaurant they both like and have the least index sum is "Shogun" with index sum 1 (0+1).
 */

class Solution {
public:
    vector<string> findRestaurant(vector<string> &list1, vector<string> &list2) {
        unordered_map<string, int> andy, doris;
        for (int i = 0; i < list1.size(); ++i)
            andy[list1[i]] = i;
        for (int j = 0; j < list2.size(); ++j)
            doris[list2[j]] = j;
        vector<string> interset;
        unsigned long minIndex = list1.size() + list2.size();
        for (const auto &p: andy) {
            if (doris.count(p.first)) {
                int dex = p.second + doris[p.first];
                if (dex == minIndex)
                    interset.emplace_back(p.first);
                else if (dex < minIndex) {
                    minIndex = static_cast<unsigned long>(dex);
                    interset.clear();
                    interset.emplace_back(p.first);
                }
            }
        }
        return interset;
    }
};

#endif //LEETCODE_599_HPP
