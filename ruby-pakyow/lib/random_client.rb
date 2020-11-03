require 'httparty'

class RandomClient
  RANDOM_ORG_API_KEY = ENV.fetch('RANDOM_ORG_API_KEY')

  def self.invoke(request)
    HTTParty.post(
      'https://api.random.org/json-rpc/2/invoke',
      body: request.to_json,
      headers: {
        'Content-Type': 'application/json'
      }
    ).parsed_response
  end

  def self.integers(amount, size)
    response = invoke({
      id: 1,
      jsonrpc: '2.0',
      method: 'generateIntegers',
      params: {
        apiKey: RANDOM_ORG_API_KEY,
        max: size,
        min: 1,
        n: amount
      }
    })

    response['result']['random']['data']
  end
end
