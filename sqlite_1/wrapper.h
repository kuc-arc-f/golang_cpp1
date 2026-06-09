#ifndef WRAPPER_H
#define WRAPPER_H

#ifdef __cplusplus
extern "C" {
#endif

void todo_add(const char* text);
void todo_list();
void todo_delete(int id);

// Goから呼ぶための関数。文字列は const char* で受ける
void call_cpp_print(const char* text);

// 1. 文字列を取得する関数
char* get_cpp_message();

// 2. 取得した文字列のメモリを解放する関数
void free_cpp_message(char* ptr);

#ifdef __cplusplus
}
#endif

#endif