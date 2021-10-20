#include <regex.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define INPUT_FILE_2 "inputFiles/day2.txt"
#define BUFFER_MAX 100

typedef struct {
    int min, max;
    char keyChar, *pass;
} cleanPassLine;

bool testPasswordOne(char* password);
bool testPasswordTwo(char* password);
cleanPassLine* generateCpl(char* passLine);
char* getMatch(int index, regmatch_t* matches, char* passLine);
void runTest();

int main() {
    puts("Part one:");
    runTest(testPasswordOne);
    puts("\nPart two:");
    runTest(testPasswordTwo);
}

void runTest(bool (*testPassword)(char*)) {

    FILE* inputFile = fopen(INPUT_FILE_2, "r");

    int okPasswords = 0;

    for (char buffer[BUFFER_MAX]; fgets(buffer, BUFFER_MAX, inputFile);) {
        if (testPassword(buffer)) {
            okPasswords++;
        }
    }
    printf("Ok passwords: %d\n", okPasswords);

    fclose(inputFile);
}

bool testPasswordOne(char* passwordLine) {
    cleanPassLine* cpl = generateCpl(passwordLine);
    if (!cpl) {
        printf("Failed to parse line: %s", passwordLine);
        exit(1);
    }

    int keyCharCount = 0;
    char* copyp = cpl->pass;
    for (char c = *copyp; c; c = *++copyp) {
        if (c == cpl->keyChar) {
            keyCharCount++;
        }
    }

    bool ok = keyCharCount <= cpl->max && keyCharCount >= cpl->min;
    free(cpl->pass);
    free(cpl);
    return ok;
}

bool testPasswordTwo(char* passwordLine) {
    cleanPassLine* cpl = generateCpl(passwordLine);
    if (!cpl) {
        printf("Failed to parse line: %s", passwordLine);
        exit(1);
    }

    int correctPos = 0;

    char* copyP = cpl->pass - 1;
    if (*(copyP + cpl->min) == cpl->keyChar) {
        correctPos++;
    }
    if (*(copyP + cpl->max) == cpl->keyChar) {
        correctPos++;
    }

    free(cpl->pass);
    free(cpl);
    return correctPos == 1;
}

cleanPassLine* generateCpl(char* passLine) {
    char* regex = "([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)";
    size_t maxGroups = 5;

    regex_t reg;
    regmatch_t matches[maxGroups];

    if (regcomp(&reg, regex, REG_EXTENDED)) {
        puts("Failed to compile regex");
        return NULL;
    }

    if (regexec(&reg, passLine, maxGroups, matches, 0)) {
        puts("Failed to find match");
        return NULL;
    }
    regfree(&reg);

    char* min = getMatch(1, matches, passLine);
    int minNum = atoi(min);
    free(min);

    char* max = getMatch(2, matches, passLine);
    int maxNum = atoi(max);
    free(max);

    char* keyCharP = getMatch(3, matches, passLine);
    char keyChar = *keyCharP;
    free(keyCharP);

    char* passP = getMatch(4, matches, passLine);

    cleanPassLine* cpl = malloc(sizeof(cleanPassLine));
    *cpl = (cleanPassLine) { minNum, maxNum, keyChar, passP };
    return cpl;
}

char* getMatch(int index, regmatch_t* matches, char* passLine) {
    char* cpy = malloc(sizeof(char) * (strlen(passLine) + 1));
    strcpy(cpy, passLine);
    cpy[matches[index].rm_eo] = '\0';
    sprintf(cpy, "%s", cpy + matches[index].rm_so);
    return cpy;
}