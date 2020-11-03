class Dice
  def initialize(size = 6, amount = 1)
    @amount = amount
    @size = size
  end

  def amount
    @amount
  end

  def size
    @size
  end

  def self.parse(string)
    pieces = string.split('d')

    if pieces.size == 0 then
      Dice.new()
    elsif pieces.size == 1
      Dice.new(pieces[0])
    else
      Dice.new(pieces[1], pieces[0].empty? ? 1 : pieces[0])
    end
  end
end
