# longest substring repeating character replace

You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.

 
## Examples

```
Example 1:

Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.
Example 2:

Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
 

Constraints:

1 <= s.length <= 105
s consists of only uppercase English letters.
0 <= k <= s.length
```

## 解析

要找出最長連續相同字元組成的子字串 並且可以替換 k 個字元

第1是 連續相同字元組成的子字串 必定是出現頻率最多次, 因為要替換的字元較少

思考當這個當這個 slide window 內的除去最多頻率的字元 所需替換的字元數 小於等於 k

這個 slide window 則是候選字元

舉例來說: 假設 s = "ABABBA" ,k =2

當 slide window 左界 = 0 , 右界 = 4 則 slide window 包含的字元為 "ABABB"

這時 出現最多次的字元是 B, 除去 B ， 需要替換的字元數為 2 所以是候選的 slide window

所以每次擴張 右界需要檢查條件就是  (right - left + 1） - 目前出現最多頻率的字元數 是否 <= k

而當 （right - left + 1) - 目前出現最多頻率的字元數 > k 時

需要移動左界 讓 slide window 符合條件

## 程式碼

```golang
package solution

func characterReplacement(s string, k int) int {
	hitMap := make(map[byte]int)
	left := 0
	count := 0
	result := 0
	for right := 0; right < len(s); right++ {
		hitMap[s[right]]++
		if hitMap[s[right]] > count {
			count = hitMap[s[right]]
		}
		for right-left+1-count > k {
			hitMap[s[left]]--
			left++
		}
		if right-left+1 > result {
			result = right - left + 1
		}
	}
	return result
}
```