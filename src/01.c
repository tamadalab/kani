#include <stdio.h>
#define N 10

void show_data(int sort_num[N]){
  for(int i = 0; i<N; i++){
    printf("%d ",sort_num[i]);
  }
  printf("\n");
}

void bubble_sort(int sort_num[N]){
  printf("バブルソート\n");
  // 最後の要素を除いて、すべての要素を並べ替えます
  for(int i=0; i<N-1; i++){
    show_data(sort_num);
    // 下から上に順番に比較します
    for(int j=N-1; j>i; j--){
      // 上の方が大きいときは互いに入れ替えます
      if(sort_num[j]<sort_num[j-1]){
        int t = sort_num[j];
        sort_num[j] = sort_num[j-1];
        sort_num[j-1] = t;
      }
    }
  }
}

void select_sort(int sort_num[N]){
  printf("選択ソート\n");
  int max_index = 0;
  for(int i = 0;i < N-1 ; i++){   
    show_data(sort_num);
    max_index = i;
    //  現在の状況で、最大の値の入っているインデックスの位置を探す
    for(int j = i+1; j < N ; j++){
      if(sort_num[max_index] < sort_num[j]){
        max_index = j;
      }
    }
    //  先頭要素より、大きい値があれば、入れ替える。
    if(max_index != i){
      //  値の入れ替え
      int tmp = sort_num[i];
      sort_num[i] = sort_num[max_index];
      sort_num[max_index] = tmp;
    }
  }
}

int main(void){
  FILE *fp;
  int num[N];

  fp = fopen("data.txt", "r");

  if(fp == NULL){
    printf("ファイルをオープンできませんでした。\n");
    return 0;
  }
  int i = 0;
  while(fscanf(fp, "%d", &num[i]) != EOF){
    i++;
  }
  fclose(fp);

  bubble_sort(num);
  select_sort(num);

  return 0;
}