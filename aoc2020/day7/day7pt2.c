#include "day7.h"

extern hashMap* bagMap;
extern arrayList* bagList;

int findBagsWithinBag(char* bagName, int multiplier) {

    arrayList* currentBagList = (arrayList*)getValueForKey(bagMap, bagName);

    if (!currentBagList) {
        printf("No bags found for key: %s\n", bagName);
        return 0;
    }

    int total = 0;

    for (int i = 0; i < getSize(currentBagList); i++) {
        bag* b = getItemAt(currentBagList, i);
        total += b->count * multiplier;
        total += multiplier * findBagsWithinBag(b->color, b->count);
    }

    return total;
}