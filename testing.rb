require 'rubygems'
require 'resque'

class FlavorSaver
  @queue = :flavors
  
end

1000.times do
  Resque.enqueue(FlavorSaver,Time.now,"blue")
  sleep(0.75)
end