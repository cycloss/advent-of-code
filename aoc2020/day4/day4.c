//first version, not very robust, only part one

#include <hashTable.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

//took way longer than it should have done because of the fact that the last passport doesn't have a blank line after it

char* requiredCodes[] = { "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid" };

bool validatePassport(hashTable* ht) {
    for (int i = 0; i < sizeof(requiredCodes) / sizeof(requiredCodes[0]); i++) {
        if (!tableContains(ht, requiredCodes[i])) {
            return false;
        }
    }
    return true;
}

int main() {
    FILE* inFile = fopen("inputFiles/day4.txt", "r");
    hashTable* ht = createHashTable(hashString, stringComparator);
    int validPP = 0, i = 0;
    for (char c = fgetc(inFile), buffer[20];; c = fgetc(inFile)) {
        if (c == ' ' || c == '\n' || c == EOF) {
            if (buffer[0] != '\n' || c == EOF) {
                buffer[3] = '\0';
                char* temp = malloc(sizeof(char) * 4);
                strcpy(temp, buffer);
                addTableItem(ht, temp);
            }
            if (buffer[0] == '\n' || c == EOF) {
                if (validatePassport(ht)) {
                    validPP++;
                }
                freeTable(ht, true);
                ht = createHashTable(hashString, stringComparator);
            }
            if (c == EOF)
                break;
            buffer[i = 0] = c;
        } else {
            buffer[i++] = c;
        }
    }
    printf("Valid: %d\n", validPP), fclose(inFile);
}