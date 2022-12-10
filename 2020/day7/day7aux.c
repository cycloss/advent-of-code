#include "day7.h"

extern hashMap* bagMap;
extern arrayList* bagList;

void bagKvpPrinter(keyValPair* kvpp) {
    printf("[%s : ", (char*)kvpp->key);
    for (int i = 0; i < getSize(kvpp->val); i++) {
        bag* bagp = getItemAt(kvpp->val, i);
        printf("(%i %s), ", bagp->count, bagp->color);
    }
    printf("]\n");
}

void strPrinter(void* str) {
    puts((char*)str);
}

char* getBagName(char** line) {
    char* bagNameReg = "^([a-z]+ [a-z]+) bags contain? ";
    size_t matchCount = 3;

    regex_t reg;
    regmatch_t matches[matchCount];

    if (regcomp(&reg, bagNameReg, REG_EXTENDED) || regexec(&reg, *line, matchCount, matches, 0)) {
        puts("Failed to execute bag name regex");
        exit(1);
    }

    regmatch_t fullMatch = matches[0];
    int fullLen = fullMatch.rm_eo - fullMatch.rm_so;

    regmatch_t nameMatch = matches[1];
    int len = nameMatch.rm_eo - nameMatch.rm_so;
    char* matchStr = malloc(sizeof(char) * len + 1);
    matchStr[len] = '\0';
    memcpy(matchStr, *line + nameMatch.rm_so, len);
    (*line) += fullLen;
    regfree(&reg);
    return matchStr;
}

//double pointer so address can be modified and incremented for next bag
bag* getNextBag(char** line) {
    char* enclosedBagReg = "([1-9]+) ([a-z]+ [a-z]+) bags?[ .,]?[ ]?";
    int matchCount = 4;
    regex_t reg;
    regmatch_t matches[matchCount];

    if (regcomp(&reg, enclosedBagReg, REG_EXTENDED) || regexec(&reg, *line, matchCount, matches, 0)) {
        regfree(&reg);
        return NULL;
    }

    bag* currentBag = malloc(sizeof(bag));

    for (int i = 1;; i++) {

        regmatch_t match = matches[i];

        if (match.rm_so == -1) {
            break;
        }

        int len = match.rm_eo - match.rm_so;
        char* matchStr = malloc(sizeof(char) * len + 1);
        matchStr[len] = '\0';
        memcpy(matchStr, *line + match.rm_so, len);

        if (i == 1) {
            currentBag->count = atoi(matchStr);
            free(matchStr);
        } else if (i == 2) {
            currentBag->color = matchStr;
            regmatch_t fullMatch = matches[0];
            int fullLen = fullMatch.rm_eo - fullMatch.rm_so;
            (*line) += fullLen;
        } else {
            puts("Too many matches in bag match");
            puts(matchStr);
            exit(1);
        }
    }
    regfree(&reg);
    return currentBag;
}

void freeBag(void* bagp) {
    bag* b = (bag*)bagp;
    free(b->color);
    free(b);
}

void freeMapPair(keyValPair* kvpp) {
    char* key = kvpp->key;
    arrayList* l = kvpp->val;
    iterateListItems(l, freeBag);
    freeAl(l, false);
    free(key);
}

void freeBagMap() {
    iterateMapPairs(bagMap, freeMapPair);
    freeMap(bagMap, false);
}