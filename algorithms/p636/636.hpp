
#ifndef LEETCODE_636_HPP
#define LEETCODE_636_HPP

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
#include <sstream>

using namespace std;

class Line {
public:
    int fid;
    bool start;
    int time;

    Line(string s) {
        auto col1 = s.find(":", 0);
        auto col2 = s.find(":", col1 + 1);
        fid = stoi(s.substr(0, col1));
        start = s.substr(col1 + 1, col2 - (col1 + 1)) == "start";
        time = stoi(s.substr(col2 + 1));
    }
};

class Solution {
public:
    vector<int> exclusiveTime(int n, vector<string> &logs) {
        vector<int> times(n, 0);
        stack<int> ids;
        int time = 0;
        for (auto log: logs) {
            Line line(log);
            if (line.start) {
                if (!ids.empty()) {
                    times[ids.top()] += line.time - time;
                }
                ids.push(line.fid);
                time = line.time;
            } else {
                line.time++;
                times[ids.top()] += line.time - time;
                ids.pop();
                time = line.time;
            }
        }
        return times;
    }
};


#endif //LEETCODE_636_HPP
