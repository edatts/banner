# Banner

This module contains a toy utility for printing text to the terminal in block ANSI format.

## Usage

As package:

```
b, err := banner.NewFromText("My text")
if err != nil {
	return err
}

b.Print()
```

```
b := banner.New()
if err := b.Fill("My text"); err != nil {
    return err
}

b.Print()
```

From command line:

```
```
