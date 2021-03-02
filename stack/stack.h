#ifndef _STACK_H_
#define _STACK_H_

#ifndef NULL
#define NULL 0
#endif

#ifndef TRUE
#define TRUE 1
#endif

#ifndef FALSE
#define FALSE 0
#endif

#ifndef BOOL
#define BOOL int
#endif

typedef char ElemType;


typedef struct {
    int len;
    ElemType *buf;
    int top;
} stack;


stack* createStack(int len);

stack * stackPush(stack *s, ElemType e);

BOOL stackPop(stack *s, ElemType* e);


#endif

