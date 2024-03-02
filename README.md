# new-wc-go

*Technologies used:* ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![Git](https://img.shields.io/badge/git-%23F05033.svg?style=for-the-badge&logo=git&logoColor=white)

Implementing a replica of Unix' `wc` tool for the coding challenge: [click here](https://codingchallenges.fyi/challenges/challenge-wc)

## Requirements:

- Go installed and configured in PATH.

## Build

Execute the following command in the root directory of the project.

```go
    go build -o nwc .
```

Successful build will generate a binary in the root directory with name "nwc".

## Execution

1. To count number of bytes, execute the following command.

```bash
    ./nwc -c test.txt
```

2. To count number of lines, execute the following command.

```bash
    ./nwc -l test.txt
```

3. To count number of words, execute the following command.

```bash
    ./nwc -w test.txt
```

4. To count number of characters, execute the following command.

```bash
    ./nwc -m test.txt
```

**Note:** If no command-line options are passed, nwc will print count of lines, words and bytes by default. For example:

```bash
    ./nwc test.txt
```

**Note:** It is also possible to chain out of another command as an input for nwc. For example:

```bash
    cat test.txt | ./nwc
```

## Contribution

Please feel free to create pull requests if you like to add any feature. Also, feel free to take snippets of code from this project.