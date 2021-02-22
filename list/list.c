#include <stdlib.h>
#include <stdio.h>
#include "list.h"

typedef struct listNode {
    void *value;
    struct listNode *next;
} listNode;

typedef struct list {
    listNode *head;
    listNode *tail;
    unsigned long len;
} list;

/*
 *  each element in 4 components: list, head, tail, node
 *
 */

list* listCreate() {
    struct list *list;
    if ((list = (struct list*) malloc(sizeof(struct list))) == NULL) {
        return NULL;
    }
	
    list->head = NULL;
    list->tail = NULL;
    list->len = 0;
}

void emptyList(list *list) {
        
}

list* listAddNodeHead(list *list, void *value) {
    listNode *node;
    if ((node = (listNode*)malloc(sizeof(listNode))) == NULL) {
        return NULL;
    }
	
    node->value = value;
	if (list->len == 0) { // set tail
		list->head = list->tail = node;
		node->next = NULL;
	} else {
		node->next = list->head;
	    list->head = node;
	}

	list->len ++;
	
	return list;
}

list *listAddNodeTail(list *list, void *value) {
	listNode *node;
	if ((node = (listNode *)malloc(sizeof(listNode))) == NULL) {
		return NULL;
	}
	
	node->value = value;
	if (list->len == 0) {
		list->head = list->tail = node;
		node->next = NULL;
	} else {
		list->tail->next = node;
		list->tail = node;
		node->next = NULL;
	}
	
	list->len ++;
}


void main() {
	
}





