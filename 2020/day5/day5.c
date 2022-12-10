

#include <hashTable.h>
#include <stdio.h>
#include <stdlib.h>
#include <utilities.h>

void intPrinter(void* p) {
    printf("%d\n", *(int*)p);
}

int main() {
    timer(START);
    FILE* in = fopen("day5.txt", "r");
    hashTable* ht = createHashTable(intHash, intComp);
    int max = 0;
    for (char buff[12]; fgets(buff, 12, in);) {
        int low = 0, high = 127;
        for (int i = 0; i < 7; i++) {
            if (buff[i] == 'F') {
                high = (high - 1 - low) / 2 + low;
            } else {
                low = (high + 1 - low) / 2 + low;
            }
        }
        int l = 0, r = 7;
        for (int i = 7; i < 10; i++) {
            if (buff[i] == 'L') {
                r = (r - 1 - l) / 2 + l;
            } else {
                l = (r + 1 - l) / 2 + l;
            }
        }
        int* id = malloc(sizeof(int));
        *id = low * 8 + l;
        max = *id > max ? *id : max;
        addTableItem(ht, id);
    }
    printf("Max ID: %d\n", max);

    for (int i = 0, i2 = 1, i3 = 2; i < 1021; i++, i2++, i3++) {
        bool first = tableContains(ht, &i), second = tableContains(ht, &i2), third = tableContains(ht, &i3);
        if (first && !second && third) {
            printf("My seat ID: %d\n", i2);
            break;
        }
    }
    freeTable(ht, true);
    fclose(in);
    timer(STOP);
}