#include "string_array.h"
#include <stdio.h>
#include <string.h>

StringArray read_file(char filename[]) {
  FILE *fptr;
  fptr = fopen("test_input.txt", "r");

  StringArray arr = init_string_array(4);

  char buffer[1024];
  while (fgets(buffer, sizeof(buffer), fptr)) {
    buffer[strcspn(buffer, "\n")] = 0;
    append_string(&arr, buffer);
  }

  fclose(fptr);
  return arr;
}

int main(int argc, char *argv[]) {
  char filename[] = "test_input.txt";
  StringArray lines = read_file(filename);

  for (int i = 0; i < lines.size; i++) {
    printf("Line %d: %s\n", i, lines.data[i]);
  }

  // free_string_array(&lines);
  return 0;
}
