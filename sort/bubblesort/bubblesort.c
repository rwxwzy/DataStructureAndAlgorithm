#include <stdio.h>

void bubblesort(int a[], int size) {
    int i, j;
    int changed;
    int tmp;
    for (i = 0; i < size - 1; i ++) {
        changed = 0;
        for (j = 0; j < size - 1; j ++) {
            
            if (a[j] > a[j + 1]) {
                tmp = a[j + 1];
                a[j + 1] = a[j];
                a[j] = tmp;
                changed = 1;
            }
        }
        
        if (changed == 0) {
            break;
        }
        
    }
    
}


void main() {
    int a[10] = { 9, 8, 7, 6, 5, 4,3, 2, 1, 0};
    bubblesort(a, sizeof(a)/ sizeof(a[0]));
    int i;
    for (i = 0; i < sizeof(a)/ sizeof(a[0]); i ++) {
        printf("%d ", a[i]);
    }
    
    printf("\n");
    
}