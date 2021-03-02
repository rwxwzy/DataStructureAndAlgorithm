#include <stdio.h>
#include <stdlib.h>
#include "stack.h"


stack* createStack(int len) {
    stack *s;
    if ((s = (stack*)malloc(sizeof(stack))) == NULL) {
        return NULL;
    }
    if ((s->buf = (ElemType*)malloc(sizeof(ElemType))) == NULL) {
        return NULL;
    }
    
    s->len = len;
    s->top = -1;
    
    return s;
}

stack * stackPush(stack *s, ElemType e) {
    if (s->top < s->len - 1){
        printf("stack is full, realloc\n");
        ElemType *tmp = realloc(s->buf, 2 * s->len);
        if (!tmp) {
            printf("realloc failure\n");
            // free(s->buf);
            return NULL;
        }
        s->buf = tmp;
        s->len = 2 * s->len;
    }
    s->top ++;
    s->buf[s->top] = e;
    return s;
}

BOOL stackPop(stack *s, ElemType* e) {
    if (s->top == -1){
        printf("No data in stack\n");
        return FALSE;
    }
    
    *e = s->top;
    s->top --;
}

void stackFree(stack *s) {
    free(s->buf);
    free(s);
}



