# Linear Regression

Run this from the current directory:

```console
go run main.go
```

Enjoy ;)

### Example

Given dummy data of the type `y = 3x` the algorithm will learn the right coefficient of the linear equation using the
stochastic gradient descent.

```console
Epoch 99 | Loss 0.0000 | Learning coefficient: 3.00

1.9041 ┼───╮
1.7137 ┤   ╰───╮
1.5233 ┤       ╰─╮
1.3329 ┤         ╰──╮
1.1425 ┤            ╰──╮
0.9520 ┤               ╰──╮
0.7616 ┤                  ╰──╮
0.5712 ┤                     ╰───╮
0.3808 ┤                         ╰─────╮
0.1904 ┤                               ╰──────────╮
0.0000 ┤                                          ╰────────────
```