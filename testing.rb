require 'rubygems'
require 'resque'

class InquirySaver
  @queue = :inquiries
  
end

100.times do
  Resque.enqueue(InquirySaver,1,"blue")
end