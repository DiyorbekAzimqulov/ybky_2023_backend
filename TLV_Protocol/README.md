# Type, Length, Value Protocol

## Introduction

The Type, Length, Value (TLV) protocol is a simple protocol that is used to encapsulate data in a message.
[JPRQ](https://jprq.io/) project is used LV part of TLV protocol to encapsulate data in a message.

## JPRQ LV protocol analysis

You can find the source code here: [JPRQ LV protocol](https://github.com/azimjohn/jprq/blob/master/server/events/events.go)
There are 3 types of events: TunnelOpened, ConnectionReceived, TunnerRequested.
They are struct and all of them implemented `.encode()` and `.decode()` methods.
The `.encode()` method is used to encode the event into a byte array.
The `.decode()` method is used to decode the byte array into an event. So byte array comes and with decode method that byte array is converted to an event and stored.

To convert from golang data type to byte array, [GOB](https://pkg.go.dev/encoding/gob) encoding/decoding is used.
