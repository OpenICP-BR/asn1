# asn1
A fork of the golang's encoding/asn1

# Changes

  * Unexported fields are just ignored instead of returning errors.
  * Support for `asn1:"-"` to ignore certain fields.
  * When marshalling `BeforeASN1Marshalling` will be called on any type that implements `BeforeASN1MarshallingI`.
  * When unmarshalling, a nil pointer to a struct will be dealt with by creating anew struct to put there.
  * Added marshalling option `asn1:"octect"` which will cause the value to be marshalled as an octet string. 
