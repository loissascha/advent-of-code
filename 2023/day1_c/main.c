#include <stdio.h>

int main(int argc, char *argv[]) {
  puts("Hello world");

  FILE *fptr;
  fptr = fopen("test_input.txt", "r");

  char myString[8];
  while (fgets(myString, 8, fptr)) {
    printf("read: %s\n", myString);
  }

  fclose(fptr);

  return 0;
}
