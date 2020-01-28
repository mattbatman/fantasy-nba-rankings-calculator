"""
create_points_data and create_points_for_category are meant to be used
together. create_points_data will cycle through data, sort it by category, and
call create_points_for_category for the sorted data at each category.
"""


def create_points_data(stat_data, categories):
    """
    Sorts data for each category and determines: where that value ranks after
    sorting, how many players had that value for that category, and the fantasy
    points due to a player that had that value at that category.
    Returns a dict of dicts that looks like:
    {
        'FG%': {
            '44.9': {'count': 1, 'due': 1, 'ranks_inverse': [1]},
            '53.4': {'count': 1, 'due': 2, 'ranks_inverse': [2]},
            '54.2': {'count': 1, 'due': 3, 'ranks_inverse': [3]},
        },
        'FT%': {
            '73.1': {'count': 1, 'due': 1, 'ranks_inverse': [1]},
            '82.8': {'count': 1, 'due': 2, 'ranks_inverse': [2]},
            '85.8': {'count': 1, 'due': 3, 'ranks_inverse': [3]},
        },
        ...
    }
    """
    ranks = {}

    for category in categories:
        is_not_tov = category != 'TOV'
        sorted_data = sorted(stat_data, key=lambda k: k[category]) if is_not_tov else sorted(
            stat_data, key=lambda k: k[category], reverse=True)
        ranks[category] = create_points_for_category(sorted_data, category)

    return ranks


def create_points_for_category(data_sorted, category):
    """ Takes a list of sorted player data and the category that points are being
    assigned for. Returns a dict that looks like:
    {
        '1000': {
            'ranks_inverse': [539, 540],
            'count': 2,
            'due': 539.5
        },
        { ... },
        ...
    }
    """
    points_due_for_category = {}
    for index, player in enumerate(data_sorted):
        value_for_player = player[category]
        key_in_data = str(value_for_player)
        inverse_rank_for_player = index + 1
        if key_in_data in points_due_for_category:
            points_due_for_category[key_in_data]['ranks_inverse'].append(
                inverse_rank_for_player)
            points_due_for_category[key_in_data]['count'] = len(
                points_due_for_category[key_in_data]['ranks_inverse'])
            points_due_for_category[key_in_data]['due'] = sum(
                points_due_for_category[key_in_data]['ranks_inverse']) / points_due_for_category[key_in_data]['count']
        else:
            points_due_for_category[key_in_data] = {
                'count': 1,
                'due': inverse_rank_for_player,
                'ranks_inverse': [inverse_rank_for_player]
            }

    return points_due_for_category
