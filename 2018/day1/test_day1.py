import unittest

from day1 import assignment1

class MyTestCase(unittest.TestCase):
    def test_assignment1(self):
        self.assertEqual(420, assignment1())


if __name__ == '__main__':
    unittest.main()
