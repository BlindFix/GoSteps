/*

Дан список неотрицательных целых чисел, повторяющихся элементов в списке нет. Нужно преобразовать его в строку, сворачивая соседние по числовому ряду числа в диапазоны.
Примеры:

    [1,4,5,2,3,9,8,11,0] => "0-5,8-9,11"
    [1,4,3,2] => "1-4"
    [1,4] => "1,4"
    [1,2] => "1-2"

*/
import (
    "sort"
    "strconv"
    "strings"
)
func compress(l []int) string {
    if len(l) == 0 {
        return ""
    }

    l = sort.Slice(l)

    var resSlice []string
    var nowCount, lElem int = l[0], 1

    for i := 1; i < len(l); i++ {
        if l[i] == l[i - 1] + 1 {
            nowCount++
        } else {
            if nowCount == 1 {
                resSlice = append(resSlice, strconv.Itoa(lElem))
            } else {
                resSlice = append(resSlice, strconv.Itoa(lElem) + "-" + strconv.Itoa(lElem + nowCount - 1))
            }
            lElem = l[i]
            nowCount = 1
        }
    }
    if nowCount == 1 {
        resSlice = append(resSlice, strconv.Itoa(lElem))
    } else {
        resSlice = append(resSlice, strconv.Itoa(lElem) + "-" + strconv.Itoa(lElem + nowCount - 1))
    }

    return strings.Join(resSlice, ",")
}

/*

На вход приходит список строк вида:

[
    "key.subkey.subkey2=1",
    "key.subkey=2",
    "key.subkey3=3",
    "key2.subkey4=5"
]

Преобразовать в структуру вида (все строки заполняют одну структуру):

    type Properties struct {
        Value *int
        Inner map[string]*Properties
    }

Сигнатура метода:
    func parse(input []string) *Properties
Формат всегда корректный, значение есть всегда.
Данные складываются вот так:

[
   "key": {
       "subkey": {
           "value": 2,
           "subkey2": {
               "value": 1
           }
       },
       "subkey3": {
           "value": 3
       }
   },
   "key2": {
       "subkey4": {
          "value": 5
       }
   }
]

*/

import (
    "sort"
    "strconv"
    "strings"
)

type Properties struct {
    Value *int
    Inner map[string]*Properties
}

func parse(input []string) *Properties {
    res := make(Properties)

    for i := 0; i < len(input); i++ {
        splitted := strings.Split(input[i], "=")
        nodes := strings.Split(splitted[0], ".")
        val := strconv.Atoi(splitted[0])

        nowNode := res
        for i := 0; i < len(nodes); i++ {
            if nowNode.Inner == nil {
                nowNode.Inner = make(map[string]*Properties)
            }
            nowNode = nowNode.Inner[nodes[i]]
        }

        nowNode.Value = new(int)
        *nowNode.Value = val
        // nowNode.Value = &[]int{val}[0]
    }

    return res
}
