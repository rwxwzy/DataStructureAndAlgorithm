#include <stdio.h>

void bubblesort(int a[], int size) {
    int i, j;
    int changed;
    int tmp;
    int step = 0;
    for (i = 0; i < size - 1; i ++) {
        changed = 0;
	// size - 1 - i, is already sorted
        for (j = 0; j < size - 1 - i; j ++) {
        // for (j = 0; j < size - 1; j ++) {
            
            if (a[j] > a[j + 1]) {
                tmp = a[j + 1];
                a[j + 1] = a[j];
                a[j] = tmp;
                changed = 1;
            }
	    step ++;
        }
        
        if (changed == 0) {
            break;
        }
        
    }

    printf("step: %d\n", step);
    
}


void main() {
    int a[11] = { 9, 8, 7, 6, 5, 4,3, 3,  2, 1, 0};
    bubblesort(a, sizeof(a)/ sizeof(a[0]));
    int i;
    for (i = 0; i < sizeof(a)/ sizeof(a[0]); i ++) {
        printf("%d ", a[i]);
    }
    
    printf("\n");
    
}
