def find(search_list, value):
    lo, hi = 0, len(search_list)-1
    while lo <= hi:
        print(lo, hi)
        mid = lo + (hi-lo)//2
        if search_list[mid] == value:
            return mid
        if search_list[mid] < value:
            lo = mid+1
        else:
            hi = mid-1
    raise ValueError("value not in array")
            
