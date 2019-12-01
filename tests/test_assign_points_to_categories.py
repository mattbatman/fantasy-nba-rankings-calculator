import unittest

from assign_points_to_categories import create_points_data, create_points_for_category
from config import stat_categories
from convert_stats_to_numbers import convert_stats_to_numbers


class TestRanksByCategory(unittest.TestCase):
    def test_create_points_for_category(self):
        data = [
            {
                'PLAYER': 'Dennis Rodman',
                'PTS': 4
            },
            {
                'PLAYER': 'Meta World Peace',
                'PTS': 10
            },
            {
                'PLAYER': 'Reggie Miller',
                'PTS': 20
            }
        ]

        actual = create_points_for_category(data, 'PTS')

        expected = {
            '4': {
                'count': 1,
                'due': 1,
                'ranks_inverse': [1]
            },
            '10': {
                'count': 1,
                'due': 2,
                'ranks_inverse': [2]
            },
            '20': {
                'count': 1,
                'due': 3,
                'ranks_inverse': [3]
            }
        }

        self.assertEqual(actual, expected)

    def test_create_ranks_data(self):
        data = [
            {
                "SORT_AT_PULL": "1",
                "PLAYER": "LeBron James",
                "TEAM": "CLE",
                "AGE": "33",
                "GP": "82",
                "W": "50",
                "L": "32",
                "MIN": "3026",
                "PTS": 2251.0,
                "FGM": "857",
                "FGA": "1580",
                "FG%": 54.2,
                "3PM": 149.0,
                "3PA": "406",
                "3P%": "36.7",
                "FTM": "388",
                "FTA": "531",
                "FT%": 73.1,
                "OREB": "97",
                "DREB": "612",
                "REB": 709.0,
                "AST": 747.0,
                "TOV": 347.0,
                "STL": 116.0,
                "BLK": 71.0,
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
                "TEAM": "HOU",
                "AGE": "28",
                "GP": "72",
                "W": "59",
                "L": "13",
                "MIN": "2551",
                "PTS": 2191.0,
                "FGM": "651",
                "FGA": "1449",
                "FG%": 44.9,
                "3PM": 265.0,
                "3PA": "722",
                "3P%": "36.7",
                "FTM": "624",
                "FTA": "727",
                "FT%": 85.8,
                "OREB": "41",
                "DREB": "348",
                "REB": 389.0,
                "AST": 630.0,
                "TOV": 315.0,
                "STL": 126.0,
                "BLK": 50.0,
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
                "PTS": 2110.0,
                "FGM": "780",
                "FGA": "1462",
                "FG%": 53.4,
                "3PM": 55.0,
                "3PA": "162",
                "3P%": "34",
                "FTM": "495",
                "FTA": "598",
                "FT%": 82.8,
                "OREB": "187",
                "DREB": "644",
                "REB": 831.0,
                "AST": 174.0,
                "TOV": 162.0,
                "STL": 115.0,
                "BLK": 193.0,
                "PF": "159",
                "FP": "4130.2",
                "DD2": "50",
                "TD3": "1",
                "+/-": "294",
                "SEASON": "2017-2018"
            },
        ]

        actual = create_points_data(data, stat_categories)

        expected = {
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
            '3PM': {
                '55.0': {'count': 1, 'due': 1, 'ranks_inverse': [1]},
                '149.0': {'count': 1, 'due': 2, 'ranks_inverse': [2]},
                '265.0': {'count': 1, 'due': 3, 'ranks_inverse': [3]},
            },
            'PTS': {
                '2110.0': {'count': 1, 'due': 1, 'ranks_inverse': [1]},
                '2191.0': {'count': 1, 'due': 2, 'ranks_inverse': [2]},
                '2251.0': {'count': 1, 'due': 3, 'ranks_inverse': [3]},
            },
            'REB': {
                '389.0': {'count': 1, 'due': 1, 'ranks_inverse': [1]},
                '709.0': {'count': 1, 'due': 2, 'ranks_inverse': [2]},
                '831.0': {'count': 1, 'due': 3, 'ranks_inverse': [3]},
            },
            'AST': {
                '174.0': {'count': 1, 'due': 1, 'ranks_inverse': [1]},
                '630.0': {'count': 1, 'due': 2, 'ranks_inverse': [2]},
                '747.0': {'count': 1, 'due': 3, 'ranks_inverse': [3]},
            },
            'STL': {
                '115.0': {'count': 1, 'due': 1, 'ranks_inverse': [1]},
                '116.0': {'count': 1, 'due': 2, 'ranks_inverse': [2]},
                '126.0': {'count': 1, 'due': 3, 'ranks_inverse': [3]},
            },
            'BLK': {
                '50.0': {'count': 1, 'due': 1, 'ranks_inverse': [1]},
                '71.0': {'count': 1, 'due': 2, 'ranks_inverse': [2]},
                '193.0': {'count': 1, 'due': 3, 'ranks_inverse': [3]},
            },
            'TOV': {
                '347.0': {'count': 1, 'due': 1, 'ranks_inverse': [1]},
                '315.0': {'count': 1, 'due': 2, 'ranks_inverse': [2]},
                '162.0': {'count': 1, 'due': 3, 'ranks_inverse': [3]},
            }
        }

        self.assertEqual(actual, expected)
