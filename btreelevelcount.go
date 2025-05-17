package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHight := BTreeLevelCount(root.Left)
	rightHight := BTreeLevelCount(root.Right)

	if leftHight > rightHight {
		return leftHight + 1
	}
	return rightHight + 1
}
