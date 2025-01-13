# Linear-Stats

The program is able to read from a file and print the `Linear Regression Line` formula and the `Pearson Correlation Coefficient` of the data provided. In other words, the program is be able to read the data present in the path passed as argument. The data in the file will be presented as the following example:

```console
234
189
113
121
344
145
125
114
110
```

This data represents a graph in which the values of the x axis are the number of the lines (0, 1, 2, 3, 4, 5 ...) and the values of the y axis are the actual numbers (189, 113, 121, 114, 145, 110...).

## *NOTE:*

*Make sure the file with the data is in the root directory i.e Linear Stats*

- To add the data to the program, create a file with the data and save it in the root directory of the program.
- To quickly set `0` as a `y` value, just add an empty space. Thus, all empty spaces will be treated as intentional `0` values.

To run the program: ` go run . main.go data.txt`

After reading the file, the program executes the Linear Regression and the Pearson Correlation Coefficient functions and prints the results in the following manner (the following expressions are only examples):

```console
Linear Regression Line: y = 0.001174x + -2252.335949
Pearson Correlation Coefficient: 0.0026653118
```

## References

[Linear Regression](https://en.wikipedia.org/wiki/Linear_regression)

[Pearson Correlation Coefficient](https://en.wikipedia.org/wiki/Pearson_correlation_coefficient)

## Author

[Adioz Daniel](https://github.com/adiozdaniel)
