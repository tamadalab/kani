#! /usr/bin/env python3

import time
import os
import csv
import sys

recommend_message = f'Recommend "git add" since it looks the working program ran well.'
recommend_message_ja = f'作業中のプログラムのテストが終われば，"git add"を実行しましょう．'

targetCommands = ['javac', 'gcc', 'clang', 'go build']

home = os.environ['KANI_HOME']
projectDir = os.environ['KANI_PROJECT_DIR']
branch = os.environ['KANI_CURRENT_BRANCH']
revision = os.environ['KANI_CURRENT_REVISION']
prevcmd = os.environ['KANI_PREV_COMMANDS']
status = int(os.environ['KANI_PREV_COMMANDS_STATUS'])

class Target:
    def __init__(self, target_command, execfile_parser):
        self.target_command = target_command
        self.execfile_parser = execfile_parser

    def is_target(self, command):
        return command.startswith(self.target_command + " ")

    def store_path(self):
        parent = os.path.join(projectDir, ".kani", 'compilations')
        if not os.path.exists(parent):
            os.makedirs(parent)
        path = os.path.join(parent, self.target_command.replace(' ', '_') + '.csv')
        return path

    def execfile(self, command):
        return self.execfile_parser(command)

def execfile_java_parser(command):
    cs = command.split()
    for i in cs:
        if i.endswith('.java'):
            return i.replace('.java', '')

def execfile_default_parser(command):
    index = command.find('-o ')
    if index < 0:
        return 'a.out'
    cs = command.split(' ')
    for i in range(0, len(cs)):
        if cs[i] == '-o':
            return cs[i + 1]
    return 'a.out'

def store_cmd(target, prevcmd):
    with open(target.store_path(), mode='a') as f:
        w = csv.writer(f)
        w.writerow([time.time(), prevcmd, status])

def store_exec(target, prevcmd):
    with open(os.path.join(projectDir, '.kani', 'compilations', 'execfile'), mode='w') as f:
        f.writelines([target.execfile(prevcmd)])

def find_exec_filename():
    path = os.path.join(projectDir, '.kani', 'compilations', 'execfile')
    if os.path.exists(path):
        with open(path, mode='r') as f:
            return f.readline()
    return None

def is_recommendation_mode_on():
    path = os.path.join(projectDir, '.kani', 'compilations', 'recommends')
    return os.path.exists(path)

def update_recommendation_mode(mode):
    path = os.path.join(projectDir, '.kani', 'compilations', 'recommends')
    if mode and not os.path.exists(path):
        with open(path, mode='w') as f:
            f.write('')
    elif os.path.exists(path):
        os.remove(path)

def is_exec():
    execfile = find_exec_filename()
    if execfile is None:
        return False
    return prevcmd.find(execfile) >= 0

def recommends_add():
    print(recommend_message_ja)
    sys.exit(1)

def reset_all():
    parent = os.path.join(projectDir, '.kani', 'compilations')
    recommends = os.path.join(parent, 'recommends')
    if os.path.exists(recommends):
        os.remove(recommends)
    execfile = os.path.join(parent, 'execfile')
    if os.path.exists(execfile):
        os.remove(execfile)


targets = list(filter(lambda x: x.is_target(prevcmd),
    map(lambda x: Target(x[0], x[1]), [('javac', execfile_java_parser), ('gcc', execfile_default_parser), ('clang', execfile_default_parser), ('go build', execfile_default_parser)])))

if is_exec():
    # TODO 終了ステータスが0なら，add を推薦．
    # 推薦状態を on．終了ステータスが0以外なら推薦状態を off．
    if status == 0:
        update_recommendation_mode(True)
        recommends_add()
    else:
        update_recommendation_mode(False)
elif len(targets) == 1:
    store_cmd(targets[0], prevcmd)
    store_exec(targets[0], prevcmd)
    update_recommendation_mode(False) # 推薦状態を off．
elif prevcmd.startswith('git add'):
    reset_all()
elif is_recommendation_mode_on():
    recommends_add()

sys.exit(0)
