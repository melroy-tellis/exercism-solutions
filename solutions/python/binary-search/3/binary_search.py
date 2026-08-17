"""
binary_search.py: Implementation of binary search.
"""
def find(search_list, value):
    """
    Find the index of a value in a list.

    Arguments:
      search_list: The list to search.
      value: The value to find.
    Returns:
      The index of the value in the list.
    Raises:
      ValueError if the value is not found in the list.
    """
    low, high = 0, len(search_list)-1
    while low <= high:
        mid = low + (high-low)//2
        if search_list[mid] == value:
            return mid
        if search_list[mid] < value:
            low = mid+1
        else:
            high = mid-1
    raise ValueError("value not in array")
            
