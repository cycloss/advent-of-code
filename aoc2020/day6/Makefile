
test: ./day6
	./day6

./day6: day6.c
	gcc -lGenericStructures day6.c -o day6

debug: day6.c
	gcc -g -lGenericStructures day6.c -o day6

leakCheck: ./day6
	valgrind ./day6
