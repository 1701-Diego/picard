# Picard - Captain's Logs Fetcher

[![Join the chat at https://gitter.im/1701-Diego/picard](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/1701-Diego/picard?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Picard streams Doppler logs for a given log-guid

## To Install

```
go get github.com/1701-diego/picard
```

## To Use

```
picard LOG-GUID
```

Set the Doppler address with the CAPTAINS_LOGS environment:

```
export DOPPLER=wss://doppler.ketchup.cf-app.com:4443
```

Defaults to ketchup without the environment variable.
The address for a local Diego Edge box can be set via: 

```
export DOPPLER=ws://doppler.192.168.11.11.xip.io
```