#!/usr/bin/env python
# -*- coding: utf-8 -*-


class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def getMinimumDifference(self, root):
        return self.doDifference(root, None)
    
    def doDifference(self, root, prev):
        if not root:
            return prev
        
        largest = self.getLarget(root.left)
        if largest:
            difference = abs(largest.val - root.val)
            prev = prev if prev and prev < difference else difference
        smallest = self.getSmallest(root.right)
        if smallest:
            difference = abs(smallest.val - root.val)
            prev = prev if prev and prev < difference else difference
        
        left = self.doDifference(root.left, prev)
        right = self.doDifference(root.right, prev)

        if left:
            prev = prev if prev and prev < left else left
        if right:
            prev = prev if prev and prev < right else right
        return prev


    def getLarget(self, root):
        if not root:
            return None
        while root.right:
            root = root.right
        return root

    def getSmallest(self, root):
        if not root:
            return None
        while root.left:
            root = root.left
        return root

if __name__ == '__main__':
    root = TreeNode(1)
    root.right = TreeNode(3)
    root.right.left = TreeNode(2)
    assert Solution().getMinimumDifference(root) == 1