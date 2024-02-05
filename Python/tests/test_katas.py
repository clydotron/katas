import codewars_test as test
from . import hamster_me

@test.describe("Fixed Tests")
def fixed_tests():
    @test.it('Basic Test Cases')
    def basic_test_cases():
        test.assert_equals(hamster_me("hamster", "hamster"), "h1a1m1s1t1e1r1")
        test.assert_equals(hamster_me("hamster", "helpme"), "h1e1h5m4m1e1")
        test.assert_equals(hamster_me("hmster", "hamster"), "h1t8m1s1t1e1r1")
        test.assert_equals(hamster_me("hhhhammmstteree", "hamster"), "h1a1m1s1t1e1r1")
        test.assert_equals(hamster_me("hamster", "abcdefghijklmnopqrstuvwxyz") , "a1a2a3a4e1e2e3h1h2h3h4h5m1m2m3m4m5r1s1t1t2t3t4t5t6t7")
        test.assert_equals(hamster_me("f", "abcdefghijklmnopqrstuvwxyz"), "f22f23f24f25f26f1f2f3f4f5f6f7f8f9f10f11f12f13f14f15f16f17f18f19f20f21")