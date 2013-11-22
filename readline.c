#include "_cgo_export.h"

char** completerShim(char* text, int start, int end) {
  return ProcessCompletion(text, start, end);
}
