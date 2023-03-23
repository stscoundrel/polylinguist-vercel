# Polylinguist Vercel Cloud function

Simple Vercel cloud functions for running [Polylinguist](https://github.com/stscoundrel/polylinguist).

This repo is hardwired for my username, as it uses my access token and custom settings. Feel free to fork it & just change folder structure, username and settings as you please. Access token is expected to be an environment variable, which you can set in Vercel.

Contains endpoints for JSON and text output:
- [Most used languages (text)](https://polylinguist.vercel.app/api/username/stscoundrel/text)
- [Most used languages (json)](https://polylinguist.vercel.app/api/username/stscoundrel/json)
