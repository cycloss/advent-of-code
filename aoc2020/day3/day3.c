#include <stdio.h>
#include <stdlib.h>

int getTrees(int across, int down) {
    FILE* inFile = fopen("inputFiles/day3.txt", "r");
    int pos = 0, downp = 0, trees = 0;
    for (char buffer[33]; fgets(buffer, 33, inFile); downp++) {
        if (downp % down == 0) {
            (buffer[pos % 31] == '#') ? trees++ : trees;
            pos += across;
        }
    }
    fclose(inFile);
    return trees;
}

int main() {
    printf("Part one trees encountered: %d\n", getTrees(3, 1));
    int pt2 = getTrees(1, 1), vectors[][2] = { { 3, 1 }, { 5, 1 }, { 7, 1 }, { 1, 2 } };
    for (int i = 0; i < sizeof(vectors) / sizeof(vectors[0]); i++) {
        pt2 *= getTrees(vectors[i][0], vectors[i][1]);
    }
    printf("\nPart two trees encountered: %u\n", pt2);
}