#ifndef WRAPPER_H
#define WRAPPER_H

#ifdef __cplusplus
extern "C" {
#endif

// Goから呼ぶためのC言語スタイルの関数
int call_cpp_add(int a, int b);

#ifdef __cplusplus
}
#endif

#endif