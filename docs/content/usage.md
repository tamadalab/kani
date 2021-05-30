---
title: ":runner: Usage"
---

### :shoe: Initialization

At first, write the following snippet in your login shell initializer (`.bashrc`, or `.zshrc`)

```sh
eval "$(kani init -)"
```

### :running_shirt_with_sash: Setup repository

* Choose your `git` repositories to run `kani`,
* Change directory to your local `git` repository,
* Run `kani init` command for preparing the `kani` running on the repository.

### :coffee: Stop `kani` temporarily in a particular repository

If want to stop recoding command temporary, run `kani disable` (`kani` creates a `disable` file on `.kani` directory).
Then, want to start recoding again, run `kani enable`.

### :broken_heart: Stop `kani`  permanently in a particular repository

If typing `kani deinit` in a particular repository, kani never records anymore.
After that, typing `kani init`, `kani` starts recoding again.
