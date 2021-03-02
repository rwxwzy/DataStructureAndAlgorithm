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

/* stop when fast->next == NULL for odd
0        1        2        3        4        5        6        7        8
        slow     fast   
                 slow               fast
                           slow                      fast 
                                    slow                               fast// stop
*/

/* stop when fast->next->next == NULL for even
0        1        2        3        4        5        6        7        8        9
        slow     fast   
                 slow               fast
                           slow                      fast 
                                    slow                               fast //stop
*/
// So, when fast->next == NULL or fast->next->next == NULL, stop

BOOL isPalindrome(list *list) {
    if (list->head == NULL || list->head->next == NULL) { // 0 or 1
        return TRUE;
    }
    
    if (list->head->next == NULL) { // 2
        if (list->match(list->head->value, list->head->next->value)) {
            return TRUE;
        } else {
            return FALSE;
        }
    }
    
    listNode *slow, *fast, *prev, *next;
    slow = fast = next = list->head;
    prev = NULL; // for first node
    
    // find middle and reverse the next pointer for the first middle part
    while(1) {
        if (fast->next != NULL && fast->next->next != NULL) {
                // no change at first
                fast = fast->next->next;
                
                //slow related
                next = slow->next;
                slow->next = prev;
                prev = slow;
                slow = next;                              
        } else {
            break;
        }
    }
    
    // slow stop middle for odd, or left middle for even
    // when fast stop, slow->next is not handle in the loop
    // so change here
    next = slow->next;
    slow->next = prev;
    
    listNode * forward, *reverse;
    if (list->len % 2 == 0) {
        forward = next;
        reverse = slow;
    } else {
        forward = next;
        reverse = prev;
    }
    
   
    while(forward) {
        printf("%c, %c\n", *(char*)reverse->value, *(char*) forward->value);
        if (! (list->match(reverse->value, forward->value))) {
            // release
            return FALSE;
        }
        
        forward = forward->next;
        reverse = reverse->next;
    }
    // recover
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
     printf("-------------------start\n");
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
    char d[8] = {'a', 'b', 'c', 'd', 'e', 'c', 'b', 'a'};
    
    char e[1] = {'a'};
    char f[2] = {'a', 'b'};
    char g[2] = {'a', 'a'};
    char h[3] = {'a', 'b', 'a'};
    char i[4] = {'a', 'b', 'b', 'c'};
    

    isArrayPalindrome(a, sizeof(a)/ sizeof(a[0]));
    isArrayPalindrome(b, sizeof(b)/ sizeof(b[0]));
    isArrayPalindrome(c, sizeof(c)/ sizeof(c[0]));
    isArrayPalindrome(d, sizeof(d)/ sizeof(d[0]));
    isArrayPalindrome(e, sizeof(e)/ sizeof(e[0]));
    isArrayPalindrome(f, sizeof(f)/ sizeof(f[0]));
    isArrayPalindrome(g, sizeof(g)/ sizeof(g[0]));
    isArrayPalindrome(h, sizeof(h)/ sizeof(h[0]));
    
    isArrayPalindrome(i, sizeof(i)/ sizeof(h[0]));
}




