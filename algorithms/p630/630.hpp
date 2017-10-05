#include <iostream>
#include <queue>
#include <algorithm>

using namespace std;

class Solution {
public:
    int scheduleCourse(vector<vector<int>> &courses) {

        sort(courses.begin(), courses.end(), [](const vector<int> &x, const vector<int> &y) -> bool {
            return x[1] < y[1];
        });

        int s = 0, c = 0;

        priority_queue<int> q;

        for (auto &b : courses) {
            int t = b[0], d = b[1];
            if (s + t <= d) {
                s += t;
                c++;
                q.push(t);
            } else if (q.size() > 0 && t < q.top()) {
                s += t - q.top();
                q.pop();
                q.push(t);
            }
        }
        return c;
    }
};

