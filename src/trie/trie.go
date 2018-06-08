// trie 字典树实现
package trie

//字典树结构
type Trie struct{
	root *Node
}
//树形结构节点
type Node struct{
	chirldren map[rune] *Node
	isEnd bool
}
//新建节点
func newNode() *Node {
	n := new(Node)
	n.chirldren = make(map[rune] *Node)
	n.isEnd = false

	return n
}
//创建字典树
func NewTrie() *Trie{
	t := new(Trie)
	t.root = newNode()
	return t
}

//字典树添加敏感词
func (t *Trie) Add(word string){
	chars := []rune(word)

	if(len(chars) <= 0){
		return
	}

	node := t.root

	for _,char := range chars{

		if _,ret := node.chirldren[char]; !ret{
			node.chirldren[char] = newNode()
		}

		node = node.chirldren[char]

	}

	node.isEnd = true

}

//在字典树中查找完全匹配的关键词
func (t *Trie) Search(word string) bool {
	chars := []rune(word)

	result := false

	if(len(chars) <= 0){
		return result
	}

	node := t.root

	for _,char := range chars{
		if _,ret := node.chirldren[char]; ret{

			if node.chirldren[char].isEnd {
				result = true
			}

			node = node.chirldren[char]
		}else{
			return false
		}
	}

	return result
}

func (t *Trie) ReplaceWord(word string) string {
	chars := []rune(word)

	var box []rune

	var startKey int = 0

	var startFlag bool = true

	if(len(chars) <= 0){
		return ""
	}

	node := t.root

	for key,char := range chars{
		if _,ret := node.chirldren[char]; ret{
			box = append(box,rune(42))
			if startFlag {
				startKey = key
				startFlag = false
			}
			if node.chirldren[char].isEnd {
				for bKey,bRlt := range box{

					chars[startKey+bKey] = bRlt

				}
			}
			node = node.chirldren[char]
		}else{
			startFlag 	= true
			box 		= box[0:0]
			node 		= t.root
		}
	}

	return string(chars)
}