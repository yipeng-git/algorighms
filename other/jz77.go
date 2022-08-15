package other

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// jz6
func printListFromTailToHead(head *ListNode) []int {
	// write code here
	if head == nil {
		return []int{}
	}
	if head.Next == nil {
		return []int{head.Val}
	}
	return append(printListFromTailToHead(head.Next), head.Val)
}

// jz24
func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	if pHead == nil {
		return nil
	}
	curr := pHead
	next := curr.Next
	curr.Next = nil
	for {
		if next == nil {
			return curr
		}
		tmp := next.Next
		next.Next = curr
		curr = next
		next = tmp
	}

}

// jz25
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}
	var curr *ListNode
	if pHead1.Val < pHead2.Val {
		curr = pHead1
		pHead1 = pHead1.Next
	} else {
		curr = pHead2
		pHead2 = pHead2.Next
	}
	head := curr
	curr.Next = nil
	for {
		if pHead1 == nil {
			curr.Next = pHead2
			return head
		}
		if pHead2 == nil {
			curr.Next = pHead1
			return head
		}
		if pHead1.Val < pHead2.Val {
			curr.Next = pHead1
			pHead1 = pHead1.Next
			curr.Next = nil
		} else {
			curr.Next = pHead2
			pHead2 = pHead2.Next
			curr.Next = nil
		}
	}
}

// jz52
func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	switchCnt := 0
	origin1 := pHead1
	origin2 := pHead2
	for {
		if switchCnt > 2 {
			return nil
		}
		if pHead1 == nil || pHead2 == nil {
			return nil
		}
		if pHead1 == pHead2 {
			return pHead1
		}
		pHead1 = pHead1.Next
		pHead2 = pHead2.Next
		if pHead1 == nil {
			pHead1 = origin2
			switchCnt += 1
		}
		if pHead2 == nil {
			pHead2 = origin1
		}
	}

}

// jz23
func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	if pHead == nil {
		return nil
	}
	slow := pHead
	fast := pHead
	for {
		fast = fast.Next
		if fast == nil {
			return nil
		}
		fast = fast.Next
		if fast == nil {
			return nil
		}
		slow = slow.Next
		if fast == nil {
			return nil
		}
		if fast == slow {
			break
		}
	}
	//fast is currently at the meeting point
	slow = pHead
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

// jz22
func FindKthToTail(pHead *ListNode, k int) *ListNode {
	// write code here
	if pHead == nil || k == 0 {
		return nil
	}
	cnt := 1
	rHead := pHead
	next := pHead.Next
	for next != nil {
		tmp := next.Next
		next.Next = rHead
		rHead = next
		next = tmp
		cnt += 1
	}
	if cnt < k {
		return nil
	}
	next = rHead.Next
	rHead.Next = nil
	for i := 1; i < k; i += 1 {
		tmp := next.Next
		next.Next = rHead
		rHead = next
		next = tmp
	}
	return rHead
}

// jz76
func deleteDuplication(pHead *ListNode) *ListNode {
	// write code here
	if pHead == nil {
		return nil
	}
	dup := false
	next := pHead.Next
	head := pHead
	for next != nil && next.Val == head.Val {
		dup = true
		next = next.Next

	}
	head.Next = deleteDuplication(next)
	if dup {
		return head.Next
	}
	return head
}

// jz18
func deleteNode(head *ListNode, val int) *ListNode {
	// write code here
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	head.Next = deleteNode(head.Next, val)
	return head
}

// jz55
func TreeDepth(pRoot *TreeNode) int {
	// write code here
	if pRoot == nil {
		return 0
	}
	return 1 + max(TreeDepth(pRoot.Left), TreeDepth(pRoot.Right))
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}

// jz77
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param pRoot TreeNode类
 * @return int整型二维数组
 */
func Print(pRoot *TreeNode) [][]int {
	// write code here
	var curStack []*TreeNode
	nextStack := []*TreeNode{}
	res := [][]int{{}}
	if pRoot == nil {
		return [][]int{}
	}
	fromLeft := false
	curStack = append(curStack, pRoot)
	row := 0

	for {
		if len(curStack) == 0 && len(nextStack) == 0 {
			return res
		}
		if len(curStack) == 0 {
			fromLeft = !fromLeft
			curStack = nextStack
			nextStack = []*TreeNode{}
			row += 1
			res = append(res, []int{})
		}
		var cur *TreeNode
		if !fromLeft {
			cur = curStack[0]
			curStack = curStack[1:]
		} else {
			cur = curStack[len(curStack)-1]
			curStack = curStack[:len(curStack)-1]
		}
		if !fromLeft {
			if cur.Left != nil {
				nextStack = append(nextStack, cur.Left)
			}
			if cur.Right != nil {
				nextStack = append(nextStack, cur.Right)
			}
		} else {
			if cur.Right != nil {
				nextStack = append([]*TreeNode{cur.Right}, nextStack...)
			}
			if cur.Left != nil {
				nextStack = append([]*TreeNode{cur.Left}, nextStack...)
			}
		}
		res[row] = append(res[row], cur.Val)
	}
}

// jz54
// KthNode 查找第k小的节点的值
func KthNode(proot *TreeNode, k int) int {
	// write code here
	if k == 0 {
		return -1
	}
	arr := dfs(proot, []int{})
	if len(arr) < k-1 {
		return -1
	}
	return arr[k-1]
}

func dfs(cur *TreeNode, arr []int) []int {
	if cur == nil {
		return arr
	}
	arr = dfs(cur.Left, arr)
	arr = append(arr, cur.Val)
	return dfs(cur.Right, arr)
}

// jz7
func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	// write code here
	if len(pre) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: pre[0],
	}
	//     if len(pre) == 1 {
	//         return root
	//     }
	i := 0
	for {
		if vin[i] == root.Val {
			break
		}
		i += 1
	}
	root.Left = reConstructBinaryTree(pre[1:i+1], vin[:i])
	root.Right = reConstructBinaryTree(pre[i+1:], vin[i+1:])
	return root
}

// jz26
func HasSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	// write code here
	if pRoot1 == nil && pRoot2 == nil {
		return false
	}
	if pRoot1 == nil || pRoot2 == nil {
		return false
	}
	if isSameTree(pRoot1, pRoot2) {
		return true
	}
	return HasSubtree(pRoot1.Left, pRoot2) || HasSubtree(pRoot1.Right, pRoot2)
}

func isSameTree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	if pRoot1 == nil && pRoot2 == nil {
		return true
	}
	if pRoot2 == nil {
		return true
	}
	if pRoot1 == nil {
		return false
	}
	return pRoot1.Val == pRoot2.Val && isSameTree(pRoot1.Left, pRoot2.Left) && isSameTree(pRoot1.Right, pRoot2.Right)
}

// jz27
func Mirror(pRoot *TreeNode) *TreeNode {
	// write code here
	if pRoot == nil {
		return nil
	}
	left := pRoot.Left
	pRoot.Left = Mirror(pRoot.Right)
	pRoot.Right = Mirror(left)
	return pRoot
}

// jz32
func PrintFromTopToBottom(root *TreeNode) []int {
	// write code here
	return bfs(root)
}

func bfs(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	fifo := []*TreeNode{root}
	var cur *TreeNode
	for len(fifo) > 0 {
		cur = fifo[0]
		fifo = fifo[1:]
		res = append(res, cur.Val)
		if cur.Left != nil {
			fifo = append(fifo, cur.Left)
		}
		if cur.Right != nil {
			fifo = append(fifo, cur.Right)
		}
	}
	return res
}

/*
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。如果是则返回 true ,否则返回 false 。假设输入的数组的任意两个数字都互不相同。

数据范围： 节点数量 0 \le n \le 10000≤n≤1000 ，节点上的值满足 1 \le val \le 10^{5}1≤val≤10
5
  ，保证节点上的值各不相同
要求：空间复杂度 O(n)O(n) ，时间时间复杂度 O(n^2)O(n
2
 )
提示：
1.二叉搜索树是指父亲节点大于左子树中的全部节点，但是小于右子树中的全部节点的树。
2.该题我们约定空树不是二叉搜索树
3.后序遍历是指按照 “左子树-右子树-根节点” 的顺序遍历
4.参考下面的二叉搜索树，示例 1
*/

// jz33
func VerifySquenceOfBST(sequence []int) bool {
	// write code here
	if len(sequence) == 0 {
		return false
	}
	if len(sequence) == 1 {
		return true
	}
	root := sequence[len(sequence)-1]
	i := len(sequence) - 2
	for ; i > 0; i -= 1 {
		if sequence[i] < root {
			break
		}
	}

	if i <= 1 {
		return VerifySquenceOfBST(sequence[i : len(sequence)-1])
	}
	for j := 0; j < i; j += 1 {
		if sequence[j] > root {
			return false
		}
	}
	return VerifySquenceOfBST(sequence[:i]) && VerifySquenceOfBST(sequence[i:len(sequence)-1])
}

// jz82
func hasPathSum(root *TreeNode, sum int) bool {
	// write code here
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

// jz34
//func FindPath( root *TreeNode ,  expectNumber int ) [][]int {
//	// write code here
//	var res [][]int
//	if root == nil {
//		return res
//	}
//	if root.Left == nil && root.Right == nil && root.Val == expectNumber {
//		return [][]int{{root.Val}}
//	}
//	left := FindPath(root.Left, expectNumber - root.Val)
//	right := FindPath(root.Right, expectNumber - root.Val)
//
//	for i, _ := range left {
//		res = append(res, append([]int{root.Val}, left[i]...))
//	}
//	for i, _ := range right {
//		res = append(res, append([]int{root.Val}, right[i]...))
//	}
//	return res
//}

// jz36
func Convert(pRootOfTree *TreeNode) *TreeNode {
	// write code here
	if pRootOfTree == nil {
		return nil
	}
	h, _ := preOrder(pRootOfTree)
	return h
}

func preOrder(node *TreeNode) (head *TreeNode, tail *TreeNode) {
	if node.Left == nil && node.Right == nil {
		return node, node
	}
	var rH, rT *TreeNode
	if node.Right != nil {
		rH, rT = preOrder(node.Right)
		node.Right = rH
		if rH != nil {
			rH.Left = node
		}
	} else {
		rT = node
	}
	var lH, lT *TreeNode
	if node.Left != nil {
		lH, lT = preOrder(node.Left)
		lT.Right = node
		node.Left = lT
	} else {
		lH = node
	}
	return lH, rT
}

// jz79
func IsBalanced_Solution(pRoot *TreeNode) bool {
	// write code here
	if pRoot == nil {
		return true
	}
	_, b := depthAndBalanced(pRoot)
	return b
}

func depthAndBalanced(node *TreeNode) (depth int, balanced bool) {
	if node.Left == nil && node.Right == nil {
		return 1, true
	}
	var lD int
	var lB bool
	var rD int
	var rB bool
	if node.Left != nil {
		lD, lB = depthAndBalanced(node.Left)
	} else {
		lD = 0
		lB = true
	}
	if node.Right != nil {
		rD, rB = depthAndBalanced(node.Right)
	} else {
		rD = 0
		rB = true
	}
	var maxDep int
	if rD > lD {
		maxDep = rD
	} else {
		maxDep = lD
	}
	return maxDep + 1, abs(lD-rD) <= 1 && lB && rB
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

type TreeLinkNode struct {
	Val   int
	Left  *TreeLinkNode
	Right *TreeLinkNode
	Next  *TreeLinkNode
}

// jz8
func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
	if pNode.Right != nil {
		cur := pNode.Right
		for {
			if cur.Left == nil {
				return cur
			}
			cur = cur.Left
		}
	}

	cur := pNode
	parent := pNode.Next
	for {
		if parent == nil {
			return nil
		}
		if parent.Left == cur {
			return parent
		}
		if parent.Right == cur {
			cur = parent
			parent = parent.Next
		}
	}
	return nil
}

// jz28
func isSymmetrical(pRoot *TreeNode) bool {
	// write code here
	if pRoot == nil {
		return true
	}
	return symmetricHelper(pRoot.Left, pRoot.Right)
}

func symmetricHelper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return symmetricHelper(left.Left, right.Right) && symmetricHelper(left.Right, right.Left)
}

// JZ78
func Print1(pRoot *TreeNode) [][]int {
	// write code here
	var res [][]int
	if pRoot == nil {
		return res
	}
	stack := []*TreeNode{pRoot}
	nextStack := make([]*TreeNode, 0, 0)
	row := 0
	for len(stack) > 0 {
		if len(res) <= row {
			res = append(res, []int{})
		}
		curNode := stack[0]
		stack = stack[1:]
		if curNode.Left != nil {
			nextStack = append(nextStack, curNode.Left)
		}
		if curNode.Right != nil {
			nextStack = append(nextStack, curNode.Right)
		}
		res[row] = append(res[row], curNode.Val)
		if len(stack) == 0 {
			row += 1
			stack = nextStack
			nextStack = make([]*TreeNode, 0, 0)
		}
	}
	return res
}

func Serialize(root *TreeNode) string {
	// write code here
	if root == nil {
		return "#"
	}
	builder := strings.Builder{}
	curStack := []*TreeNode{root}
	nextStack := make([]*TreeNode, 0)
	for len(curStack) > 0 {
		node := curStack[0]
		curStack = curStack[1:]
		if curStack == nil {
			builder.WriteString("#")
		} else {
			builder.WriteString(strconv.FormatInt(int64(node.Val), 10))
		}
		if len(curStack) > 0 || len(nextStack) > 0 {
			builder.WriteString(",")
		}
	}
	return builder.String()
}
func Deserialize(s string) *TreeNode {
	// write code here
	if len(s) == 0 || s == "#" {
		return nil
	}
	return nil
}

// jz84
func FindPath(root *TreeNode, sum int) int {
	// write code here
	if root == nil {
		return res
	}
	findPathDfs(root, sum)
	FindPath(root.Left, sum)
	FindPath(root.Right, sum)
	return res
}

var res = 0

// jz84
func findPathDfs(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	if root.Val == sum {
		res += 1
	}
	findPathDfs(root.Left, sum-root.Val)
	findPathDfs(root.Right, sum-root.Val)
}

// jz86
func lowestCommonAncestorBinaryTree(root *TreeNode, o1 int, o2 int) int {
	// write code here
	return lowestHelper(root, o1, o2).Val
}

func lowestHelper(root *TreeNode, o1, o2 int) *TreeNode {
	if root == nil || root.Val == o1 || root.Val == o2 {
		return root
	}
	left := lowestHelper(root.Left, o1, o2)
	right := lowestHelper(root.Right, o1, o2)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

// jz68
func lowestCommonAncestor(root *TreeNode, p int, q int) int {
	// write code here
	if p == root.Val || q == root.Val {
		return root.Val
	}
	if p < root.Val && q > root.Val {
		return root.Val
	}
	if p < root.Val && q < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return lowestCommonAncestor(root.Right, p, q)
}

// jz9
//var stack1 [] int
//var stack2 [] int
//
//var popCnt = 0
//var pushCnt = 0
//
//func Push(node int) {
//	if pushCnt % 2 == 0 {
//		stack1 = append(stack1, node)
//	} else {
//		stack2 = append(stack2, node)
//	}
//	pushCnt += 1
//}
//
//func Pop() int{
//	var res int
//	if popCnt % 2 == 0 {
//		res = stack1[0]
//		stack1 = stack1[1:]
//	} else {
//		res = stack2[0]
//		stack2 = stack2[1:]
//	}
//	popCnt += 1
//	return res
//}

// jz30
var stack []int

func Push(node int) {
	// write code here
	stack = append([]int{node}, stack...)
}
func Pop() {
	// write code here
	stack = stack[1:]
}
func Top() int {
	// write code here
	return stack[0]
}
func Min() int {
	// write code here
	min := stack[0]
	for i, _ := range stack {
		if stack[i] < min {
			min = stack[i]
		}
	}
	return min
}

// jz31
func IsPopOrder(pushV []int, popV []int) bool {
	// write code here
	if len(pushV) != len(popV) {
		return false
	}
	if len(pushV) == 0 && len(popV) == 0 {
		return true
	}
	if pushV[0] == popV[len(popV)-1] {
		return IsPopOrder(pushV[1:], popV[:len(popV)-1])
	} else {
		if pushV[0] == popV[0] {
			return IsPopOrder(pushV[1:], popV[1:])
		} else if pushV[len(pushV)-1] == popV[len(popV)-1] {
			return IsPopOrder(pushV[:len(pushV)-1], popV[:len(popV)-1])
		}
		start := pushV[0]
		for i, _ := range popV {
			if start == popV[i] {
				return IsPopOrder(pushV[0:i+1], popV[0:i+1]) && IsPopOrder(pushV[i+1:], popV[i+1:])
			}
		}
		return false
	}
}

// jz73
func ReverseSentence(str string) string {
	// write code here
	arr := strings.Split(str, " ")
	builder := strings.Builder{}
	for i := len(arr) - 1; i >= 0; i -= 1 {
		builder.WriteString(arr[i])
		if i > 0 {
			builder.WriteString(" ")
		}
	}
	return builder.String()
}

// jz59
func maxInWindows(num []int, size int) []int {
	// write code here
	var res []int
	i := 0
	j := size - 1
	for j < len(num) {
		max := num[i]
		for k := i; k <= j; k += 1 {
			if num[k] > max {
				max = num[k]
			}
		}
		res = append(res, max)
		j += 1
		i += 1
	}
	return res
}

// jz53
func GetNumberOfK(data []int, k int) int {
	// write code here
	pos := positionHelper(data, 0, len(data)-1, k)
	if pos < 0 {
		return 0
	}
	cnt := 0
	for i := pos; i < len(data); i += 1 {
		if data[i] == k {
			cnt += 1
		} else {
			break
		}
	}
	for i := pos - 1; i >= 0; i -= 1 {
		if data[i] == k {
			cnt += 1
		} else {
			break
		}
	}
	return cnt
}

func positionHelper(data []int, start, end, target int) int {
	if end < start {
		return -1
	}
	if start == end {
		if data[start] == target {
			return start
		}
		return -1
	}
	mid := (start + end) / 2
	if data[mid] == target {
		return mid
	} else if data[mid] < target {
		return positionHelper(data, mid+1, end, target)
	} else {
		return positionHelper(data, start, mid-1, target)
	}
}

// jz4
func Find(target int, array [][]int) bool {
	// write code here
	i := 0
	j := len(array[0]) - 1
	if j < 0 {
		return false
	}
	for i < len(array) && j >= 0 {
		if array[i][j] == target {
			return true
		}
		if target > array[i][j] {
			i += 1
		} else {
			j -= 1
		}
	}
	return false
}

// jz11
func minNumberInRotateArray(rotateArray []int) int {
	// write code here
	return minRotateHelper(rotateArray, 0, len(rotateArray)-1)
}

func minRotateHelper(rotate []int, start, end int) int {
	if start == end {
		return rotate[start]
	}
	if end-start == 1 {
		if rotate[start] < rotate[end] {
			return rotate[start]
		}
		return rotate[end]
	}
	mid := (start + end) / 2
	if rotate[mid] > rotate[end] {

		return minRotateHelper(rotate, mid, end)
	} else {
		for i := mid; i <= end; i += 1 {
			if rotate[i] < rotate[mid] {
				return minRotateHelper(rotate, mid, end)
			}
		}
		return minRotateHelper(rotate, start, mid)
	}
}

// jz38
func Permutation(str string) []string {
	// write code here
	if len(str) == 1 {
		return []string{str}
	}
	used := make(map[byte]bool, len(str))
	var res []string
	for i := 0; i < len(str); i += 1 {
		if _, ok := used[str[i]]; !ok {
			used[str[i]] = true
			tmp := Permutation(str[0:i] + str[i+1:])
			for j, _ := range tmp {
				res = append(res, string(str[i])+tmp[j])
			}
		}
	}
	return res
}

// jz44
// 012345678910111213141516
func findNthDigit(n int) int {
	//raw := n
	// write code here
	if n < 10 {
		return n
	}
	nLen := 1
	for true {
		nTotal := 9 * int(math.Pow10(nLen-1))

		if n < nTotal {
			break
		}
		n -= nTotal
		nLen += 1
	}
	a := int(math.Pow10(nLen-1)) + (n-1)/nLen
	fmt.Println(a)
	digit := (n - 1) % nLen
	str := strconv.Itoa(a)
	res, _ := strconv.Atoi(string(str[digit]))
	return res
}

func lenNum(n int) int {
	cnt := 0
	for n > 0 {
		n = n % 10
		cnt += 1
	}
	return cnt
}

// jz44
func findNthDigit2(n int) int {
	digit := 1
	start := 1
	sum := 9
	for n > sum {
		n -= sum
		start *= 10
		digit += 1
		sum = 9 * start * digit
	}
	str := strconv.Itoa(start + (n-1)/digit)
	index := (n - 1) % digit
	d, _ := strconv.Atoi(string(str[index]))
	return d
}

// jz42
func FindGreatestSumOfSubArray(array []int) int {
	// write code here
	var res int
	dp := make([]int, len(array))
	dp[0] = array[0]
	res = dp[0]
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for i := 1; i < len(array); i += 1 {
		dp[i] = max(dp[i-1]+array[i], array[i])
		res = max(res, dp[i])
	}
	return res
}

// jz42
func FindGreatestSumOfSubArray_sol2(array []int) int {
	// write code here
	var res int
	cur := array[0]
	res = array[0]

	for i := 1; i < len(array); i += 1 {
		cur = max(cur+array[i], array[i])
		res = max(res, cur)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// jz85
func FindGreatestSumOfSubArray2(array []int) []int {
	// write code here

	dp := make([]int, len(array))
	dp[0] = array[0]
	res := dp[0]
	left, right := 0, 0
	resL, resR := 0, 0
	for i := 1; i < len(array); i += 1 {
		right += 1
		dp[i] = max(array[i]+dp[i-1], array[i])
		if array[i] > array[i]+dp[i-1] {
			left = right
		}
		if dp[i] > res || dp[i] == res && (right-left+1) > (resR-resL+1) {
			res = dp[i]
			resL = left
			resR = right
		}
	}

	return array[resL : resR+1]
}

// jz69
func jumpFloor(number int) int {
	// write code here
	if number <= 2 {
		return number
	}
	a := 1
	b := 2
	for i := 3; i <= number; i += 1 {
		tmp := a + b
		a = b
		b = tmp
	}
	return b
}

// jz10
func Fibonacci(n int) int {
	// write code here
	var first = 1
	var second = 1
	if n == 1 {
		return first
	} else if n == 2 {
		return second
	}
	for i := 2; i < n; i += 1 {
		j := first
		first = second
		second = j + first
	}
	return second
}

// jz19
func match(str string, pattern string) bool {
	// write code here
	if len(str) == 0 {
		return true
	}

	newPattern, target, req := popPattern(pattern)
	res := false
	if req == "" {
		return string(str[0]) == target && match(str[1:], newPattern)
	}
	if req == "#" {
		return string(str[0]) == target && (match(str[1:], pattern))
	}
	return res
}

func popPattern(pattern string) (p, target, requirement string) {
	target = string(pattern[0])
	pattern = pattern[1:]
	if len(pattern) > 0 {
		requirement = string(pattern[0])
		if requirement == "." || requirement == "*" {
			pattern = pattern[1:]
		} else {
			requirement = ""
		}
	}
	return pattern, target, requirement
}

// jz71
// 一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。求该青蛙跳上一个n级的台阶(n为正整数)总共有多少种跳法。
func jumpFloorII(number int) int {
	// write code here
	if number <= 2 {
		return number
	}
	res := 2
	for i := 2; i < number; i += 1 {
		res *= 2
	}
	return res
}

// jz70
func rectCover(number int) int {
	// write code here
	if number <= 3 {
		return number
	}
	a := 2
	b := 3
	for i := 4; i <= number; i += 1 {
		tmp := a + b
		a = b
		b = tmp
	}
	return b
}

// jz63
func maxProfit(prices []int) int {
	// write code here
	profit := 0
	min := prices[0]
	for i := 1; i < len(prices); i += 1 {
		newProfit := prices[i] - min
		if newProfit > profit {
			profit = newProfit
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return profit
}

// jz47
func maxValue(grid [][]int) int {
	// write code here
	for i := 1; i < len(grid[0]); i += 1 {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < len(grid); i += 1 {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < len(grid[0]); i += 1 {
		for j := 1; j < len(grid); j += 1 {
			var tmp int
			if grid[j][i-1] > grid[j-1][i] {
				tmp = grid[j][i-1]
			} else {
				tmp = grid[j-1][i]
			}
			grid[j][i] += tmp
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

// jz48
func lengthOfLongestSubstring(s string) int {
	// write code here
	l, r := 0, 0
	cnt := 0
	m := map[byte]int{}
	for r < len(s) {
		m[s[r]] = m[s[r]] + 1
		if m[s[r]] != 1 {
			for m[s[r]] > 1 {
				m[s[l]] = m[s[l]] - 1
				if m[s[l]] == 0 {
					delete(m, s[l])
				}
				l += 1
			}
		}
		if len(m) > cnt {
			cnt = len(m)
		}
		r += 1
	}
	return cnt
}

// jz46
func solve(nums string) int {
	// write code here
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	if nums[1] == '0' {
		dp[1] = 1
	} else {
		if (nums[0] <= '2' && nums[1] <= '6') || (nums[0] < '2') {
			dp[1] = 2
		} else {
			dp[1] = 1
		}
	}
	for i := 2; i < len(nums); i += 1 {
		if nums[i] == '0' {
			if nums[i-1] <= '2' {
				dp[i] = dp[i-2]
			} else {
				dp[i] = 0
			}
		} else {
			if nums[i-1] == '0' {
				dp[i] = dp[i-2]
			} else if (nums[i-1] <= '2' && nums[i] <= '6') || (nums[i-1] == '1') {
				dp[i] = dp[i-2] + dp[i-1]
			} else {
				dp[i] = dp[i-1]
			}
		}
	}
	return dp[len(nums)-1]
}

// jz12
func hasPath(matrix [][]byte, word string) bool {
	// write code here
	find := false
	for i := 0; i < len(matrix); i += 1 {
		for j := 0; j < len(matrix[0]); j += 1 {
			if matrix[i][j] == word[0] {
				matrix[i][j] = '1'
				find = find || hasPathHelper(matrix, word[1:], i, j)
				matrix[i][j] = word[0]
			}
		}
	}
	return find
}

func hasPathHelper(matrix [][]byte, word string, curRow, curCol int) bool {
	if len(word) == 0 {
		return true
	}
	find := false
	var i, j int
	i = curRow - 1
	j = curCol
	if !(i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0])) {
		if matrix[i][j] == word[0] {
			matrix[i][j] = '1'
			find = find || hasPathHelper(matrix, word[1:], i, j)
			matrix[i][j] = word[0]
		}
	}
	i = curRow + 1
	j = curCol
	if !(i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0])) {
		if matrix[i][j] == word[0] {
			matrix[i][j] = '1'
			find = find || hasPathHelper(matrix, word[1:], i, j)
			matrix[i][j] = word[0]
		}
	}
	i = curRow
	j = curCol - 1
	if !(i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0])) {
		if matrix[i][j] == word[0] {
			matrix[i][j] = '1'
			find = find || hasPathHelper(matrix, word[1:], i, j)
			matrix[i][j] = word[0]
		}
	}
	i = curRow
	j = curCol + 1
	if !(i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0])) {
		if matrix[i][j] == word[0] {
			matrix[i][j] = '1'
			find = find || hasPathHelper(matrix, word[1:], i, j)
			matrix[i][j] = word[0]
		}
	}
	return find
}

// jz13
func movingCount(threshold int, rows int, cols int) int {
	// write code here
	res = 0
	matrix := make([][]bool, rows)
	for i := 0; i < rows; i += 1 {
		matrix[i] = make([]bool, cols)
	}
	helper(threshold, rows, cols, 0, 0, matrix)
	return res
}

func cnt(in int) int {
	sum := 0
	for in > 0 {
		sum += in % 10
		in = in / 10
	}
	return sum
}

func helper(threshold int, rows int, cols int, curRow, curCol int, matrix [][]bool) {
	if curRow >= rows || curCol >= cols || matrix[curRow][curCol] {
		return
	}
	if cnt(curRow)+cnt(curCol) > threshold {
		return
	}
	res += 1
	matrix[curRow][curCol] = true
	helper(threshold, rows, cols, curRow+1, curCol, matrix)
	helper(threshold, rows, cols, curRow, curCol+1, matrix)
}

// jz3
func duplicate(numbers []int) int {
	// write code here
	m := map[int]bool{}
	for i, _ := range numbers {
		if m[numbers[i]] {
			return numbers[i]
		}
		m[numbers[i]] = true
	}
	return -1
}

// jz51
func InversePairs(data []int) int {
	// write code here
	cnt := 0
	mergeSort(data, 0, len(data)-1, &cnt)
	return cnt
}

func mergeSort(data []int, start, end int, count *int) {
	if start == end {
		return
	}
	mid := (start + end) / 2
	mergeSort(data, start, mid, count)
	mergeSort(data, mid+1, end, count)
	for i := mid; i > 0; i -= 1 {
		//if data[i] >
	}
}

// jz40
func GetLeastNumbers_Solution(input []int, k int) []int {
	// write code here
	GetLeastNumbers_Solution_helper(input, 0, len(input)-1)
	return input[:k]
}

func GetLeastNumbers_Solution_helper(input []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	GetLeastNumbers_Solution_helper(input, start, mid)
	GetLeastNumbers_Solution_helper(input, mid+1, end)
	tmp := make([]int, end-start+1)
	copy(tmp, input[start:end+1])
	i := start
	j := mid + 1
	index := 0
	for i <= mid && j <= end {
		if input[i] > input[j] {
			tmp[index] = input[j]
			j += 1
		} else {
			tmp[index] = input[i]
			i += 1
		}
		index += 1
	}
	if i <= mid {
		copy(tmp[index:], input[i:mid+1])
	} else {
		copy(tmp[index:], input[j:end+1])
	}
	copy(input[start:end+1], tmp)
}

// jz41
//func Insert(num int) {
//	index := search(num)
//	rawLen := len(arr)
//	arr = append(arr, 1)
//	fmt.Println(arr[index : rawLen+1])
//	copy(arr[index+1:], arr[index:rawLen+1])
//	arr[index] = num
//}
//
//func search(target int) int {
//	if len(arr) == 0 {
//		return 0
//	}
//	start, mid := 0, 0
//	end := len(arr) - 1
//
//	for start < end {
//		mid = (start + end) / 2
//		if mid == start {
//			if arr[mid] < target {
//				return mid + 1
//			} else {
//				return mid
//			}
//		}
//		if arr[mid] == target {
//			return mid
//		}
//		if arr[mid] < target {
//			start = mid
//		} else {
//			end = mid
//		}
//	}
//	if arr[mid] < target {
//		return mid + 1
//	} else {
//		return mid
//	}
//}
//
//var arr = []int{}
//
//func GetMedian() float64 {
//	if len(arr)%2 == 1 {
//		return float64(arr[len(arr)/2])
//	}
//	return (float64(arr[len(arr)/2]) + float64(arr[len(arr)/2-1])) / 2
//}

func Insert(num int) {
	arr = append(arr, num)
}

var arr = []int{}

func GetMedian() float64 {
	mSort(arr, 0, len(arr)-1)
	if len(arr)%2 == 1 {
		return float64(arr[len(arr)/2])
	}
	return (float64(arr[len(arr)/2]) + float64(arr[len(arr)/2-1])) / 2
}

func mSort(input []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	GetLeastNumbers_Solution_helper(input, start, mid)
	GetLeastNumbers_Solution_helper(input, mid+1, end)
	tmp := make([]int, end-start+1)
	copy(tmp, input[start:end+1])
	i := start
	j := mid + 1
	index := 0
	for i <= mid && j <= end {
		if input[i] > input[j] {
			tmp[index] = input[j]
			j += 1
		} else {
			tmp[index] = input[i]
			i += 1
		}
		index += 1
	}
	if i <= mid {
		copy(tmp[index:], input[i:mid+1])
	} else {
		copy(tmp[index:], input[j:end+1])
	}
	copy(input[start:end+1], tmp)
}

// jz65
func Add(num1 int, num2 int) int {
	// write code here
	sum := num1
	add := num2
	for add != 0 {
		tmp := sum ^ add
		add = (sum & add) << 1
		sum = tmp
	}
	return sum
}

func Add2(num1 int, num2 int) int {
	// write code here
	sum := num1
	add := num2
	for add != 0 {
		tmp := sum ^ add
		add = (sum & add) << 1
		sum = tmp
	}
	return sum
}

func NumberOf1(n int) int {
	// write code here
	cnt := 0
	if n < 0 {
		fmt.Printf("%x\n", n)
		n = n & 0xffffffff
		fmt.Printf("%b\n", n)
	}
	for n != 0 {
		cnt += n & 1
		n = n >> 1
	}
	return cnt
}
