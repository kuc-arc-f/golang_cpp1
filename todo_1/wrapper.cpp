#include "wrapper.h"
#include <cstdlib>
#include <cstring>
//#include <curl/curl.h>

#include "include/mycpp.hpp"
#include "include/my_todo.hpp"

void todo_delete(int id){
    MyTodo todo_helper("");
    todo_helper.todo_delete_handler(id);    
}

void todo_add(const char* text){
    MyTodo todo_helper("");
    todo_helper.todo_add_handler(std::string(text));    
}

void todo_list(){
    MyTodo todo_helper("");
    todo_helper.todo_list_handler();    
}

void call_cpp_print(const char* text) {
    Printer printer;
    // const char* を std::string に変換してC++のメソッドに渡す
    printer.PrintMessage(std::string(text));
}

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