#include <stdio.h>
#include <stdlib.h>
#include "list.h"

void print(void *value) {
    printf("%c\n", *(char*)value);
}

BOOL charMatch(void *value1, void* value2) {
    if (*(char*) value1 == *(char*)value2) {
        return TRUE;
    } else {
        return FALSE;
    }
}

/* stop when pos2->next == NULL for odd
0        1        2        3        4        5        6        7        8
        pos1     pos2   
                 pos1               pos2
                           pos1                      pos2 
                                    pos1                               pos2// stop
*/

/* stop when pos2->next->next == NULL for even
0        1        2        3        4        5        6        7        8        9
        pos1     pos2   
                 pos1               pos2
                           pos1                      pos2 
                                    pos1                               pos2 //stop
*/
// So, when pos2->next == NULL or pos2->next->next == NULL, stop

BOOL isPalindrome(list *list) {
    if (list->len == 1) {
        return TRUE;
    }
    
    if (list->len == 2) {
        if (list->match(list->head->value, list->head->next->value)) {
            return TRUE;
        } else {
            return FALSE;
        }
    }
    
    listNode *pos1, *pos2, *prev, *next;
    pos1 = pos2 = next = list->head;
    prev = NULL; // for first node
    
    while(1) {
        if (pos2->next != NULL) {
            if (pos2->next->next != NULL) {
                // no change at first
                pos2 = pos2->next->next;
                
                //pos1 related
                next = pos1->next;
                pos1->next = prev;
                prev = pos1;
                pos1 = next;                              
            } else {
                break;
            }
        } else {
            break;
        }
    }
    
    listNode * forward, *reverse;
    if (list->len % 2 == 0) {
        // in the condition of pos2->next->next == NULL
        forward = pos1->next;
        reverse = pos1;
    } else {
        forward = pos1->next;
        reverse = prev;
    }
    
    while(forward) {
        
        if (! (list->match(reverse->value, forward->value))) {
            // release
            return FALSE;
        }
        printf("%c, %c\n", *(char*)reverse->value, *(char*) forward->value);
        
        forward = forward->next;
        reverse = reverse->next;
    }
    //release
    return TRUE;
}

//only to verify list
BOOL isArrayPalindrome(char a[], int size) {
    list *list = listCreate(print, charMatch);
    int i;
    char *tmp;
    for(i = 0; i < size; i ++ ) {
        tmp = (char*)malloc(sizeof(char));
        *tmp = a[i];
        listAddNodeTail(list, tmp);
    }
    listOutput(list);
    if(isPalindrome(list)) {
        printf("True\n");
    } else {
        printf("False\n");
    }
    
    // listRelease(list);
    
    return TRUE;
}

void main() {
    char a[7] = {'a', 'b', 'c', 'd', 'c', 'b', 'a'};
    char b[7] = {'a', 'b', 'c', 'd', 'c', 'a', 'a'};
    
    char c[8] = {'a', 'b', 'c', 'd', 'd', 'c', 'b', 'a'};
    char d[8] = {'a', 'b', 'c', 'd', 'd', 'c', 'b', 'a'};
    
    char e[1] = {'a'};
    char f[2] = {'a', 'b'};
    char g[2] = {'a', 'a'};
    
    isArrayPalindrome(a, sizeof(a)/ sizeof(a[0]));
    isArrayPalindrome(b, sizeof(b)/ sizeof(b[0]));
    isArrayPalindrome(c, sizeof(c)/ sizeof(c[0]));
    isArrayPalindrome(d, sizeof(d)/ sizeof(d[0]));
    isArrayPalindrome(e, sizeof(e)/ sizeof(e[0]));
    isArrayPalindrome(f, sizeof(f)/ sizeof(f[0]));
    isArrayPalindrome(g, sizeof(g)/ sizeof(g[0]));
}




