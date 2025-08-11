package kafka

type RequestHeaderV2 struct {
	RequestAPIkey int16
	RequestAPIver int16
	CorrelationID int32
	ClientID      []byte
}
