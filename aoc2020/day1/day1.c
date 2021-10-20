#include <stdio.h>
#include <stdlib.h>

#define INPUT_FILE_1 "day1.txt"
#define MAX_CHAR 200

void partOne() {

    FILE* dayOneInput = fopen(INPUT_FILE_1, "r");

    char buckets[2000] = { 0 };

    for (char buffer[MAX_CHAR]; fgets(buffer, MAX_CHAR, dayOneInput);) {
        int num = atoi(buffer);
        int searchVal = 2020 - num;
        if (buckets[searchVal]) {
            printf("Nums: %d, %d\n", num, searchVal);
            printf("Answer: %d\n", num * searchVal);
            break;
        }
        buckets[num] = 1;
    }

    fclose(dayOneInput);
}

#define BUCKET_COUNT 10000

typedef struct {
    int num1;
    int num2;
} numPair;

numPair* createNumPair(int num1, int num2) {
    numPair* np = malloc(sizeof(numPair));
    *np = (numPair) { num1, num2 };
    return np;
}

void freePairs(numPair** buckets, int len) {
    for (int i = 0; i < len; i++) {
        if (buckets[i]) {
            free(buckets[i]);
        }
    }
}

void partTwo() {

    FILE* inputFile = fopen(INPUT_FILE_1, "r");
    FILE* inputFile2 = fopen(INPUT_FILE_1, "r");

    numPair* buckets[BUCKET_COUNT] = { NULL };
    for (char buffer[MAX_CHAR]; fgets(buffer, MAX_CHAR, inputFile);) {
        int num1 = atoi(buffer);
        for (char buffer2[MAX_CHAR]; fgets(buffer2, MAX_CHAR, inputFile2);) {
            int num2 = atoi(buffer2);
            int addition = num2 + num1;
            if (!buckets[addition]) {
                buckets[addition] = createNumPair(num1, num2);
            }
        }
        fseek(inputFile2, 0, SEEK_SET);
    }

    for (char buffer[MAX_CHAR]; fgets(buffer, MAX_CHAR, inputFile2);) {
        int num = atoi(buffer);
        int searchVal = 2020 - num;
        if (buckets[searchVal]) {
            printf("Nums: %d, %d, %d\n", buckets[searchVal]->num1, buckets[searchVal]->num2, num);
            printf("Answer: %d\n", buckets[searchVal]->num1 * buckets[searchVal]->num2 * num);
            break;
        }
    }

    freePairs(buckets, BUCKET_COUNT);
    fclose(inputFile);
    fclose(inputFile2);
}

int main() {
    printf("Part one:\n");
    partOne();
    printf("\nPart two:\n");
    partTwo();
    return 0;
}