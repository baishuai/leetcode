
#ifndef LEETCODE_331_HPP
#define LEETCODE_331_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <iterator>
#include <unordered_set>
#include <set>
#include <numeric>
#include <stack>
#include <sstream>
#include <string>

using namespace std;

class Solution {
public:
    bool isValidSerialization(string preorder) {
        stringstream ss(preorder);
        string item;
        vector<string> tokens;
        while (getline(ss, item, ',')) {
            tokens.push_back(item);
        }

        stack<int> st;

        for (int i = 0; i < tokens.size(); i++) {
            string t = tokens[i];
            if (t == "#") {
                if (st.size() == 0) {
                    return i == 0 && tokens.size() == 1;
                } else if (st.top() == 0) {
                    st.top()++;
                } else {
                    st.top()++;
                    while (st.size() > 0 && st.top() == 2) {
                        st.pop();
                    }
                }
            } else {
                if (st.size() > 0) {
                    st.top()++;
                } else if (i > 0) {
                    return false;
                }
                st.push(0);
            }
        }
        return st.size() == 0;
    }
};


#endif //LEETCODE_331_HPP
