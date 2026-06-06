#include "wrapper.h"
#include "mycpp.hpp"

int call_cpp_add(int a, int b) {
    MyCpp mycpp("example");
    return mycpp.Add(a, b); // 内部でC++のインスタンスを作って実行
}