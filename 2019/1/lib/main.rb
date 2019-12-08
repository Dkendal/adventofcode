def fuel(mass)
  (mass / 3.0).floor - 2
end

def fuel_r(mass)
  fuel_mass = (mass / 3.0).floor - 2
  return 0 if fuel_mass <= 0
  fuel_mass + fuel_r(fuel_mass)
end

def main
  input = STDIN.each_line.map(&:to_i)

  print "answer 1: "
  puts input.sum { |mass| fuel(mass) }

  print "answer 2: "
  puts input.sum { |mass| fuel_r(mass) }
end
