#!/usr/bin/env python
# -*- coding: utf-8 -*-

class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def convertBST(self, root):
        return self.greater(root, 0)

    def greater(self, root, prev):
        if not root:
            return None
        self.greater(root.right, prev)
        right_lesat = self.findRightLeast(root.right)
        if right_lesat:
            root.val = root.val + right_lesat.val
        else:
            root.val = root.val + prev
        self.greater(root.left, root.val)
        return root
    
    def findRightLeast(self, root):
        re = root
        while re:
            if re.left:
                re = re.left
            else:
                return re
        return None

def preOrder(root):
    if root:
        print root.val
        preOrder(root.left)
        preOrder(root.right)

if __name__ == '__main__':
    root = TreeNode(2)
    root.left = TreeNode(0)
    root.right = TreeNode(3)
    root.left.left = TreeNode(-4)
    root.left.right = TreeNode(1)
    preOrder(root)
    print

    root = Solution().convertBST(root)
    preOrder(root)
