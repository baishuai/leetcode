//
// Created by bai on 17-7-2.
//

#ifndef LEETCODE_635_HPP
#define LEETCODE_635_HPP


#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <map>
#include <unordered_set>
#include <numeric>
#include <sstream>

using namespace std;


class LogSystem {
public:
    LogSystem() {

    }

    void put(int id, string timestamp) {
        this->logs[convert(timestamp)] = id;
    }

    vector<int> retrieve(string s, string e, string gra) {
        long long l = convert(s), h = convert(e), gran = 1;

        if (gra == "Year") gran = 10000000000;
        else if (gra == "Month") gran = 100000000;
        else if (gra == "Day") gran = 1000000;
        else if (gra == "Hour") gran = 10000;
        else if (gra == "Minute") gran = 100;
        else if (gra == "Second") gran = 1;

        l = l / gran * gran;
        h = h / gran * gran;
        h += gran;

        vector<int> res;
        for_each(logs.lower_bound(l), logs.lower_bound(h), [&res](pair<long long, int> p) {
            res.push_back(p.second);
        });
        return res;
    }

private:

    map<long long, int> logs;

    long long convert(string &s) {
        long long y, m, d, H, M, S;
        char _;
        stringstream ss(s);
        ss >> y >> _ >> m >> _ >> d >> _ >> H >> _ >> M >> _ >> S;
        return ((((y * 100 + m) * 100 + d) * 100 + H) * 100 + M) * 100 + S;
    }
};


/**
 * Your LogSystem object will be instantiated and called as such:
 * LogSystem obj = new LogSystem();
 * obj.put(id,timestamp);
 * vector<int> param_2 = obj.retrieve(s,e,gra);
 */


#endif //LEETCODE_635_HPP
