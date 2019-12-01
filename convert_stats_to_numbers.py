def convert_stats_to_numbers(data, stat_categories):
    """Convert data as strings to floats.
    Loop over the data (expected to be a list of dicts), then
    loop over the stat_categories (expected to be a list of strings).
    For the datum at that category as a key, convert the value to a float.
    """
    for datum in data:
        for category in stat_categories:
            datum[category] = float(datum[category])

    return data
