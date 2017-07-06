
#ifndef LEETCODE_501_HPP
#define LEETCODE_501_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <set>
#include <numeric>
#include <functional>
#include "../common/leetcode.hpp"

using namespace std;


struct pcnt {
    int v;
    int cnt;

    void reset(int v) {
        this->cnt = 1;
        this->v = v;
    }
};


class Solution {
public:
    vector<int> findMode(TreeNode *root) {
        if (root == nullptr) {
            return {};
        }
        deque<pcnt> res;

        preTravel(root, [&](int v) {
            if (res.size() == 0) {
                res.emplace_back(pcnt{v, 1});
            } else {
                if (v == res.back().v) {
                    res.back().cnt++;
                } else {
                    if (res.back().cnt < res.front().cnt) {
                        res.back().reset(v);
                    } else {
                        while (res.front().cnt < res.back().cnt) {
                            res.pop_front();
                        }
                        res.emplace_back(pcnt{v, 1});
                    }

                }
            }
        });

        if (res.back().cnt < res.front().cnt) {
            res.pop_back();
        }
        while (res.front().cnt < res.back().cnt) {
            res.pop_front();
        }

        vector<int> ans;
        for_each(res.begin(), res.end(), [&ans](pcnt p) {
            ans.emplace_back(p.v);
        });
        return ans;
    }

private:

    void preTravel(TreeNode *root, function<void(int)> f) {
        if (root != nullptr) {
            preTravel(root->left, f);
            f(root->val);
            preTravel(root->right, f);
        }
    }
};


#endif //LEETCODE_501_HPP
