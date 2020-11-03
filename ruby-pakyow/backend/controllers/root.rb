require 'json'

require 'dice.rb'
require 'random_client.rb'

controller :root do
  get "/" do
    @connection << File.read("../shared/index.go.html")
  end

  get "/styles.css" do
    @connection << File.read("../shared/styles.css")
  end

  get "/main.js" do
    @connection << File.read("../shared/main.js")
  end

  get "/roll/:dice" do
    dice = Dice.parse(@connection.params['dice'])

    @connection << {
      result: RandomClient.integers(dice.amount, dice.size).sum
    }.to_json
  end
end
