# volumetest

Story: There are over 100,000 flights a day, with millions of people and cargo being transferred around the world. With so many people, and different carrier/agency groups it can be hard to track where a person might be. In order to determine the flight path of a person, we must sort through all of their flight records.

Goal: To create a microservice API that can help us understand and track how a particular person’s flight path may be queried. The API should accept a request that includes a list of flights, which are defined by a source and destination airport code. These flights may not be listed in order and will need to be sorted to find the total flight paths starting and ending airports.

Examples:

[['SFO', 'EWR']]                                                                            => ['SFO', 'EWR']

[['ATL', 'EWR'], ['SFO', 'ATL']]                                                            => ['SFO', 'EWR']

[['IND', 'EWR'], ['SFO', 'ATL'], ['GSO', 'IND'], ['ATL', 'GSO']]                            => ['SFO', 'EWR']


# run

go run main.go

curl -X POST  -d '{"flights": [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]}' http://localhost:8080/

Request data format:

type FlightsRequest struct {
    Flights [][]string `json:"flights"`
}

{"flights": [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]}

Response data format:

type FlightsResponse struct {
    Result []string `json:"result"`
    Error  string   `json:"error"`
}

{"result":["SFO","EWR"],"error":""}