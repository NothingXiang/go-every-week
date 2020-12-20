package gee

import "strings"

type node struct {
	// 待匹配路由，例如 /p/:lang
	// 只有叶子节点该字段才不为空
	pattern string

	// 路由中的一部分，例如 :lang
	part string

	// 子节点，例如 [doc, tutorial, intro]
	children []*node

	// 是否精确匹配，part 含有 : 或 * 时为true
	isWild bool
}

// 返回第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 返回所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 重点方法1 insert： 按照指定规则插入路径。递归方法
// param: pattern 完整url
// param: parts	按照'/'拆分的url
// param: height 高度，用于标识当前对应哪个parts
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)

	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// 重点方法2： search,用于匹配路由
// param: parts 按照'/'拆分的url
// param: height 高度，用于标识当前对应哪个parts
// return: 按照前序遍历方式，最早匹配成功的叶子节点
func (n *node) search(parts []string, height int) *node {
	// 遍历的停止条件 保证匹配到叶子节点
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		// 这个是啥情况？
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	// 前序遍历
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
