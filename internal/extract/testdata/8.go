package testdata

// входные параметры: 1 шт невалидный параметр, структура. Проверить ошибку.
// выходные параметры: (float32, err error)

type Triplet struct {
	beginning float32
	middle    float32
	end       float32
}

type GoDao8 struct {
	// language=PostgreSQL
	TripletSum func(triplet Triplet) (sum float32, err error) `
        select $1::float4 + $2::float4 + $3::float4;`
}
