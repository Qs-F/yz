# exp `yz`

This is the experimental repo for language processing software.

In this project, the last goal is to make markdown processor.  
`yz` mainly provides the way to deal with replacing model of proccess.

The core idea of `yz` is to build the ast along with inconsist expressions, like `# title` and `#title`.

In addition to this, the goal is to make the html products, so any parse error must be ignored. In other words, it should not be treated as error.

## Definition

pp = pre parse

## Design

### Dividing the Contents into Blocks

### Pre-divide into Small Blocks vs. Large Blocks

e.g. for small blocks:

```
1. one
2. two
3. three
```

interprets above into three node, afterwards, concatanates and makes one upper-layer.

e.g. for large blocks:

```
1. one
2. two
3. three
```

interprets above into one big node, afterwards, makes three lower-layer.
