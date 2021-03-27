---
title: ":red_car: How to work"
---

## Mechanism for recoding inside the `kani`

Using zsh hook mechanism for recoding.

## Directory layout of `kani`

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
    ├── find-project-dir.sh (script for finding the project root from the current directory)
    └── is-target-project.sh (script for checking the current directory is the target of kani)
```
