#include <stdio.h>

void main() {
    int a[11] = {9, 8, 7, 6, 5, 4, 3, 3, 2, 1, 0};

    int size = sizeof(a) / sizeof(a[0]);

    int i, j;
    int pos;
    int min;
    for (i = 0; i < size; i ++) {
        pos = i;
        for (j = i; j < size; j ++) {
            if (a[pos] > a[j]) {
		pos = j;
	    }
	}
	min = a[pos];
	a[pos] = a[i];
	a[i] = min;
    }

    for (i = 0; i < size; i ++) {
        printf("%d ", a[i]);
    }
    printf("\n");
}
