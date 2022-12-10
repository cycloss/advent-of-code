#include <arrayList.h>
#include <hashMap.h>
#include <regex.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    char* color;
    int count;
} bag;

void bagKvpPrinter(keyValPair* kvpp);
void strPrinter(void* str);
char* getBagName(char** line);
bag* getNextBag(char** line);
void freeBag(void* bagp);
void freeMapPair(keyValPair* kvpp);
void freeBagMap();

int findBagsWithinBag(char* bagName, int multiplier);