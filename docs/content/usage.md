---
title: ":runner: Usage"
---

## Initialization

Before using `kani`, write the following snippet into your `.zshrc`.

```sh
eval "$(git kani init -)"
```

## Enable `kani` on your git repository.

Run the following command in your git repository.

```sh
git kani init
```

Then, `kani` creates `.kani` folder on the top of your git repository (same as `.git` folder).
