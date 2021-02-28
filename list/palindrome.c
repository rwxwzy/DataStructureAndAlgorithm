#include <stdio.h>
#include <stdlib.h>
#include "list.h"

void print(void *value) {
    printf("%c\n", *(char*)value);
}

BOOL isPalindrome(list *list) {
    
}

void main() {
    list *list = listCreate(print);
    
    char a[7] = {'a', 'b', 'c', 'd', 'c', 'b', 'a'};
    char b[7] = {'a', 'b', 'c', 'd', 'c', 'a', 'a'};
    
    char c[8] = {'a', 'b', 'c', 'd', 'd', 'c', 'b', 'a'};
    char d[8] = {'a', 'b', 'c', 'd', 'd', 'c', 'b', 'a'};
    int i;
    char *tmp;
    for(i = 0; i < sizeof(a)/ sizeof(a[0]); i ++ ) {
        tmp = (char*)malloc(sizeof(char));
        *tmp = a[i];
        listAddNodeTail(list, tmp);
    }
    listOutput(list);
    listRelease(list);
    
}