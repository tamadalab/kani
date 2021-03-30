---
title: ":red_car: How to work"
---

## Mechanism for recoding inside the `kani`

Using bash/zsh hook mechanism for recoding.
Note that `kani` runs on `bash`, you must install [`rcaloras/bash-preexec`](https://github.com/rcaloras/bash-preexec).


## Directory layout of `kani`

### `KANI_HOME`

The directory layout of `KANI_HOME` is like below.

```sh
kani
├── README.md
├── analyses
│   ├── recommend.py (decision of recommendation)
│   └── commit_guide.txt (recommendation text)
├── bin
│   └── git-kani
└── scripts # utility scripts (includes hook function)
    ├── chpwd_hook.sh
    ├── periodic_hook.sh
    ├── precmd_hook.sh
    ├── preexec_hook.sh
    ├── find-project-dir.sh  # script for finding the project root from the current directory
    ├── is-target-project.sh # script for checking the current directory is the target of kani
    └── store_db.sh          # script for storing the command line histories into the database.
```

### `.kani` on the project root.

Also, `kani` creates `.kani` directory on the project root (same location as `.git` directory).
The `kani` stores various recorded data into `.kani` directory, such as command line histories (`kani.sqlite`), previous command (`prev_cmd`), and etc.

## Database

### Schema

| Name        | Type    | Primary | Null | Default           | Note |  
|-------------|---------|---------|------|-------------------|------|
| id          | INTEGER | Yes     | No   | N/A               |      |
| datetime    | TEXT    | No      | No   | CURRENT_TIMESTAMP | UTC  |
| command     | TEXT    | No      | No   | ""                |      |
| status_code | INTEGER | No      | No   | ""                |      |
| branch      | TEXT    | No      | Yes  |                   |      |
| revision    | TEXT    | No      | Yes  |                   |      |
| shell       | shell   | No      | No   |                   |      |

### `CREATE TABLE`

```sql
CREATE TABLE histories ( \
  id          INTEGER PRIMARY KEY, \
  datetime    TEXT    DEFAULT CURRENT_TIMESTAMP, \
  command     TEXT    NOT NULL, \
  status_code INTEGER NOT NULL, \
  branch      TEXT, \
  revision    TEXT, \
  shell       TEXT  \
)```
