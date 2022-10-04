#include <stdio.h>
#include <stdlib.h>

int some_sum(int n, int m) {
  double s = 0;
  for(int i = 0; i <= n; i++)
    s += pow((double)i,(double)m);
  s = (int) s;
  return s;
}

int main(){
  FILE *S1, *S2;
  int n, m;
  
  S1 = fopen("data.txt", "r");
  fscanf(S1, "%d %d", &n, &m);
  fclose(S1);
  
  S2 = fopen("data.txt", "w"); 
  /*
  открытие существующего файла для записи
  стирает всё его содержимое
  */
  fprintf(S2, "%d\n", some_sum(n,m));
  fclose(S2);
  return 0;
}