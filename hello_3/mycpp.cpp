#include "mycpp.h"
#include <iostream>

void Printer::PrintMessage(const std::string& msg) {
    std::cout << "[C++側が出力] " << msg << std::endl;
}