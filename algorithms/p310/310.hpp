//
// Created by bai on 17-6-30.
//

#ifndef LEETCODE_310_HPP
#define LEETCODE_310_HPP

/**
For a undirected graph with tree characteristics, we can choose any node as the root.
 The result graph is then a rooted tree. Among all possible rooted trees,
 those with minimum height are called minimum height trees (MHTs).
 Given such a graph, write a function to find all the MHTs and return a list of their root labels.

Format
The graph contains n nodes which are labeled from 0 to n - 1.
 You will be given the number n and a list of undirected edges (each edge is a pair of labels).

You can assume that no duplicate edges will appear in edges.
 Since all edges are undirected, [0, 1] is the same as [1, 0] and thus will not appear together in edges.

 */

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <numeric>

using namespace std;

struct Node {
    unordered_set<int> neighbor;

    bool isLeaf() const {
        return neighbor.size() == 1;
    }
};

class Solution {
public:
    vector<int> findMinHeightTrees(int num, vector<pair<int, int>> &edges) {

        vector<Node> nodes((unsigned long) num);
        for (auto p : edges) {
            nodes[p.first].neighbor.insert(p.second);
            nodes[p.second].neighbor.insert(p.first);
        }

        //find all leaves
        vector<int> outLayer;

        if (num <= 2) {
            outLayer.resize((unsigned long) num);
            iota(outLayer.begin(), outLayer.end(), 0);
            return outLayer;
        }

        vector<int> innerLayer;
        for (int i = 0; i < num; ++i) {
            if (nodes[i].isLeaf()) {
                outLayer.emplace_back(i);
            }
        }

        while (true) {
            for (int i : outLayer) {
                for (auto n : nodes[i].neighbor) {
                    nodes[n].neighbor.erase(i);
                    if (nodes[n].isLeaf()) {
                        innerLayer.emplace_back(n);
                    }
                }
            }
            if (innerLayer.empty()) {
                break;
            }
            outLayer.swap(innerLayer);
            innerLayer.clear();
        }
        return outLayer;
    }
};


#endif //LEETCODE_310_HPP
