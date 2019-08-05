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

### Isolated vs. Dependent Node

In general, node can be categorized into the two types:

#### Isolated Node -- Context-free Node

Isolated Node is the things which does not need context.

Like in markdown, the heading notation `#` or `##` can be the example.  
The reason is, if we write `# Title`, it should be interpreted as `<h1>Title</h1>` without knowing the context.

### Dependent Node -- Sequencial Context Node

Dependent Node is the things which need context.

Like in markdown, the list notation is:

```
1. one
2. two
3. three
5. four
```

must be interpreted as:

```
<ol>
  <li>one</li>
  <li>two</li>
  <li>three</li>
  <li>four</li>
</ol>
```
