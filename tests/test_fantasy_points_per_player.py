import unittest

from assign_points_to_categories import create_points_data, create_points_for_category
from config import stat_categories
from convert_stats_to_numbers import convert_stats_to_numbers
from fantasy_points_per_player import create_fantasy_points_per_player, create_fantasy_points_per_player_turnovers_half


raw_data = [
    {
        "SORT_AT_PULL": "1",
        "PLAYER": "LeBron James",
        "TEAM": "CLE", "AGE": "33",
        "GP": "82",
        "W": "50",
        "L": "32",
        "MIN": "3026",
        "PTS": "2251",
        "FGM": "857",
        "FGA": "1580",
        "FG%": "54.2",
        "3PM": "149",
        "3PA": "406",
        "3P%": "36.7",
        "FTM": "388",
        "FTA": "531",
        "FT%": "73.1",
        "OREB": "97",
        "DREB": "612",
        "REB": "709",
        "AST": "747",
        "TOV": "347",
        "STL": "116",
        "BLK": "71",
        "PF": "136",
        "FP": "4436.3",
        "DD2": "52",
        "TD3": "18",
        "+/-": "105",
        "SEASON": "2017-2018"
    },
    {
        "SORT_AT_PULL": "2",
        "PLAYER": "James Harden",
        "TEAM": "HOU", "AGE": "28",
        "GP": "72",
        "W": "59",
        "L": "13",
        "MIN": "2551",
        "PTS": "2191",
        "FGM": "651",
        "FGA": "1449",
        "FG%": "44.9",
        "3PM": "265",
        "3PA": "722",
        "3P%": "36.7",
        "FTM": "624",
        "FTA": "727",
        "FT%": "85.8",
        "OREB": "41",
        "DREB": "348",
        "REB": "389",
        "AST": "630",
        "TOV": "315",
        "STL": "126",
        "BLK": "50",
        "PF": "169",
        "FP": "3815.8",
        "DD2": "31",
        "TD3": "4",
        "+/-": "522",
        "SEASON": "2017-2018"
    },
    {
        "SORT_AT_PULL": "3",
        "PLAYER": "Anthony Davis",
        "TEAM": "NOP",
        "AGE": "25",
        "GP": "75",
        "W": "45",
        "L": "30",
        "MIN": "2727",
        "PTS": "2110",
        "FGM": "780",
        "FGA": "1462",
        "FG%": "53.4",
        "3PM": "55",
        "3PA": "162",
        "3P%": "34",
        "FTM": "495",
        "FTA": "598",
        "FT%": "82.8",
        "OREB": "187",
        "DREB": "644",
        "REB": "831",
        "AST": "174",
        "TOV": "162",
        "STL": "115",
        "BLK": "193",
        "PF": "159",
        "FP": "4130.2",
        "DD2": "50",
        "TD3": "1",
        "+/-": "294",
        "SEASON": "2017-2018"
    },
]

formatted_data = convert_stats_to_numbers(raw_data, stat_categories)

points_due_per_value = create_points_data(
    formatted_data, stat_categories)
            

class TestFantasyPointsPerPlayer(unittest.TestCase):
    def test_fantasy_points_per_player(self):
            
        actual = create_fantasy_points_per_player(
            formatted_data, points_due_per_value, stat_categories)
            
        expected = [{
            'FG%': 3,
            'FT%': 1,
            '3PM': 2,
            'PTS': 3,
            'REB': 2,
            'AST': 3,
            'STL': 2,
            'BLK': 2,
            'TOV': 1,
            'TOTAL': 19,
            'PLAYER': 'LeBron James'
        }, {
            'FG%': 1,
            'FT%': 3,
            '3PM': 3,
            'PTS': 2,
            'REB': 1,
            'AST': 2,
            'STL': 3,
            'BLK': 1,
            'TOV': 2,
            'TOTAL': 18,
            'PLAYER': 'James Harden'
        }, {
            'FG%': 2,
            'FT%': 2,
            '3PM': 1,
            'PTS': 1,
            'REB': 3,
            'AST': 1,
            'STL': 1,
            'BLK': 3,
            'TOV': 3,
            'TOTAL': 17,
            'PLAYER': 'Anthony Davis'
        }]

        self.assertEqual(actual, expected)
        
        
    def test_fantasy_points_per_player_turnovers_half(self):
            
        actual = create_fantasy_points_per_player_turnovers_half(
            formatted_data, points_due_per_value, stat_categories)
            
        expected = [{
            'FG%': 3,
            'FT%': 1,
            '3PM': 2,
            'PTS': 3,
            'REB': 2,
            'AST': 3,
            'STL': 2,
            'BLK': 2,
            'TOV': 0.5,
            'TOTAL': 18.5,
            'PLAYER': 'LeBron James'
        }, {
            'FG%': 1,
            'FT%': 3,
            '3PM': 3,
            'PTS': 2,
            'REB': 1,
            'AST': 2,
            'STL': 3,
            'BLK': 1,
            'TOV': 1,
            'TOTAL': 17,
            'PLAYER': 'James Harden'
        }, {
            'FG%': 2,
            'FT%': 2,
            '3PM': 1,
            'PTS': 1,
            'REB': 3,
            'AST': 1,
            'STL': 1,
            'BLK': 3,
            'TOV': 1.5,
            'TOTAL': 15.5,
            'PLAYER': 'Anthony Davis'
        }]

        self.assertEqual(actual, expected)
