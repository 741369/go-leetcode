/**********************************************
** @Des: 914 x-of-a-kind-in-a-deck-of-cards
** @Author: liuzhiwang
** @Last Modified time: 2020/3/27 下午7:05
***********************************************/

package main

import "fmt"

func main() {
	//arr := []int{1, 2, 3, 4, 4, 3, 2, 1}
	arr := []int{1, 1, 1, 2, 2, 2, 3, 3, 3}
	fmt.Println(hasGroupsSizeX(arr))

}

func hasGroupsSizeX(deck []int) bool {
	deckLen := len(deck)
	count := make([]int, 10000)
	for _, v := range deck {
		count[v]++
	}

	for i := 2; i <= deckLen; i++ {
		if deckLen%i == 0 {
			flag := true
			for _, v := range count {
				if v%i != 0 {
					flag = false
					break
				}
			}
			if flag {
				return true
			}
		}
	}

	return false
}

/**
class Solution {
    public boolean hasGroupsSizeX(int[] deck) {
        int N = deck.length;
        int[] count = new int[10000];
        for (int c: deck)
            count[c]++;

        List<Integer> values = new ArrayList();
        for (int i = 0; i < 10000; ++i)
            if (count[i] > 0)
                values.add(count[i]);

        search: for (int X = 2; X <= N; ++X)
            if (N % X == 0) {
                for (int v: values)
                    if (v % X != 0)
                        continue search;
                return true;
            }

        return false;
    }
}
*/
