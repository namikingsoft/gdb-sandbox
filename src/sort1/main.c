#include <stdio.h>

void printarray(int *,int);
void sort(int *,int);
void swap(int *,int *);

/* program entry */
int main(int argc, char **argv) {
  int array[4] = {4, 1, 3, 2};
  printarray(&array[0], 4);
  sort(array, 4);
  printarray(&array[0], 4);
  return 0;
}

/* printout */
void printarray(int *array, int length) {
  int i;
  for (i = 0; i < length; i++) {
    printf("%d ", array[i]);
  }
  printf("\n");
}

/* bubble sort */
void sort(int *array, int length) {
  int i, j;
  for (i = 0; i < length - 1; i++) {
    for (j = 0; j < length - i - 1; j++) {
      if (array[j] > array[j+1]) {
        swap(&array[j], &array[j+1]);
      }
    }
  }
}

/* swap a <=> b */
void swap(int *a, int *b) {
  int temp;
  temp = *a;
  *a = *b;
  *b = temp;
}
