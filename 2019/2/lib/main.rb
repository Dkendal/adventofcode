require "minitest/autorun"

class Test < Minitest::Test
  def test_example_1
    assert_equal call([1, 0, 0, 0, 99]), [2, 0, 0, 0, 99]
  end

  def test_example_2
    assert_equal call([2, 3, 0, 3, 99]), [2, 3, 0, 6, 99]
  end

  def test_example_3
    assert_equal call([2, 4, 4, 5, 99, 0]), [2, 4, 4, 5, 99, 9801]
  end

  def test_example_4
    assert_equal call([1, 1, 1, 4, 99, 5, 6, 0, 99]), [30, 1, 1, 4, 2, 5, 6, 0, 99]
  end
end

def call(data)
  data
    .each_slice(4) { |op, a, b, c|
    case op
    when 1
      data[c] = data[a] + data[b]
    when 2
      data[c] = data[a] * data[b]
    when 99
      break
    end
  }
  data
end

def main
  bak = STDIN.each_line(",").map(&:to_i)

  data = bak.dup
  data[1] = 12
  data[2] = 2
  print "answer 1: "
  puts call(data)[0]

  (0..99).to_a.repeated_permutation(2) do |x, y|
    data = bak.dup
    data[1] = x
    data[2] = y
    if call(data)[0] == 19690720
      print "answer 2: "
      puts 100 * x + y
    end
  end
end
