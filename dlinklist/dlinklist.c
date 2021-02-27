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
    
    if (list->len == 0) {// clear
        list->head = list->tail = node;
        node->next = node->prev = NULL;
    } else {    
        node->next = list->head;
        node->prev = NULL;
        list->head = node;
    }
    
    list->len ++;
    return list;
}

dLinkList* dLinkListAddTailNode(dLinkList *list, void *value) {
    dLinkListNode *node;
    if ((node = (dLinkListNode*)malloc(sizeof(dLinkListNode))) == NULL) {
        return NULL;
    }
    
    node->value = value;
    if (list->len == 0) {
        list->head = list->tail = node;
        node->next = node->prev = NULL;
    } else {
        node->prev = list->tail;
        node->next = NULL;
        list->tail->next = node;
        list->tail = node;
    }
    
    list->len ++;
    return list;
}

dLinkList *dLinkListInsertNode(dLinkList *list, dLinkListNode *ref_node, void *value, BOOL after) {
     dLinkListNode *node;
     if ((node = (dLinkListNode*)malloc(sizeof(dLinkList))) == NULL) {
         return NULL;
     }
     node->value = value;
     
     if (after) {
         node->prev = ref_node;
         node->next = ref_node->next;
         
         ref_node->next = node;
         if (node->next != NULL) {
             node->next->prev = node;
         }
         if (list->tail == ref_node) {
             list->tail = node;
         }
     } else {
         node->next = ref_node;
         node->prev = ref_node->prev;
         
         if (node->prev != NULL) {
             node->prev->next = node;
         }
         ref_node->prev = node;
         
         if (list->head == ref_node) {
             list->head = node;
         }
     }
     
     list->len ++;
     return list;
}






void main() {
}










