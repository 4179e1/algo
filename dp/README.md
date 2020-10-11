01背包问题的dptable迭代过程，输入参数为
costs  = [1, 1, 1, 1, 2]
values = [1, 2, 3, 4, 5]
volume = 3

Row表示item (costs/values)
Col表示容量

```
        Col 0   Col 1   Col 2   Col 3
Row 0:  0       0       0       0
Row 1:  0       1       0       0
Row 2:  0       0       0       0
Row 3:  0       0       0       0
Row 4:  0       0       0       0
Row 5:  0       0       0       0

        Col 0   Col 1   Col 2   Col 3
Row 0:  0       0       0       0
Row 1:  0       1       1       0
Row 2:  0       0       0       0
Row 3:  0       0       0       0
Row 4:  0       0       0       0
Row 5:  0       0       0       0

        Col 0   Col 1   Col 2   Col 3
Row 0:  0       0       0       0
Row 1:  0       1       1       1
Row 2:  0       0       0       0
Row 3:  0       0       0       0
Row 4:  0       0       0       0
Row 5:  0       0       0       0

        Col 0   Col 1   Col 2   Col 3
Row 0:  0       0       0       0
Row 1:  0       1       1       1
Row 2:  0       2       0       0
Row 3:  0       0       0       0
Row 4:  0       0       0       0
Row 5:  0       0       0       0

        Col 0   Col 1   Col 2   Col 3
Row 0:  0       0       0       0
Row 1:  0       1       1       1
Row 2:  0       2       3       0
Row 3:  0       0       0       0
Row 4:  0       0       0       0
Row 5:  0       0       0       0

        Col 0   Col 1   Col 2   Col 3
Row 0:  0       0       0       0
Row 1:  0       1       1       1
Row 2:  0       2       3       3
Row 3:  0       0       0       0
Row 4:  0       0       0       0
Row 5:  0       0       0       0

        Col 0   Col 1   Col 2   Col 3
Row 0:  0       0       0       0
Row 1:  0       1       1       1
Row 2:  0       2       3       3
Row 3:  0       3       0       0
Row 4:  0       0       0       0
Row 5:  0       0       0       0

        Col 0   Col 1   Col 2   Col 3
Row 0:  0       0       0       0
Row 1:  0       1       1       1
Row 2:  0       2       3       3
Row 3:  0       3       5       0
Row 4:  0       0       0       0
Row 5:  0       0       0       0

        Col 0   Col 1   Col 2   Col 3
Row 0:  0       0       0       0
Row 1:  0       1       1       1
Row 2:  0       2       3       3
Row 3:  0       3       5       6
Row 4:  0       0       0       0
Row 5:  0       0       0       0

        Col 0   Col 1   Col 2   Col 3
Row 0:  0       0       0       0
Row 1:  0       1       1       1
Row 2:  0       2       3       3
Row 3:  0       3       5       6
Row 4:  0       4       0       0
Row 5:  0       0       0       0

        Col 0   Col 1   Col 2   Col 3
Row 0:  0       0       0       0
Row 1:  0       1       1       1
Row 2:  0       2       3       3
Row 3:  0       3       5       6
Row 4:  0       4       7       0
Row 5:  0       0       0       0

        Col 0   Col 1   Col 2   Col 3
Row 0:  0       0       0       0
Row 1:  0       1       1       1
Row 2:  0       2       3       3
Row 3:  0       3       5       6
Row 4:  0       4       7       9
Row 5:  0       0       0       0

        Col 0   Col 1   Col 2   Col 3
Row 0:  0       0       0       0
Row 1:  0       1       1       1
Row 2:  0       2       3       3
Row 3:  0       3       5       6
Row 4:  0       4       7       9
Row 5:  0       4       0       0

        Col 0   Col 1   Col 2   Col 3
Row 0:  0       0       0       0
Row 1:  0       1       1       1
Row 2:  0       2       3       3
Row 3:  0       3       5       6
Row 4:  0       4       7       9
Row 5:  0       4       7       0

        Col 0   Col 1   Col 2   Col 3
Row 0:  0       0       0       0
Row 1:  0       1       1       1
Row 2:  0       2       3       3
Row 3:  0       3       5       6
Row 4:  0       4       7       9
Row 5:  0       4       7       9

        Col 0   Col 1   Col 2   Col 3
Row 0:  0       0       0       0
Row 1:  0       1       1       1
Row 2:  0       2       3       3
Row 3:  0       3       5       6
Row 4:  0       4       7       9
Row 5:  0       4       7       9
```