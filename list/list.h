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

typedef void (*output)(void *value);

typedef struct list {
    listNode *head;
    listNode *tail;
    unsigned long len;
    void (*free)(void *ptr);
    int (*match)(void *ptr, void *key);
    void (*output)(void *value);
} list;

list *listAddNodeTail(list *list, void *value);
list* listCreate(output output);
void listOutput(list *list);
void listRelease(list *list);


#endif
