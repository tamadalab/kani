#include <stdio.h>

int total_num = 20;

// ターミナルにソートした結果を出力する
void show_data(int num[total_num]){
  for(int i=0; i<total_num; i++){
    printf("%d ", num[i]);
  }
  printf("\n");
}

// 選択ソート
// 最大値最初うちを探索し，配列最後の要素と入れ替えていく．
void selection_sort(int sort_num[total_num]){
  printf("選択ソート\n");
  int i,j,min,tmp,count=0;
  for(i=0; i<total_num; i++){
    min = sort_num[i];
    // for文で回して，minより小さい値があれば置き換える．
    // for(j=i+1; j<total_num; j++){
    //   if(sort_num[j] < min){
    //     min = j;
    //   }
    // }

    if(min != i){
      tmp = sort_num[i];
      sort_num[i] = sort_num[min];
      sort_num[min] = tmp;
      count++;
    }
  }
  show_data(sort_num);
}

// 素数以外を表示
// 素数以外を表示する
void prime_number(int num[total_num]){
  printf("素数以外を表示\n");
  int i,j,flag;
  int count=0;
  int prime_num[total_num];
  for(i=0; i<total_num; i++){
    flag = 0;
    for(j=2; j<num[i]; j++){
      if(num[i]%j == 0){
        flag = 1;
        break;
      }
    }
    // flagが1の時，prime_numに値を追加していく．
    // if(flag==1){
    //   prime_num[count] = num[i];
    //   count++;
    // }
  }
  total_num = count;
  show_data(prime_num);
  selection_sort(prime_num);
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