import unittest

from pyfakefs.fake_filesystem_unittest import TestCase

from day1 import *


class TestDay1(TestCase):
    def setUp(self):
        self.testlist = ['+1', '+2', '+1', '-1', '+5']
        self.setUpPyfakefs()


    def test_parse_input_file(self):
        self.fs.create_file("input.txt", contents='\n'.join(self.testlist))
        self.assertEqual(self.testlist, parse_input_file("input.txt"))


    def test_parse_numbers(self):
        self.assertEqual(12, parse_numbers('+12'))
        self.assertEqual(-12, parse_numbers('-12'))
        self.assertRaises(TypeError, parse_numbers, '12')
        self.assertRaises(TypeError, parse_numbers, 12)


    def test_assignment1(self):
        self.assertEqual(8, assignment1(self.testlist))


    def test_assignment2(self):
        self.assertEqual(3, assignment2(self.testlist))


if __name__ == '__main__':
    unittest.main()
