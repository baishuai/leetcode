
#ifndef LEETCODE_399_HPP
#define LEETCODE_399_HPP

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


class Solution {
public:
    vector<double>
    calcEquation(vector<pair<string, string>> equations, vector<double> &values, vector<pair<string, string>> queries) {
        unordered_map<string, unordered_map<string, double >> quot;
        for (int i = 0; i < values.size(); ++i) {
            auto &p = equations[i];
            quot[p.first][p.first] = 1;
            quot[p.second][p.second] = 1;
            quot[p.first][p.second] = values[i];
            quot[p.second][p.first] = 1 / values[i];
        }


        for (auto &k : quot) {
            for (auto &i : k.second) {
                for (auto &j: k.second) {
                    if (i != j) {
                        quot[i.first][j.first] = quot[i.first][k.first] * quot[k.first][j.first];
                    }
                }
            }
        }
        vector<double> res;
        for (auto &q : queries) {
            res.push_back(quot[q.first].count(q.second) ? quot[q.first][q.second] : -1);
        }
        return res;
    }

};


#endif //LEETCODE_399_HPP
