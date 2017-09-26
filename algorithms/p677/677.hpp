
#ifndef LEETCODE_677_HPP
#define LEETCODE_677_HPP

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

struct TrieNode {
    int val;
    unordered_map<char, TrieNode *> children;
};

class MapSum {
public:
    /** Initialize your data structure here. */
    MapSum() {
        root = new TrieNode();
    }

    void insert(string key, int val) {
        auto iter = root;
        for (char c: key) {
            if (iter->children.count(c) == 0 || iter->children[c] == nullptr) {
                iter->children[c] = new TrieNode();
            }
            iter = iter->children[c];
        }
        iter->val = val;
    }

    int sum(string prefix) {
        auto iter = root;
        for (char c: prefix) {
            if (iter->children.count(c) == 0 || iter->children[c] == nullptr) {
                return 0;
            }
            iter = iter->children[c];
        }
        return sumTail(iter);
    }

    int sumTail(TrieNode *node) {
        if (node == nullptr)
            return 0;
        int ans = node->val;
        for (auto &p : node->children) {
            ans += sumTail(p.second);
        }
        return ans;
    }

    void free(TrieNode *node) {
        if (node == nullptr)
            return;
        for (auto &p : node->children) {
            free(p.second);
        }
        delete node;
    }

    virtual ~MapSum() {
        free(root);
    }

private:
    TrieNode *root;
};

/**
 * Your MapSum object will be instantiated and called as such:
 * MapSum obj = new MapSum();
 * obj.insert(key,val);
 * int param_2 = obj.sum(prefix);
 */

#endif //LEETCODE_677_HPP
