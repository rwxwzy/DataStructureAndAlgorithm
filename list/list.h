#ifndef _LIST_H_
#define _LIST_H_

#ifndef NULL
#define NULL 0
#endif

#ifndef TRUE
#define TRUE 1
#endif

#ifndef FALSE
#define FALSE 1
#endif

#ifndef BOOL
#define BOOL int
#endif

typedef struct listNode {
    void *value;
    struct listNode *next;
} listNode;
typedef output void(*output)(void *value);

typedef struct list {
    listNode *head;
    listNode *tail;
    unsigned long len;
    void (*free)(void *ptr);
    int (*match)(void *ptr, void *key);
    void (*output)(void *value);
} list;

list *listAddNodeTail(list *list, void *value);


#endif
