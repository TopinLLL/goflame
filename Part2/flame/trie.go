package flame

import "strings"

//前缀树的节点
type node struct {
	pattern string //待匹配路由，只有叶子节点才可能被赋值为非空
	part string //待匹配路由的一部分，当前要处理的部分 例如:lang
	children []*node //子节点
	isWild bool //是否是精准匹配，含有：或*时为true
}

//查找第一个匹配的节点，用于插入
func(n *node)matchChild(part string)*node{
	for _,child:=range n.children{
		if child.part==part||child.isWild{
			return child
		}
	}
	return nil
}

//查找所有匹配的节点，用于匹配
func(n *node)matchChildren(part string)[]*node{
	nodes:=make([]*node,0)
	for _,child:=range n.children{
		if child.part==part||child.isWild{
			nodes=append(nodes,child)
		}
	}
	return nodes
}

//插入part
func(n *node)insert(pattern string,parts []string,height int){
	//递归退出条件是匹配至parts的最后一段
	if len(parts)==height{
		//只有叶子节点pattern不置为空
		n.pattern=pattern
		return
	}
	//当前part
	part:=parts[height]
	//查找第一个匹配的孩子，直到招不到为止
	child:=n.matchChild(part)
	//构造child节点并且插入到前缀树中
	if child == nil{
		child = &node{
			part:     part,
			isWild:   part[0]==':'||part[0]=='*',
		}
		n.children=append(n.children,child)
	}
	child.insert(pattern,parts,height+1)
}

//找到匹配节点
func (n *node)search(parts []string,height int)*node{
	//递归退出条件是找到parts的最后一部分，或者路径中出现‘*’
	if len(parts)==height||strings.HasPrefix(n.part,"*"){
		if n.pattern ==""{
			return nil
		}
		return n
	}

	part:=parts[height]
	children:=n.matchChildren(part)

	for _,child:=range children{
		result:=child.search(parts,height+1)
		if result!=nil{
			return result
		}
	}
	return nil
}



