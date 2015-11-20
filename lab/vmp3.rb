#!/usr/bin/env ruby
workdir = ARGV[0] || '.'
Dir.glob("#{workdir}/**/*.{mkv,MKV,mp4,webm}").each { |file| puts `ffmpeg -i '#{file}' -y '#{file.sub(/\.\w+$/, ".mp3")}'` if !File.exist?(file.sub(/\.\w+$/, ".mp3"))}
