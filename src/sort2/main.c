#include <stdio.h>
#include <stdlib.h>

int cmp_int(const void * a_, const void * b_) {
  const int * a = a_, * b = b_;
  return a[0] - b[0];
}

int main() {
  int i;
  int a[] = { 7, 6, 3, 8, 5, 4 };
  int n = sizeof(a)/sizeof(int);
  qsort(a, n, sizeof(int), cmp_int);
  for (i = 0; i < n; i++) {
    printf("%d ", a[i]);
  }
  printf("\n");
  return 0;
}
