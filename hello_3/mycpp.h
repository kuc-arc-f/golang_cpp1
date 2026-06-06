#ifndef MYCPP_H
#define MYCPP_H
#include <string>

class Printer {
public:
    // C++本来のstd::stringを使うメソッド
    void PrintMessage(const std::string& msg);
};


class Provider {
public:
    std::string GetMessage() {
        return "C++の世界からこんにちは！";
    }
};

#endif