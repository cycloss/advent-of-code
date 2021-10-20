
#include <arrayList.h>
#include <hashTable.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define BUFF_SIZE 40
#define ALPH_SIZE 26

hashTable* ht;
int letters[ALPH_SIZE] = { 0 };

void printer(void* str) {
    printf("%s", (char*)str);
}

void addLineCharsToHt(char* line) {
    for (int i = 0; i < strlen(line); i++) {
        if (line[i] == '\n') {
            continue;
        }
        char* c = malloc(sizeof(char));
        letters[line[i] - 'a']++;
        *c = line[i];
        if (!addTableItem(ht, c)) {
            free(c);
        }
    }
}

int everyoneCount = 0;

void updateEveryoneCount(int groupSize) {

    for (int i = 0; i < ALPH_SIZE; i++) {
        if (letters[i] == groupSize) {
            everyoneCount++;
        }
    }
    memset(letters, 0, sizeof(letters));
}

int processGroup(arrayList* group) {

    int groupSize = getSize(group);

    for (int i = 0; i < groupSize; i++) {
        char* line = getItemAt(group, i);
        addLineCharsToHt(line);
    }
    updateEveryoneCount(groupSize);
    int unique = getTableSize(ht);
    clearAl(group, true);
    clearTable(ht, true);
    return unique;
}

int main() {

    FILE* in = fopen("day6.txt", "r");

    ht = createHashTable(intHash, intComp);

    arrayList* al = createArrayList();
    int total = 0;
    for (char buff[BUFF_SIZE], *ret = fgets(buff, BUFF_SIZE, in);; ret = fgets(buff, BUFF_SIZE, in)) {
        if (!ret || buff[0] == '\n') {
            total += processGroup(al);
            if (!ret) {
                break;
            }
        } else {
            char* line = malloc(sizeof(char) * BUFF_SIZE);
            strcpy(line, buff);
            appendToAl(al, line);
        }
    }
    printf("\nTotal: %d\n", total);
    printf("Everyone total: %d\n", everyoneCount);
    fclose(in);
    freeAl(al, false);
    freeTable(ht, false);
    return 0;
}