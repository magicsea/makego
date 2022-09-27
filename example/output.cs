public class Tables
{
	public Tables(Func<string, Stream> loader){
		
 		_ConstTable = new Table.ConstTable();
        var streamConstTable = loader("ConstTable");
        var readerConstTable = new tabtoy.DataReader(streamConstTable) ;
        CheckHeader(readerConstTable, ConstTable);
        Table.ConstTable.DeserializeGMTable(ConstTable, readerConstTable);
        streamConstTable.Dispose();
		
 		_GMTable = new Table.GMTable();
        var streamGMTable = loader("GMTable");
        var readerGMTable = new tabtoy.DataReader(streamGMTable) ;
        CheckHeader(readerGMTable, GMTable);
        Table.GMTable.DeserializeGMTable(GMTable, readerGMTable);
        streamGMTable.Dispose();
		
	}

	
	public Table.ConstTable ConstTable { get { return _ConstTable; } }
	private Table.ConstTable _ConstTable;
	
	public Table.GMTable GMTable { get { return _GMTable; } }
	private Table.GMTable _GMTable;
	

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