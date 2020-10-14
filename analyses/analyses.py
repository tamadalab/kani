import glob
import os
# import subprocess

# カレントディレクトリのファイルを検索し，ファイル行数を取得する．


def search_analyses_file():
  # ファイルの検索する用のパスを返す．
  path=os.getcwd()

  # files = os.listdir(path)
  # file = [f for f in files if os.path.isfile(os.path.join(path, f))]
  # print(file)

  path += '/*.txt' # デバッグ状態 本番環境は.cに変更
  return path


def search_line_count(path):
  # ファイル行数を出力
  
  for f in glob.glob(path):
    print(os.path.split(f)[1])
    print(sum([1 for _ in open(os.path.split(f)[1])]))
  return


path = search_analyses_file()
search_line_count(path)