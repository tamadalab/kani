#include <stdio.h>

int total_num = 20;

// ターミナルにソートした結果を出力する
void show_data(int sort_num[total_num]){
  for(int i=0; i<total_num; i++){
    printf("%d ", sort_num[i]);
  }
  printf("\n");
}

// バブルソート
// 隣り合う要素の大小を比較しながら整列させていく．
void bubble_sort(int sort_num[total_num]){
  printf("バブルソート\n");
  int i,j,tmp;
  for(i=0; i<total_num; i++){
    for(j=total_num-1; j>i; j--){
      if(sort_num[j] < sort_num[j-1]){
        // 配列前後の値を入れ替える．
        // tmp = sort_num[j];
        // sort_num[j] = sort_num[j-1];
        // sort_num[j-1] = tmp;
      }
    }
  }
  show_data(sort_num);
}

// 素数を表示
// 素数だけを表示する
void prime_number(int num[total_num]){
  printf("素数を表示\n");
  int i,j,flag;
  int count=0;
  int prime_num[total_num];
  for(i=0; i<total_num; i++){
    flag = 0;
    // for文で回して，nの倍数だった場合flagを立てる．
    // for(j=2; j<num[i]; j++){
    //   if(num[i]%j == 0){
    //     flag = 1;
    //     break;
    //   }
    // }
    if(flag==0){
      prime_num[count] = num[i];
      count++;
    }
  }
  total_num = count;
  show_data(prime_num);
  bubble_sort(prime_num);
}


int main(void){

  // 以下データ読み込み用
  FILE *fp;
  int num[total_num];
  fp = fopen("data.txt", "r");
  if(fp == NULL){
    printf("ファイルをオープンできませんでした。\n");
    return 0;
  }
  int count = 0;
  while(fscanf(fp, "%d", &num[count]) != EOF){
    // printf("%d", num[i]); //デバッグ用
    count++;
  }
  fclose(fp);

  prime_number(num);

  return 0;
}