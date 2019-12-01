import unittest

from convert_stats_to_numbers import convert_stats_to_numbers

test_categories = [
    'PTS'
]

start_test_data = [
    {
        'PTS': '1',
        'OTHER': '2'
    },
    {
        'PTS': '3',
        'OTHER': '4'
    }
]


class TestConvertStatsToNumbers(unittest.TestCase):
    def test(self):
        actual = convert_stats_to_numbers(start_test_data, test_categories)
        expected = [
            {
                'PTS': 1,
                'OTHER': '2'
            },
            {
                'PTS': 3,
                'OTHER': '4'
            }
        ]
        self.assertEqual(actual, expected)
