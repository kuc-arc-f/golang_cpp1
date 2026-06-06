#include "wrapper.h"
#include "mycpp.h"
#include <cstdlib>
#include <cstring>

char* get_cpp_message() {
    Provider provider;
    std::string cpp_str = provider.GetMessage();

    // Goに渡すために、Cのヒープメモリを確保してコピーする
    // ⚠️ +1 は文字列の終わりを示す「\0」の分です
    char* c_str = (char*)std::malloc(cpp_str.length() + 1);
    std::strcpy(c_str, cpp_str.c_str());

    return c_str; // Goへポインタを返す
}

void free_cpp_message(char* ptr) {
    // Go側から呼ばれたら、確保していたメモリを解放する
    std::free(ptr);
}