
#ifndef LEETCODE_729_HPP
#define LEETCODE_729_HPP

#include <iostream>
#include <algorithm>
#include <map>

using namespace std;

class MyCalendar {
public:
    MyCalendar() = default;

    bool book(int start, int end) {
        auto next = books.lower_bound(start);
        if (next != books.end() && next->first < end)
            return false;
        if (next != books.begin() && (--next)->second > start)
            return false;
        books.insert({start, end});
        return true;
    }

private:
    map<int, int> books;

};


#endif //LEETCODE_729_HPP
