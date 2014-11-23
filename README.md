# Picard - Captain's Logs Fetcher

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