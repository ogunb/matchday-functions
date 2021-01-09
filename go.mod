module matchday

go 1.13

replace github.com/ogunb/matchday-functions/reminder => ./reminder

replace github.com/ogunb/matchday-functions/fixture => ./fixture

require github.com/ogunb/matchday-functions/fixture v0.0.0-00010101000000-000000000000 // indirect
