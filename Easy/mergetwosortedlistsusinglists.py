# Merging two Lists and returning a Sorted list using Lists

class Solution:
  def mergeTwoLists(self, l1: List, l2: List) -> List:
    for hold, val in enumerate(l2):
      l1.append(val)
    l1.sort()
    print(l1)
