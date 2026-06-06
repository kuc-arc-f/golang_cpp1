#pragma once
#include <iostream>

/*
class Calculator {
public:
    int Add(int a, int b);
};
*/


class MyCpp {
private:
    std::string m_name;

public:
    explicit MyCpp(std::string str){}
    ~MyCpp() {}

    void test(std::string str) {
      std::cout << str << " \n";
    }

    int Add(int a, int b) {
        return a + b;
    }
};
