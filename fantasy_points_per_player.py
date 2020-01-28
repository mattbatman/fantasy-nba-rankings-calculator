def create_fantasy_points_per_player(stat_data, points_due_data, categories):
    """Cycles through stat data to pull the fantasy points due for each player.
    """
    fantasy_points_per_player = []
    for player in stat_data:
        fantasy_points_for_this_player = {}
        for category in categories:
            player_stat_value = str(player[category])
            points_due = points_due_data[category][player_stat_value]['due']
            fantasy_points_for_this_player[category] = points_due
        fantasy_points_for_this_player['TOTAL'] = sum(
            fantasy_points_for_this_player.values())
        fantasy_points_for_this_player['PLAYER'] = player['PLAYER']
        fantasy_points_per_player.append(fantasy_points_for_this_player)
    return fantasy_points_per_player
