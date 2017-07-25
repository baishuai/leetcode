
#ifndef LEETCODE_535_HPP
#define LEETCODE_535_HPP

#include <iostream>
#include <vector>
#include <string>

using namespace std;


class Solution {
public:

    // Encodes a URL to a shortened URL.
    string encode(string longUrl) {
        urls.emplace_back(longUrl);
        return to_string(urls.size() - 1);
    }

    // Decodes a shortened URL to its original URL.
    string decode(string shortUrl) {
        int index = stoi(shortUrl);
        return index < urls.size() ? urls[index] : "";
    }

private:
    vector<string> urls;
};

// Your Solution object will be instantiated and called as such:
// Solution solution;
// solution.decode(solution.encode(url));


#endif //LEETCODE_535_HPP
