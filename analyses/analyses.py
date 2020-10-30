import glob
# import os
import subprocess
import csv
import sys


# カレントディレクトリのファイルを検索し，ファイル行数を取得する．
PATH = '/usr/local/opt/kani/analyses/'

def search_diff_line():
  # 差分のあるファイルを取得し，csvで出力する．
  result = subprocess.run("git diff --numstat | awk -v OFS=, '{print $1,$2,$3}' > /usr/local/opt/kani/analyses/diff_lines.csv", shell=True, text=True)

def search_analyses_file():
  # commitすべき状態か判断する．

  with open(f'{PATH}diff_lines.csv') as data:
    lines = csv.reader(data)
    criteria = 3 # 判定基準
    flag = search_error_count() # 基準を超えたものが有る場合，commitにいついての表示を出すため1にする．
    for line in lines:
      total_line = int(line[0])+int(line[1])
      if total_line > criteria:
        show_message(line[2],total_line,criteria)
        flag = 1
      
    if flag == 1:
      f = open(f'{PATH}guide_commit.txt','r',encoding='UTF-8')
      data = f.read()
      print(data)
      f.close()

def show_message(file_name,total_line,criteria):
  # commitを提案する行数を超えたら，メッセージを表示する．
  print(f'{file_name}の編集した行数が{total_line}行になったので，Commitをお勧めします．')

def search_error_count():
  # エラーのカウント数を受け取り，表示の処理を行う．
  count = int(sys.argv[1])
  if count >= 3:
    print('連続したエラーが修正されたので，Commitをお勧めします．')
    return 1  
  else:
    return 0

search_diff_line()
search_analyses_file()