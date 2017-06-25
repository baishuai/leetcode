#include <iostream>
#include <queue>
#include <algorithm>
#include <string>

using namespace std;

struct rect {
    int top;
    int left;
    int bottom;
    int right;
};

class Excel {
private:
    vector<vector<int>> data;
    vector<vector<vector<rect>>> deps;

    int calc(int r, int c) {
        int sum = 0;
        for (rect r : this->deps[r][c]) {
            for (int i = r.top; i <= r.bottom; ++i) {
                for (int j = r.left; j <= r.right; ++j) {
                    sum += this->get(i + 1, j + 'A');
                }
            }
        }
        return sum;
    }

public:
    Excel(int H, char W)
            : data(vector<vector<int>>(H, vector<int>(W - 'A' + 1))),
              deps(vector<vector<vector<rect>>>(H, vector<vector<rect>>(W - 'A' + 1))) {

    }

    void set(int r, char c, int v) {
        r -= 1;
        c -= 'A';
        this->data[r][c] = v;
        this->deps[r][c].clear();
    }

    int get(int r, char c) {
        r -= 1;
        c -= 'A';
        if (this->deps[r][c].size() == 0) {
            return this->data[r][c];
        }
        return calc(r, c);
    }

    int sum(int r, char c, vector<string> strs) {
        r -= 1;
        c -= 'A';
        this->deps[r][c].clear();
        for (auto &s : strs) {
            int t, b;
            char left, right;
            if (s.find(":") != string::npos) {
                sscanf(s.c_str(), "%c%d:%c%d", &left, &t, &right, &b);
            } else {
                sscanf(s.c_str(), "%c%d", &left, &t);
                right = left;
                b = t;
            }
            rect dr{t - 1, left - 'A', b - 1, right - 'A'};
            this->deps[r][c].emplace_back(dr);
        }

        return calc(r, c);
    }
};

/**
 * Your Excel object will be instantiated and called as such:
 * Excel obj = new Excel(H, W);
 * obj.set(r,c,v);
 * int param_2 = obj.get(r,c);
 * int param_3 = obj.sum(r,c,strs);
 */