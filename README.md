# Introduction
This is a tool collecting data from YR for a given location, and using OpenAI chatgpt for writing a descriptive text about the air temperature.

# Requirements
You need a paid OpenAI subscription to use their API.

# Setup
Rename the configs/default.json to configs/config.json, add your OpenAI key (paid access to chatgpt required) and yr user id (email).

# Build
run `$ go build` to create an executable file (rmcode.exe or similar)
