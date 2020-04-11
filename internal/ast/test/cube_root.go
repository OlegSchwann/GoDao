package test

// goDao: generate
type goDao struct {
	// language=PostgreSQL
	CubeRoot func(cubeVolume float64) (cubeSide float64) `
		select (||/ $1::float8)::float8;`

	// "||/" - the weirdest operator I've ever seen
	// (except perl, of course https://metacpan.org/pod/distribution/perlsecret/lib/perlsecret.pod)
}
