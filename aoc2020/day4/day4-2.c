//second version, more robust but more complicated, uses custom hashtable that can be found in my github repos
#include <arrayList.h>
#include <hashTable.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    bool byrOk, iyrOk, eyrOk, hgtOk, hclOk, eclOk, pidOk;
} ppValidator;

void printStringp(void* str) {
    puts((char*)str);
}

arrayList* getPassports() {
    FILE* in = fopen("inputFiles/day4.txt", "r");
    arrayList* al = createArrayList();
    char *buffer = malloc(sizeof(char) * 100), lastChar;
    int buffInd = 0;

    for (char c = fgetc(in); lastChar != EOF; c = fgetc(in)) {
        if (c == '\n' && lastChar == '\n' || c == EOF) {
            buffer[buffInd] = '\0';
            appendToAl(al, buffer);
            buffer = malloc(sizeof(char) * 100);
            buffInd = 0;
        } else if (c == ' ' || c == '\n') {
            buffer[buffInd++] = ' ';
        } else {
            buffer[buffInd++] = c;
        }
        lastChar = c;
    }
    free(buffer);
    fclose(in);
    return al;
}

bool checkPpv(ppValidator* ppvp) {
    for (int i = 0; i < 7; i++) {
        if (!((bool*)ppvp)[i]) {
            return false;
        }
    }
    return true;
}

bool partTwo = false;
hashTable* eyeColorTable;
char* codes[] = { "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid" };

void processSection(char* section, ppValidator* ppvp) {
    section[3] = '\0';
    if (strcmp(section, codes[0]) == 0) {
        if (partTwo) {
            section += 4;
            int year = atoi(section);
            year >= 1920 && year <= 2002 ? ppvp->byrOk = true : false;
        } else {
            ppvp->byrOk = true;
        }
    } else if (strcmp(section, codes[1]) == 0) {
        if (partTwo) {
            section += 4;
            int year = atoi(section);
            year >= 2010 && year <= 2020 ? ppvp->iyrOk = true : false;
        } else {
            ppvp->iyrOk = true;
        }
    } else if (strcmp(section, codes[2]) == 0) {
        if (partTwo) {
            section += 4;
            int year = atoi(section);
            year >= 2020 && year <= 2030 ? ppvp->eyrOk = true : false;
        } else {
            ppvp->eyrOk = true;
        }
    } else if (strcmp(section, codes[3]) == 0) {
        if (partTwo) {
            section += 4;
            int height = atoi(section);
            int len = strlen(section);
            if (section[len - 2] == 'c' && section[len - 1] == 'm') {
                height >= 150 && height <= 193 ? ppvp->hgtOk = true : false;
            } else if (section[len - 2] == 'i' && section[len - 1] == 'n') {
                height >= 59 && height <= 76 ? ppvp->hgtOk = true : false;
            }
        } else {
            ppvp->hgtOk = true;
        }
    } else if (strcmp(section, codes[4]) == 0) {
        if (partTwo) {
            section += 4;
            int len = strlen(section);
            bool hash = *section++ == '#';
            bool charsOk = true;
            for (int i = 0; i < len - 1; i++) {
                char c = section[i];
                if (!((section[i] >= '0' && section[i] <= '9') || (section[i] >= 'a' && section[i] <= 'f'))) {
                    charsOk = false;
                }
            }
            hash&& len == 7 && charsOk ? ppvp->hclOk = true : false;
        } else {
            ppvp->hclOk = true;
        }
    } else if (strcmp(section, codes[5]) == 0) {
        if (partTwo) {
            section += 4;
            ppvp->eclOk = tableContains(eyeColorTable, section);
        } else {
            ppvp->eclOk = true;
        }
    } else if (strcmp(section, codes[6]) == 0) {
        if (partTwo) {
            section += 4;
            int len = strlen(section);
            bool numsOk = true;
            for (int i = 0; i < len; i++) {
                if (!(section[i] <= '9' && section[i] >= '0')) {
                    numsOk = false;
                }
            }
            len == 9 && numsOk ? ppvp->pidOk = true : false;
        } else {
            ppvp->pidOk = true;
        }
    }
}

bool checkPassport(char* passport) {

    ppValidator ppv = { false, false, false, false, false, false, false };
    for (char* section = strtok(passport, " "); section; section = strtok(NULL, " ")) {
        processSection(section, &ppv);
    }
    return checkPpv(&ppv);
}

int getValidPassports(arrayList* passports) {

    int validCount = 0;
    for (int i = 0; i < getSize(passports); i++) {
        if (checkPassport(getItemAt(passports, i))) {
            validCount++;
        }
    }
    return validCount;
}

int main() {

    arrayList* al = getPassports();
    printf("Passport count: %d\n", getSize(al));
    printf("Part one valid passports: %d\n", getValidPassports(al));
    freeAl(al, true);

    partTwo = true;
    eyeColorTable = createHashTable(strHash, strComp);
    addTableItems(eyeColorTable, 7, "amb", "blu", "brn", "gry", "grn", "hzl", "oth");
    arrayList* al2 = getPassports();
    printf("Part two valid passports: %d\n", getValidPassports(al2));
    freeAl(al2, true);
    freeTable(eyeColorTable, false);
}