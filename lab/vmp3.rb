#!/usr/bin/env ruby
workdir = ARGV[0] || '.'
Dir.glob("#{workdir}/**/*.{mkv,MKV,mp4,webm}").each { |file| mp3file=file.sub(/\.\w+$/, ".mp3"); puts `ffmpeg -i '#{file}' -y '#{mp3file}'` if !File.exist?(mp3file)}
