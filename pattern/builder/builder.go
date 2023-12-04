package builder

type IBuilder interface {
	WindowType(windowType string) IBuilder
	DoorType(doorType string) IBuilder
	Floor(floor int) IBuilder
	Build() *House
}

func Builder(builderType string) IBuilder {
	switch builderType {
	case "normal":
		return &NormalBuilder{}
	case "igloo":
		return &IglooBuilder{}
	}
	return nil
}

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (builder *NormalBuilder) WindowType(windowType string) IBuilder {
	builder.windowType = windowType
	return builder
}

func (builder *NormalBuilder) DoorType(doorType string) IBuilder {
	builder.doorType = doorType
	return builder
}

func (builder *NormalBuilder) Floor(floor int) IBuilder {
	builder.floor = floor
	return builder
}

func (builder *NormalBuilder) Build() *House {
	return &House{
		windowType: builder.windowType,
		doorType:   builder.doorType,
		floor:      builder.floor,
	}
}

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (builder *IglooBuilder) WindowType(windowType string) IBuilder {
	builder.windowType = windowType
	return builder
}

func (builder *IglooBuilder) DoorType(doorType string) IBuilder {
	builder.doorType = doorType
	return builder
}

func (builder *IglooBuilder) Floor(floor int) IBuilder {
	builder.floor = floor
	return builder
}

func (builder *IglooBuilder) Build() *House {
	return &House{
		windowType: builder.windowType,
		doorType:   builder.doorType,
		floor:      builder.floor,
	}
}
