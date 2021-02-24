#include <stdlib.h>
#include <stdio.h>
#include "dlinklist.h"

typedef struct dLinkListNode {
    struct dLinkListNode *next;
    struct dLinkListNode *prev;
    void *value;
    int (*match)(void *ptr, void *key);
} dLinkListNode;

typedef struct dLinkList {
    struct dLinkListNode *head;
    struct dLinkListNode *tail;
    unsigned long len;
} dLinkList;

// check all elements in dLinkList and dLinkListNode

dLinkList *CreateDLinkList() {
    dLinkList *list;
    if ((list = (dLinkList*)malloc(sizeof(dLinkList))) == NULL) {
        return NULL;
    }
    
    list->head = NULL;
    list->tail = NULL;
    list->len = 0;
}

dLinkList *listAddHeadNode(dLinkList *list, void *value) {
    dLinkListNode *node;
    if ((node = (dLinkListNode*)malloc(sizeof(dLinkListNode))) == NULL) {
        return NULL;
    }
    node->value = value;
    
    node->next = list->head;
    node->prev = NULL;
    if (list->len = 0) {
        list->tail = node;
    }
    
    list->len ++;
    return list;
}











void main() {
}










