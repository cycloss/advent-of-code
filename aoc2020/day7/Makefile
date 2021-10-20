
test: ./day7
	./day7

./day7: day7.c
	gcc -lGenericStructures day7.c day7pt2.c day7aux.c -o day7

debug: day7.c
	gcc -g -lGenericStructures day7.c day7pt2.c day7aux.c -o day7

leakCheck: ./day7
	valgrind ./day7
