/*
Breadth First Search: Shortest Reach
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//int n: the number of nodes
//int m: the number of edges
//int edges[m][2]: start and end nodes for edges
//int s: the node to start traversals from
// Complete the bfs function below.
func bfs(n int32, m int32, edges [][]int32, s int32) []int32 {
	var dis []int32 //слайс расстояний
	var itr, num, numend int32
    for itr:=1; itr <= int(n-1); itr++{// заполним слайс растояний -1 для висящих точек. n-1 так как стартовую вершину не считаем
        dis = append(dis, -1)
    }

	for itr = 1; itr <= n; itr++ { // общай цикл по всем нодам

		for i:=0; i<=len(edges)-1; i++ { // запускаем цикл по всему слайсу для поиска ребер

			for _, num = range edges[i][:1] {  // берем из 1 позиции ноду от которой ищем
                 

            if edges[i] != nil && num == itr {  // берем в слайсе ребра только относящиеся к данной верхней ноде
                for _, numend = range edges[i][1:] {     

                    if itr>1 && dis[numend-3]>0 { // если во втором слое в исходной ноде есть вес то учитываем его 
                        dis [numend-2] = dis [numend-3]+6
                    } else{
                        dis [numend-2] = 6
                         } 
                         //fmt.Println("itr=",itr,"i =",i,"edges[i] = ", edges[i],"edges[i][:1] =",edges[i][:1],"edges[i][1:]=",edges[i][1:]," dis [itr] = ", dis[itr])
                    
                 } // ищем с кем соединяется ребро и присваиваем вес до этой ноды 
                
                }
            }       

		}

	}
	
    return dis // выводим вес маршрутов
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create("/Users/a18228591/go/src/github.com/chaosblade-exec-os/graph/log.log")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nm := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nm[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(nm[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		var edges [][]int32
		for i := 0; i < int(m); i++ {
			edgesRowTemp := strings.Split(readLine(reader), " ")

			var edgesRow []int32
			for _, edgesRowItem := range edgesRowTemp {
				edgesItemTemp, err := strconv.ParseInt(edgesRowItem, 10, 64)
				checkError(err)
				edgesItem := int32(edgesItemTemp)
				edgesRow = append(edgesRow, edgesItem)
			}

			if len(edgesRow) != int(2) {
				panic("Bad input")
			}

			edges = append(edges, edgesRow)
		}

		sTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		s := int32(sTemp)

		result := bfs(n, m, edges, s)

		for i, resultItem := range result {
			fmt.Fprintf(writer, "%d", resultItem)

			if i != len(result)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		fmt.Fprintf(writer, "\n")
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
