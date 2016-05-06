package main

type CommandProperties struct {
	Name  string
	Group string
	Desc  string
}

/*
func NewGetCommand(props CommandProperties, e *client.EntityEndpoint) *cmd.Command {
	return cmd.NewCommand(

		props.Name, props.Group, props.Desc,

		func(cmd *cmd.Command) {
			cmd.AppendArg("id", "Unique identifier")
		},

		func(cmd *cmd.Command) error {
			id, err := strconv.ParseInt(cmd.Arg("id"), 10, 64)
			if err != nil {
				return err
			}

			cr, res, err := e.Get(id)
			if err != nil {
				return err
			}

			outputResponse(cr, res)

			return nil
		},
	)
}

func NewGetAllCommand(props CommandProperties, e *client.EntityEndpoint) *cmd.Command {
	return cmd.NewCommand(

		props.Name, props.Group, props.Desc,

		func(cmd *cmd.Command) {},

		func(cmd *cmd.Command) error {
			cr, res, err := e.GetAll()
			if err != nil {
				return err
			}

			outputResponse(cr, res)

			return nil
		},
	)
}

func NewDeleteCommand(props CommandProperties, e *client.EntityEndpoint) *cmd.Command {
	return cmd.NewCommand(

		props.Name, props.Group, props.Desc,

		func(cmd *cmd.Command) {
			cmd.AppendArg("id", "Unique identifier")
		},

		func(cmd *cmd.Command) error {
			id, err := strconv.ParseInt(cmd.Arg("id"), 10, 64)
			if err != nil {
				return err
			}

			cr, res, err := e.Delete(id)
			if err != nil {
				return err
			}

			outputResponse(cr, res)

			return nil
		},
	)
}

type AddCommandArg struct {
	ArgName   string
	FieldName string
	ArgDesc   string
}

func NewAddCommand(props CommandProperties, e *client.EntityEndpoint, model interface{}, args ...AddCommandArg) *cmd.Command {
	return cmd.NewCommand(

		props.Name, props.Group, props.Desc,

		func(cmd *cmd.Command) {
			for _, val := range args {
				cmd.AppendArg(val.ArgName, val.ArgDesc)
			}
		},

		func(cmd *cmd.Command) error {
			m := reflect.New(reflect.TypeOf(model)).Elem()
			for _, val := range args {
				m.FieldByName(val.FieldName).SetString(cmd.Arg(val.ArgName))
			}

			b, err := json.Marshal(m.Interface())
			if err != nil {
				return err
			}

			cr, res, err := e.Create(string(b))
			if err != nil {
				return err
			}

			outputResponse(cr, res)

			return nil
		},
	)
}
*/
