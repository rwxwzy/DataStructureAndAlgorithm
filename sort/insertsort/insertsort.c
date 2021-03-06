#include <stdio.h>

void main() {
    int a[11] = {9, 8, 7, 6, 5, 4, 3, 3, 2, 1, 0};
    
    int i, j;
    int size = sizeof(a)/ sizeof(a[0]);
    int tmp;
    for (i = 1; i < size; i ++) {
        tmp = a[i];
        // start from back, move until after is bigger than before
        for (j = i - 1; j >= 0; j --) {
            if (tmp < a[j]) {
                a[j + 1] = a[j];
            } else {
                break;
            }
        }
        
        a[j + 1] = tmp;
    }
    
    for (i = 0; i < size; i ++) {
        printf("%d ", a[i]);
    }
    
    printf("\n");
    
}