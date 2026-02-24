# Function to create dictionary
# arr is list of tuple. tuple contain name and marks.


def create_dict(arr):

    dict = {}

    for name, mark in arr:
        dict[name] = mark
        
    return dict