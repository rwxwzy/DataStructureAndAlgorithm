#include <stdlib.h>
#include <stdio.h>
#include "list.h"



/*
 *  each element in list and node
 *
 */

list* listCreate(output output) {
    struct list *list;
    if ((list = (struct list*) malloc(sizeof(struct list))) == NULL) {
        return NULL;
    }
	
    list->head = NULL;
    list->tail = NULL;
    list->len = 0;
    list->output = output;
}

void listRelease(list *list) {
    listNode *node = list->head;
    listNode *tmp;
    while(list->len != 0) {
        tmp = node->next;
        if (list->free != NULL) {
            list->free(node->value);
            free(node);
        }
        
        node = tmp;
        list->len --;
    }
    
    free(list);
}

void listOutput(list *list) {
    if (list->output == NULL) {
        return;
    }
    
    unsigned long len = list->len;
    listNode *current = list->head;
    while(len) {
        list->output(current->value);
        current = current->next;
        len --;
    }
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

list *listInsertNodeAfter(list *list, listNode *old_node, void *value) {
    listNode *node;
    if ((node = (listNode*)malloc(sizeof(listNode))) == NULL) {
        return NULL;
    }
    node->value = value;
    
    node->next = old_node->next;
    old_node->next = node;
    
    if (list->tail == old_node) {
        list->tail = node;
    }
    
    list->len ++;
}

list *listInsertNodeBefore(list *list, listNode *old_node, void *value) {
    listNode *node;
    if ((node = (listNode*)malloc(sizeof(listNode))) == NULL) {
        return NULL;
    }
    
    BOOL found = FALSE;
    listNode *cur = list->head;
    listNode *prev = list->head;
    while (cur != NULL) {
        if (cur == old_node) {
            node->next = prev->next;
            prev->next = node;
            found = TRUE;
            break;
        }
        
        prev = cur;
        cur = cur->next;
    }
    
    if (found) {
        // at least one old_node
        if (list->head == old_node) {
            list->head = node;
        }
        
        node->value = value;
        list->len ++;
        return list;
    }

    return NULL;
}

void listDelNode(list *list, listNode *node) {
    if (list->len == 0) {
        return;
    }
    
    BOOL found = FALSE;
    listNode *cur = list->head;
    listNode *prev = list->head;
    while (cur != NULL) {
        if (cur == node) {
            prev->next = node->next;
            found = TRUE;
            break;
        }
        
        prev = cur;
        cur = cur->next;
    }
    
    if (found) {
        if (list->len == 0) {
        } else if (list->len == 1) {
            list->free(list->head);
            list->len = 0;
            list->head = NULL;
            list->tail = NULL;
        } else {
            if (list->head == node) {
                list->head = node->next;
            }
            if (list->tail == node) {
                list->tail = prev;
            }
            
            list->len --;
        }
    }
}







