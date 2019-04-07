#!/usr/bin/env python3
# -*- coding: utf-8 -*-

class Solution:
    def twoSum(self, nums, target):
        numMap = {}
        for i in range(len(nums)):
            if (target - nums[i]) in numMap:
                return [i, numMap[target - nums[i]]]
            else:
                numMap[nums[i]] = i

if __name__ == '__main__':
    assert Solution().twoSum([3, 2, 4], 6) == [1, 2]
