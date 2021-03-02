#include <stdio.h>
#include <stdlib.h>
#include "stack.h"

BOOL testMath(char buf[], int size, int *result) {
    ElemType e;
    
    stack* nums = createStack(2);
    if (nums == NULL) {
        return FALSE;
    }

    stack* symbols = createStack(2);
    if (symbols == NULL) {
        stackFree(nums);
        return FALSE;
    }
    
    int i;
    for(i = 0; i < size; i ++) {
        if (buf[i] >= '0' && buf[i] <= '9') {
            
        }
    }
    
}

void main() {

    char *buf = "1+ 2 *3 + 6";
    int result;
    if(testMath(buf, strlen(buf), &result)) {
        printf("result: %d\n", *result);
    }
}