/**********************************************
** @Des:
** @Author: lzw
** @Last Modified time: 2020/5/11 下午8:51
***********************************************/

package main

import "fmt"

func main() {
	fmt.Println("====", getLeastNumbers([]int{1, 3, 5, 2, 4}, 3))
}

func getLeastNumbers(arr []int, k int) []int {
	result := make([]int, k)
	if k < 0 || len(arr) > 10000 {
		return nil
	}
	quickSort(arr, 0, len(arr)-1)

	for i := 0; i < k; i++ {
		result[i] = arr[i]
	}

	return result
}

func quickSort(arr []int, left, right int) {
	if left < right {
		index := arr[left]
		low, high := left, right
		for low < high {
			for low < high && index <= arr[high] {
				high--
			}
			arr[low] = arr[high]
			for low < high && index >= arr[low] {
				low++
			}
			arr[high] = arr[low]
		}
		arr[low] = index
		quickSort(arr, left, low)
		quickSort(arr, low+1, right)
	}
}

/**
class Solution {
    public int[] getLeastNumbers(int[] arr, int k) {
        int[] result = new int[k];
        if(k<0 || arr.length>10000){
            return null;
        }
        quickSort(arr,0,arr.length - 1);

        for(int i = 0;i<k; i++){
            result[i] = arr[i];
        }
        return result;
    }

    public static void quickSort(int[] arr,int left,int right){
        if(left<right){
            int index = arr[left];
            int low = left,high = right;

            while(low<high){
                while(low<high && index <= arr[high]){
                    high --;
                }
                arr[low] = arr[high];
                while(low<high && index >= arr[low]){
                    low++;
                }
                arr[high] = arr[low];
            }
            arr[low] = index;

            quickSort(arr, left, low);
            quickSort(arr, low + 1,right);
        }
    }
}
*/
