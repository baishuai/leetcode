
#ifndef LEETCODE_133_HPP
#define LEETCODE_133_HPP

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


//Definition for undirected graph.
struct UndirectedGraphNode {
    int label;
    vector<UndirectedGraphNode *> neighbors;

    UndirectedGraphNode(int x) : label(x) {};
};

class Solution {
public:
    UndirectedGraphNode *cloneGraph(UndirectedGraphNode *node) {
        return clone(node);
    }

private:
    unordered_map<UndirectedGraphNode *, UndirectedGraphNode *> gmap;

    UndirectedGraphNode *clone(UndirectedGraphNode *node) {
        if (node == nullptr)
            return nullptr;
        if (gmap.count(node) > 0)
            return gmap[node];
        UndirectedGraphNode *c = new UndirectedGraphNode(node->label);
        gmap[node] = c;
        for (UndirectedGraphNode *n : node->neighbors) {
            c->neighbors.push_back(clone(n));
        }
        return c;
    }
};


#endif //LEETCODE_133_HPP
