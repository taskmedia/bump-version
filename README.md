# bump version

This minimal cli tool returns the next version depending on the given level (`patch`, `minor` or `major`)

## usage

```bash
$ ./bump-version -v "v1.2.3" -l "minor"
v1.3.0
```

## examples

Output versions on version `v1.2.3` on different levels given:

| level | version |
| ----- | ------- |
| patch | v1.2.4  |
| minor | v1.3.0  |
| major | v2.0.0  |
