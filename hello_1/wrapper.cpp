#include "wrapper.h"
#include "mycpp.h"

int call_cpp_add(int a, int b) {
    Calculator calc;
    return calc.Add(a, b); // 内部でC++のインスタンスを作って実行
}