#include "string_array.h"
#include <stdlib.h>
#include <string.h>

StringArray init_string_array(int initial_capacity) {
  StringArray arr;
  arr.data = malloc(initial_capacity * sizeof(char));
  arr.size = 0;
  arr.capacity = initial_capacity;
  return arr;
}

void append_string(StringArray *arr, const char *str) {
  if (arr->size >= arr->capacity) {
    arr->capacity *= 2;
    arr->data = realloc(arr->data, arr->capacity * sizeof(char));
  }
  arr->data[arr->size] = malloc(strlen(str) + 1);
  strcpy(arr->data[arr->size], str);
  arr->size++;
}

void free_string_array(StringArray *arr) {
  for (int i = 0; i < arr->size; i++) {
    free(arr->data[i]);
  }
  free(arr->data);
  arr->size = 0;
  arr->capacity = 0;
}
