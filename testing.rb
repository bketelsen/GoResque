require 'rubygems'
require 'resque'

class InquirySaver
  @queue = :inquiries
  
end

1000.times do
  Resque.enqueue(InquirySaver,Time.now,"blue")
  sleep(0.75)
end