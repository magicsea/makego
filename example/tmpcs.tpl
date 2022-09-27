public class Tables
{
	public Tables(Func<string, Stream> loader){
		{{range .Structs}}
 		_{{.Data 1}} = new Table.{{.Data 1}}();
        var stream{{.Data 1}} = loader("{{.Data 1}}");
        var reader{{.Data 1}} = new tabtoy.DataReader(stream{{.Data 1}}) ;
        CheckHeader(reader{{.Data 1}}, {{.Data 1}});
        Table.{{.Data 1}}.DeserializeGMTable({{.Data 1}}, reader{{.Data 1}});
        stream{{.Data 1}}.Dispose();
		{{end}}
	}

	{{range .Structs}}
	public Table.{{.Data 1}} {{.Data 1}} { get { return _{{.Data 1}}; } }
	private Table.{{.Data 1}} _{{.Data 1}};
	{{end}}

	private void CheckHeader(DataReader reader, ITable table)
    {
        var result = reader.ReadHeader(table.GetBuildID());
        if (result != FileState.OK)
        {
            Console.WriteLine($"combine file crack! {table.GetTableName()}");
            return;
        }
    }
}