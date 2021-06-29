class LRUCache:

        
# 利用字典和双链表实现
    def __init__(self, capacity: int):
        self.cap=capacity
        self.dic=dict()
        # 哨兵节点
        self.dummyHead=Node(0,0,None,None)
        self.dummyTail=Node(0,0,None,self.dummyHead)
        self.dummyHead.next=self.dummyTail

    def get(self, key: int) -> int:
        if key in self.dic:
            theNode = self.dic[key]
            value = theNode.val
            # 查询时候需要更新LRU，把查询的节点放入链表头部
            self.dic[key] = self.put(key,value)
            
            return value
        else:
            return -1

    def put(self, key: int, value: int) -> None:
        if key in self.dic: #把查询的节点放入头部
            self.remove(key)
            
        elif len(self.dic)==self.cap: # LRU满了，淘汰最后一个节点
            self.remove(self.dummyTail.prev.key)
        return self.add(key,value)

    def add(self,key,value):
        old = self.dummyHead.next
        new = Node(key,value,old,self.dummyHead)
        self.dummyHead.next= new
        old.prev=new
        self.dic[key]=new
        
        return new

    def remove(self,key):
        if key not in self.dic:
            return
        node = self.dic[key]
        node.prev.next,node.next.prev = node.next,node.prev #从链表删除节点
        self.dic.pop(key)

# 辅助类，保存字典键、值，前一个链表节点、后一个链表节点
            
class Node:
    def __init__(self,key, val,next,prev):
        self.key=key
        self.val=val
        self.next=next
        self.prev=prev



# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)