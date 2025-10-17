#ifndef STRING_ARRAY_H
#define STRING_ARRAY_H

typedef struct {
  char **data;
  int size;
  int capacity;
} StringArray;

StringArray init_string_array(int initial_capacity);

void append_string(StringArray *arr, const char *str);

void free_string_array(StringArray *arr);

#endif
