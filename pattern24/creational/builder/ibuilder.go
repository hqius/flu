package builder

// select A from table where B order by C limit D, E
// column builder
// table builder
// where builder
//   -> and or
//   -> lt lte eq gt gte
// order by
//   -> order builder
// limit
//   -> builder

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}
