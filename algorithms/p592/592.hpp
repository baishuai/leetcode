
#ifndef LEETCODE_592_HPP
#define LEETCODE_592_HPP

#include <iostream>
#include <string>
#include <sstream>

using namespace std;

class Solution
{
  public:
    string fractionAddition(string expression)
    {
        istringstream ins(expression);
        int A = 0, B = 1, a, b;
        char _;
        while (ins >> a >> _ >> b)
        {
            A = A * b + a * B;
            B *= b;
            int g = abs(gcd(A, B));
            A /= g;
            B /= g;
        }
        return to_string(A) + '/' + to_string(B);
    }

    int gcd(int a, int b) {
        while (a != b) {
            if(a>b) {    
                a = a - b;    
            } else {    
                b = b - a;    
            }                      
        }
        return a;                
    }
};

#endif // LEETCODE_592_HPP