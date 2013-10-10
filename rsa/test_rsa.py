import unittest
import rsa

class TestRSAFunctions(unittest.TestCase):

    def test_multiplicative_inverse(self):
        tests = [(120, 23, 47),(160, 7, 23),(40, 3, 27),(48, 5, 29)]
        for test in tests:
            phin, e, ans = test
            self.assertEqual(rsa.multiplicative_inverse(phin, e), ans)

if __name__ == '__main__':
    unittest.main()