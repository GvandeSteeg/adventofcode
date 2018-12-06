import unittest

from day1 import *


class MyTestCase(unittest.TestCase):
    def setUp(self):
        self.testlist = ['+1', '+2', '+1', '-1', '+5']


    def test_assignment1(self):
        self.assertEqual(8, assignment1(self.testlist))


    def test_assignment2(self):
        self.assertEqual(3, assignment2(self.testlist))


if __name__ == '__main__':
    unittest.main()
