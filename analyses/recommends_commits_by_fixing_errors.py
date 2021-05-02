#! /usr/bin/env python

import datetime
import os
import sqlite3
import sys

class Command:
    def __init__(self, datetime, command, status):
        self.datetime = datetime
        self.command = command
        self.status = status

targetCommands = ['javac', 'gcc', 'clang', 'make', 'maven', 'gradle', 'go build']

def isTargetCommand(command):
    for target in targetCommands:
        if command.startswith(target):
            return True
    return False

def readsCommandHistoryFromDB(projectDir):
    dbPath = os.path.join(projectDir, ".kani", "kani.sqlite")
    conn = sqlite3.connect(dbPath, detect_types=sqlite3.PARSE_DECLTYPES|sqlite3.PARSE_COLNAMES)
    sqlite3.dbapi2.converters['DATETIME'] = sqlite3.dbapi2.converters['TIMESTAMP']
    cursor = conn.cursor()
    commands = list()

    for row in cursor.execute("SELECT * FROM histories WHERE datetime > datetime('now', '-1 hours') ORDER BY datetime DESC"):
        if not isTargetCommand(row[2]):
            continue
        commands.append(Command(row[1], row[2], row[3]))
    conn.close()
    return commands

home = os.environ['KANI_HOME']
projectDir = os.environ['KANI_PROJECT_DIR']
branch = os.environ['KANI_CURRENT_BRANCH']
revision = os.environ['KANI_CURRENT_REVISION']
commands = readsCommandHistoryFromDB(projectDir)

sys.exit(0)
