require 'net/http'

def send_message
  url = URI('https://api.telegram.org/bot371317976:AAE0Uz46-0N0aYE3YkyR2syETUkgcbFPES8/sendMessage')

  res = Net::HTTP.post_form(url, chat_id: '103443335', text: 'something is Wrong')
end

def down?
  a = %x(docker-compose ps)
  output = a.to_s.gsub(/\s+/m, ' ').strip.split(' ')
  output.include? 'Exit'
end
if down?
  success = false
  (0..3).each do |i|
    %x(docker-compose up -d)
    sleep(10)
    unless down?
      success = true
      break
    end
  end
end
unless success
  send_message
end

