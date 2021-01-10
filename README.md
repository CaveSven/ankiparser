# ankiparser
Parse notes to create anki cards 

## Usage

Assume there is a text file with the following content:

```text
**Q:** What is the capital of France?
**A:** Paris

**Q:** From which city did Germany's first steam locomotive depart?
**A:** Nuremberg
```

Using 

```bash
ankiparser $PATH_TO_FILE $PATH_TO_TARGET_DIR
```

will create a Anki parsable file `$PATH_TO_TARGET_DIR/cards.txt` with the content

```text
What is the capital of France?\tParis
From which city did Germany's first steam locomotive depart?\tNuremberg
```


## GO HOW-TO

* [Tutorial: Create a Go module](https://golang.org/doc/tutorial/create-module)
