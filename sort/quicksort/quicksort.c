#include <stdio.h>


void print(int a[], int size) {
    int i;
    for (i = 0; i < size; i ++) {
        printf("%d ", a[i]);
    }
    printf("\n");
}


// 1 2 9 8 3  7  4 0 1 6 5

// sort by partition and then recursively
// pick right as pivot
int partition(int a[], int left, int right) {
    printf("+++++++++++++++\n");
    printf("Before Partition: \n");
    printf("left: %d, right: %d\n", left, right);
    print(a + left, right - left + 1);
    
    int i = left;
    int j = left;
    int tmp;
    for (j = left; j < right; j ++) {
        if (a[j] > a[right]) {
            if (a[i] < a[right]) {
                i = j;
            }
        } else {
            if (a[i] > a[right] ) {
                tmp = a[i];
                a[i] = a[j];
                a[j] = tmp;
            }
            i ++;
        }
        
        printf("a[i]: %d, a[j]: %d, a[right]: %d\n", a[i], a[j], a[right]);
    }
    
    // switch between i and the pivot
    tmp = a[right];
    a[right] = a[i];
    a[i] = tmp;
    
    printf("After Partition: \n");
    printf("i: %d\n", i);
    print(a + left, right - left + 1);
    printf("-------------------\n");
    
    return i;
}

void quickSort(int a[], int left, int right) {
    if (right <= left) {
        printf("right: %d <= left: %d\n", right, left);
        return ;
    }
    
    int i = partition(a, left, right);
    quickSort(a, left, i - 1);
    quickSort(a, i + 1, right);
}

void test(int a[], int size) {
    int i;
    printf("Before Sort: \n");
    print(a, size);
    printf("\n");
    
    quickSort(a, 0, size - 1);
    
    printf("\nAfter Sort: \n");
    print(a, size);
}

void main() {
    int a[] = {9, 8, 7, 6, 5, 4, 3, 3, 2, 1, 0 };
    // 1 2 9 8 3  7  4 0 1 6 5
    int b[] = {1, 2, 9, 8, 3, 7, 4, 0, 1, 6, 5};
    
    int c[] = {5, 8, 9, 23, 67, 1, 3, 7, 31, 56};
    // 1 3 5 7 8 9 23 31 56 67
 
    
    test(a, sizeof(a)/sizeof(a[0]));
    //test(b, sizeof(b)/sizeof(b[0]));

    //test(c, sizeof(c)/sizeof(c[0]));
}