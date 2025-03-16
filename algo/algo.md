
# Algorithm in LeetCode


Summarized by Wafer

## Reference
[代码随想录](https://programmercarl.com/)

[算法通关手册](https://algo.itcharge.cn/00.Introduction/01.Data-Structures-Algorithms/)

[AC算法笔记](https://github.com/AlanChaw/interview/blob/master/%E7%AE%97%E6%B3%95%E7%AC%94%E8%AE%B0.md)

## Binary Search
#### 二分查找
搜索无重复、有序的数组


#### 例题
[704. Binary Search](https://leetcode.com/problems/binary-search/description/)





## Two Pointer
#### 快慢指针
删除数组特定值的元素，等价于创建不等于目标值的慢数组

#### 对撞指针
满足无重复、有序的数组又同时满足回文条件（以0对称）


#### 例题
[27. Remove Element](https://leetcode.com/problems/remove-element/description/)

[977. Squares of a Sorted Array](https://leetcode.com/problems/squares-of-a-sorted-array/description/)


## Sliding Window
#### 滑动窗口
数组的子数组问题，换算成求窗口大小

[209. Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/description/)


## Matrix
#### 矩阵
主要是学会go怎么创建二维切片

    mat := make([][]int,n)
    for i := range mat {
        mat[i] = make([]int,n)
    }

#### 例题
[59. Spiral Matrix II](https://leetcode.com/problems/spiral-matrix-ii/description/)


## Linked List
#### 链表
链表的插入删除等操作最好是遍历到前一个`Node`，因此引入 `dummyHead` 方便遍历

如果是从dummyHead开始遍历，完全便利的结束条件是 `curr.Next != nil`

如果从`head`开始遍历，因为真实的节点是 `cur!= nil`

实在想不出来就画图

#### 例题
[203. Remove Linked List Elements](https://leetcode.com/problems/remove-linked-list-elements/description/)

[707. Design Linked List](https://leetcode.com/problems/design-linked-list/description/)


## Linked List Two Ptr
#### 链表指针
快慢指针，快的先走n步，用于删除节点。

快慢指针，反转节点，一定要记住从dummyHead开始走，会方便很多；还有一个思路是temp来存可能会丢掉的值
#### 例题
[206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/description/)

[24. Swap Nodes in Pairs](https://leetcode.com/problems/swap-nodes-in-pairs/description/)

[19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/)

[160. Intersection of Two Linked Lists](https://leetcode.com/problems/intersection-of-two-linked-lists/description/)

## Hash Table
#### slices包
这个包巨好用，点击 [此链接](https://github.com/Wafer233/testLab/blob/master/golang/slices/slices.go)
#### 切片与映射
哈希表的在golang中的表达无非两种，slice和map，那么golang在遍历他们会发生什么？

    for index, num := range nums {
        //遍历切片
    }

    for key, value := range mapping {
        //遍历映射
    }

    if value, exist := map[key]; !exist {
        //不存在时会怎么样
    }

#### 哈希表
哈希表可以用于计数，比如说在求两数之和，可以把遍历过的值存进去。

前两题直接秒了，`key` 和`value`分别是对应的字符或者数值。

2sum, 4sumII 还可以，写循环，遍历过的数字存进哈希表，后续查有没有对应的key就好了。

3sum，4sum，就双指针了，遍历start，一共要总长度-3次；然后剩下两个不就是两数之和问题吗   
#### 例题
[242. Valid Anagram](https://leetcode.com/problems/valid-anagram/description/)

[349. Intersection of Two Arrays](https://leetcode.com/problems/intersection-of-two-arrays/description/)

[1. Two Sum](https://leetcode.com/problems/two-sum/description/)

[454. 4Sum II](https://leetcode.com/problems/4sum-ii/description/)

[15. 3Sum](https://leetcode.com/problems/3sum/description/)

[18. 4Sum](https://leetcode.com/problems/4sum/description/)

[347. Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/description/)


## String
#### 字符串
字符串类型也是力扣上的常考的东西了，golang的字符串有一个特点，就是他不能改，然后它还能遍历，通常来说要有这些操作。

    for index, ch := range str {
        ch == '1'
    }

    b := []byte(str)
    for index, value := range b {
        b[index] = 'a'
    }

    r := []rune(str)
    for index, value := range b {
        r[index] = 'a'
    }



#### strings包
这个包巨好用，点击 [此链接](https://github.com/Wafer233/testLab/blob/master/golang/strings/strings.go)

#### 反转与重复子串
看到反转都要有条件反射了！对撞指针！golang交换很简单

重复子串，这个很新，移动匹配，增加一段掐头去尾能找到自己
    
    s[left], s[right] =  s[right],s[left]   

#### 例题
[344. Reverse String](https://leetcode.com/problems/reverse-string/description/)

[541. Reverse String II](https://leetcode.com/problems/reverse-string-ii/description/)

[151. Reverse Words in a String](https://leetcode.com/problems/reverse-words-in-a-string/description/)

[459. Repeated Substring Pattern](https://leetcode.com/problems/repeated-substring-pattern/description/)

## Stack
#### 栈和队列
其实这个东西在go中很好理解，一个先进先出，一个先进后出，因此想要模拟的话直接操作切片

    //stack
    push -> append(stack,val)
    pop -> stack[:len(stack)-1]
    peek -> stack[len(stack)-1]

    //queue
    inQueue -> append(queue,val)
    deQueue -> stack[1:]
    peek -> stack[0]

#### 实现功能
栈实现括号有效，逆波兰表达式

*双向队列实现滑窗最大值

#### 例题
[232. Implement Queue using Stacks](https://leetcode.com/problems/implement-queue-using-stacks/)

[225. Implement Stack using Queues](https://leetcode.com/problems/implement-stack-using-queues/description/)

[20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/description/)

[1047. Remove All Adjacent Duplicates In String](https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/description/)

[150. Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/description/)

[239. Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/description/)



## Binary Tree
#### 二叉树
无非就是操作在夹在递归函数中的位置，表示了前中后遍历，注意层序遍历要按照前序的方式进行递归

递归遍历三部曲

1. 递归函数携带参数/返回值
2. 中止条件
3. 单层递归逻辑

---
   ####  
这里的写法是我个人总结出来的。

1.只要是树的遍历就用递归

2.递归当中有三个变量，一个是递归函数携带的参数，一个是递归函数的返回值，还有一个是递归函数外的一个全局的值。

3.返回值。我目前的想法就是，返回我所需要的属性值，画一个图就很好理解，通过后序把最下面的值一点点传到前面来。例如深度（可以用在求最大或者最小深度），是否对称的bool值.

4.携带值。

5.全局值。就是我想要打印出来所有的结果，这是就搞一个全局值，例如打印所有节点 `[]int`, 打印所有路径 `[]string`


#### 例题
[144. Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/description/)

[145. Binary Tree Postorder Traversal](https://leetcode.com/problems/binary-tree-postorder-traversal/)

[94. Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/description/)

[102. Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/description/)

## Binary Tree Ops and Set

### 方法论
构造无脑前序。

一个是用数组构造，核心就是找分割点。


### 小结
#### 翻转二叉树

递归：前序，交换左右孩子

#### 构造二叉树

递归：前序，重点在于找分割点，分左右区间构造

#### 构造最大的二叉树

递归：前序，分割点为数组最大值，分左右区间构造

#### 合并两个二叉树

递归：前序，左右一起遍历



### 例题


[226. Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/description/)

[106. Construct Binary Tree from Inorder and Postorder](https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/)

[654. Maximum Binary Tree](https://leetcode.com/problems/maximum-binary-tree/description/)

[617. Merge Two Binary Trees](https://leetcode.com/problems/merge-two-binary-trees/description/)





## Binary Tree Attribute
### 方法论
对于一般的二叉树，求属性基本上都是 __后续__，最简单的例子，你要求他是是不是对称，求他的深度，都要从后面不断地返回值上来.

当然也有例外，什么时候 __前序__ 呢？我认为是你要按顺序打印到一个全局的 `[]string` 里面，就要前序了


### 小结
#### 二叉树：是否对称

递归：后序，比较的是根节点的左子树与右子树是不是相互翻转

`compare(left, right *TreeNode) bool`

#### 二叉树：求最大深度

递归：后序，返回深度

`getD(root *TreeNode) int`


#### 二叉树：求最小深度

递归：后序，返回深度，深度要分情况讨论

`getD(root *TreeNode) int`


#### 二叉树：求有多少个节点
递归：后序，返回节点个数，因为可以简算成`2^n-1`


#### 二叉树：是否平衡
递归：后序，返回深度，不满足的情况返回-1


#### 二叉树：找所有路径 **

递归：前序，这个不一样，因为你打印的中间节点就得在前面，所以得前序


#### 二叉树：求左叶子之和

递归：后序，必须三层约束条件，才能判断是否是左叶子。

#### 二叉树：求左下角的值

迭代：层序遍历找最后一行最左边

#### 二叉树：求路径总和

递归：有一个很吊的写法，这样子就可以只会传对的值
`return travel(cur.Left, curSum) || travel(cur.Right, curSum)`

[101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree/description/)

[104. Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/description/)

[111. Minimum Depth of Binary Tree](https://leetcode.com/problems/minimum-depth-of-binary-tree/description/)

[222. Count Complete Tree Nodes](https://leetcode.com/problems/count-complete-tree-nodes/description/)

[110. Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/description/)

[257. Binary Tree Paths](https://leetcode.com/problems/binary-tree-paths/description/)

[404. Sum of Left Leaves](https://leetcode.com/problems/sum-of-left-leaves/description/)

[513. Find Bottom Left Tree Value](https://leetcode.com/problems/find-bottom-left-tree-value/description/)

[112. Path Sum](https://leetcode.com/problems/path-sum/description/)


## BST

[700. Search in a Binary Search Tree](https://leetcode.com/problems/search-in-a-binary-search-tree/description/)

[98. Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/description/)

[530. Minimum Absolute Difference in BST](https://leetcode.com/problems/minimum-absolute-difference-in-bst/description/)

[501. Find Mode in Binary Search Tree](https://leetcode.com/problems/find-mode-in-binary-search-tree/description/)

[236. Lowest Common Ancestor of a Binary Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/)

[701. Insert into a Binary Search Tree](https://leetcode.com/problems/insert-into-a-binary-search-tree/description/)

[538. Convert BST to Greater Tree](https://leetcode.com/problems/convert-bst-to-greater-tree/description/)

## Backtracking

### 解决的问题
    组合问题：N个数里面按一定规则找出k个数的集合
    切割问题：一个字符串按一定规则有几种切割方式
    子集问题：一个N个数的集合里有多少符合条件的子集
    排列问题：N个数按一定规则全排列，有几种排列方式
    棋盘问题：N皇后，解数独等等


### 三部曲
    回溯函数模板返回值以及参数
    回溯函数终止条件
    回溯搜索的遍历过程
## Backtracking Combination

###
```
func combine(n int, k int) [][]int {
	ret := [][]int{}
    curComb := []int{}

	var backtracking func(startNum int)
	backtracking = func(startNum int) {

		if end condition {
            temp := append([]int{}, curComb...)
			ret = append(ret, temp)
			return
		}
        

		for i := startNum; i <= n ; i++ {
			curComb = append(curComb, nums[i])
			backtracking(i+1)  <- 这里看能不能重复
			curComb = curComb[:len(curComb)-1]
		}
		return
	}

	backtracking(0)
	return ret
}
```

### 小结

#### 组合
不能重复，`i+1`开开始递归，结束是 `len(curComb) == k` 
#### 组合 II
不能重复，`i+1`开开始递归，结束是 `len(curComb) == k && sum == n`

#### 打电话
不能重复，`i+1`开开始递归，结束就是 `index == len(digits)`,本质上没啥区别，只是`for`循环是`letters`的`len`，递归是`index+1`

#### 组合和
可以重复，`i`开开始递归
#### 组合和II
不能重复，`i+1`开开始递归，要去重，在`for`外面写一个`used := map[int]bool{}`

### 例题
[77. Combinations](https://leetcode.com/problems/combinations/description/)

[216. Combination Sum III](https://leetcode.com/problems/combination-sum-iii/description/)

[17. Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/)

[39. Combination Sum](https://leetcode.com/problems/combination-sum/description/)

[40. Combination Sum II](https://leetcode.com/problems/combination-sum-ii/description/)


## Backtracking Segmentation
### 模板

    func partition(s string) [][]string {
        
        ret := [][]string{}
        cur := []string{}
        var backtracking func(Index int)
        backtracking = func(Index int) {
            if Index ==  len(s) {
                temp := append([]string{}, cur...)
                ret = append(ret,temp)
                return
            }

            for i := Index; i< len(s); i++{
                // [Index,i+1])
                str := s[Index:i+1]
                if isPali(str) {
                    cur = append(cur,str)
                } else {
                    continue
                }
                backtracking(i+1)
                cur = cur[:len(cur)-1]
            }
        }
        backtracking(0)
        return ret
    }

func isPali(str string) bool {
    left := 0
    right := len(str) -1
    for left < right {
        if str[left] != str[right] {
            return false
        } else {
            left ++
            right --
        }
    }
    return true
}

### 例题
[131. Palindrome Partitioning](https://leetcode.com/problems/palindrome-partitioning/description/)

[93. Restore IP Addresses](https://leetcode.com/problems/restore-ip-addresses/description/)

### 例题
## Backtracking Subset

[78. Subsets](https://leetcode.com/problems/subsets/description/)

[90. Subsets II](https://leetcode.com/problems/subsets-ii/description/)

[491. Non-decreasing Subsequences](https://leetcode.com/problems/non-decreasing-subsequences/description/)

## Backtracking Permutation

[46. Permutations](https://leetcode.com/problems/permutations/description/)

[47. Permutations II](https://leetcode.com/problems/permutations-ii/description/)

## Backtracking Chessboard

[51. N-Queens](https://leetcode.com/problems/n-queens/description/)

## DP

[509. Fibonacci Number](https://leetcode.com/problems/fibonacci-number/description/)

[70. Climbing Stairs](https://leetcode.com/problems/climbing-stairs/description/)

[746. Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/description/)

[62. Unique Paths](https://leetcode.com/problems/unique-paths/description/)

[63. Unique Paths II](https://leetcode.com/problems/unique-paths-ii/description/)

[343. Integer Break](https://leetcode.com/problems/integer-break/description/)

[96. Unique Binary Search Trees](https://leetcode.com/problems/unique-binary-search-trees/description/)

## DP 0-1 Knapsack

[416. Partition Equal Subset Sum](https://leetcode.com/problems/partition-equal-subset-sum/description/)

[1049. Last Stone Weight II](https://leetcode.com/problems/last-stone-weight-ii/description/)

[494. Target Sum](https://leetcode.com/problems/target-sum/description/)

[474. Ones and Zeroes](https://leetcode.com/problems/ones-and-zeroes/description/)

## DP Complete Knapsack

[518. Coin Change II](https://leetcode.com/problems/coin-change-ii/description/)

[377. Combination Sum IV](https://leetcode.com/problems/combination-sum-iv/description/)

[322. Coin Change](https://leetcode.com/problems/coin-change/description/)

[279. Perfect Squares](https://leetcode.com/problems/perfect-squares/description/)

[139. Word Break](https://leetcode.com/problems/word-break/description/)

## DP House Robber

[198. House Robber](https://leetcode.com/problems/house-robber/description/)

[213. House Robber II](https://leetcode.com/problems/house-robber-ii/description/)

[337. House Robber III](https://leetcode.com/problems/house-robber-iii/description/)

## DP Stock

[121. Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)

[122. Best Time to Buy and Sell Stock II](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/)

[123. Best Time to Buy and Sell Stock III](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/description/)

[188. Best Time to Buy and Sell Stock IV](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/description/)

[309. Best Time to Buy and Sell Stock with Cooldown](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/)

[714. Best Time to Buy and Sell Stock with Transa](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/)

## DP Discontinuous Subsequence

[300. Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/)

[1143. Longest Common Subsequence](https://leetcode.com/problems/longest-common-subsequence/description/)

[1035. Uncrossed Lines](https://leetcode.com/problems/uncrossed-lines/description/)

[392. Is Subsequence](https://leetcode.com/problems/is-subsequence/description/)

## DP Continuous Subsequence/ Subarray

[674. Longest Continuous Increasing Subsequence](https://leetcode.com/problems/longest-continuous-increasing-subsequence/description/)

[718. Maximum Length of Repeated Subarray](https://leetcode.com/problems/maximum-length-of-repeated-subarray/description/)

[53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray/description/)

## DP Edit Distance

[115. Distinct Subsequences](https://leetcode.com/problems/distinct-subsequences/description/)

[583. Delete Operation for Two Strings](https://leetcode.com/problems/delete-operation-for-two-strings/description/)

[72. Edit Distance](https://leetcode.com/problems/edit-distance/description/)

## DP Palindrome Array

[647. Palindromic Substrings](https://leetcode.com/problems/palindromic-substrings/description/)

[516. Longest Palindromic Subsequence](https://leetcode.com/problems/longest-palindromic-subsequence/description/)


